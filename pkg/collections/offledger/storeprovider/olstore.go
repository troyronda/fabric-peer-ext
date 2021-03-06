/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package storeprovider

import (
	"time"

	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/core/common/privdata"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/rwsetutil"
	storeapi "github.com/hyperledger/fabric/extensions/collections/api/store"
	mspmgmt "github.com/hyperledger/fabric/msp/mgmt"
	"github.com/hyperledger/fabric/protos/common"
	cb "github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/ledger/rwset/kvrwset"
	pb "github.com/hyperledger/fabric/protos/transientstore"
	"github.com/pkg/errors"
	"github.com/trustbloc/fabric-peer-ext/pkg/collections/offledger/storeprovider/store/api"
	"github.com/trustbloc/fabric-peer-ext/pkg/collections/offledger/storeprovider/store/cache"
	"github.com/trustbloc/fabric-peer-ext/pkg/config"
)

var logger = flogging.MustGetLogger("ext_offledger")

type store struct {
	channelID   string
	dbProvider  api.DBProvider
	cache       *cache.Cache
	collConfigs map[common.CollectionType]*collTypeConfig
}

func newStore(channelID string, dbProvider api.DBProvider, collConfigs map[common.CollectionType]*collTypeConfig) *store {
	logger.Debugf("constructing collection data store")

	store := &store{
		channelID:   channelID,
		collConfigs: collConfigs,
		dbProvider:  dbProvider,
	}

	if config.GetOLCollCacheEnabled() {
		store.cache = cache.New(channelID, dbProvider, config.GetOLCollCacheSize())
	}

	return store
}

// Close closes the store
func (s *store) Close() {
	s.dbProvider.Close()
}

// Persist persists all data within the private data simulation results
func (s *store) Persist(txID string, privateSimulationResultsWithConfig *pb.TxPvtReadWriteSetWithConfigInfo) error {
	rwSet, err := rwsetutil.TxPvtRwSetFromProtoMsg(privateSimulationResultsWithConfig.PvtRwset)
	if err != nil {
		return errors.WithMessage(err, "error getting pvt RW set from bytes")
	}

	for _, nsRWSet := range rwSet.NsPvtRwSet {
		for _, collRWSet := range nsRWSet.CollPvtRwSets {
			if err := s.persistColl(txID, nsRWSet.NameSpace, privateSimulationResultsWithConfig.CollectionConfigs, collRWSet); err != nil {
				return err
			}
		}
	}

	return nil
}

// PutData returns the  data for the given key
func (s *store) PutData(config *cb.StaticCollectionConfig, key *storeapi.Key, value *storeapi.ExpiringValue) error {
	if value.Value == nil {
		return errors.Errorf("attempt to put nil value for key [%s]", key)
	}
	if config.Name != key.Collection {
		return errors.Errorf("invalid collection config for key [%s]", key)
	}

	key, value, err := s.beforeSave(config, key, value)
	if err != nil {
		return err
	}

	if !value.Expiry.IsZero() {
		if value.Expiry.Before(time.Now()) {
			// Already expired
			logger.Debugf("[%s] Key [%s] already expired", s.channelID, key)
			return nil
		}
	}

	db, err := s.dbProvider.GetDB(s.channelID, key.Collection, key.Namespace)
	if err != nil {
		return err
	}

	logger.Debugf("[%s] Putting key [%s] to DB", s.channelID, key)
	err = db.Put(api.NewKeyValue(key.Key, value.Value, key.EndorsedAtTxID, value.Expiry))
	if err != nil {
		return err
	}

	if s.cache != nil {
		logger.Debugf("[%s] Putting key [%s] to cache", s.channelID, key)
		s.cache.Put(key.Namespace, key.Collection, key.Key,
			&api.Value{
				Value:      value.Value,
				TxID:       key.EndorsedAtTxID,
				ExpiryTime: value.Expiry,
			},
		)
	}

	return nil
}

// GetData returns the  data for the given key
func (s *store) GetData(key *storeapi.Key) (*storeapi.ExpiringValue, error) {
	return s.getData(key.EndorsedAtTxID, key.Namespace, key.Collection, key.Key)
}

// GetDataMultipleKeys returns the  data for the given keys
func (s *store) GetDataMultipleKeys(key *storeapi.MultiKey) (storeapi.ExpiringValues, error) {
	return s.getDataMultipleKeys(key.EndorsedAtTxID, key.Namespace, key.Collection, key.Keys...)
}

// Query executes the given rich query
func (s *store) Query(key *storeapi.QueryKey) (storeapi.ResultsIterator, error) {
	db, err := s.dbProvider.GetDB(s.channelID, key.Collection, key.Namespace)
	if err != nil {
		return nil, err
	}

	results, err := db.Query(key.Query)
	if err != nil {
		return nil, err
	}

	var queryResults []*storeapi.QueryResult
	for _, result := range results {
		if result.TxID == key.EndorsedAtTxID {
			logger.Debugf("[%s] Key [%s:%s:%s] was persisted in same transaction [%s] as caller. Not adding key to result set.", s.channelID, key.Namespace, key.Collection, result.Key, key.EndorsedAtTxID)
		} else {
			r := &storeapi.QueryResult{
				Key:           storeapi.NewKey(result.TxID, key.Namespace, key.Collection, result.Key),
				ExpiringValue: &storeapi.ExpiringValue{Value: result.Value.Value, Expiry: result.ExpiryTime},
			}
			queryResults = append(queryResults, r)
		}
	}

	return newResultsIterator(queryResults), nil
}

func (s *store) persistColl(txID string, ns string, collConfigPkgs map[string]*common.CollectionConfigPackage, collRWSet *rwsetutil.CollPvtRwSet) error {
	config, exists := s.getCollectionConfig(collConfigPkgs, ns, collRWSet.CollectionName)
	if !exists {
		logger.Debugf("[%s]  config for collection [%s:%s] not found in config packages", s.channelID, ns, collRWSet.CollectionName)
		return nil
	}

	authorized, err := s.isAuthorized(ns, config)
	if err != nil {
		return err
	}
	if !authorized {
		logger.Infof("[%s] Will not store  collection [%s:%s] since local peer is not authorized.", s.channelID, ns, collRWSet.CollectionName)
		return nil
	}

	logger.Debugf("[%s] Collection [%s:%s] is of type [%s]", s.channelID, ns, collRWSet.CollectionName, config.Type)

	expiryTime, err := s.getExpirationTime(config)
	if err != nil {
		return err
	}

	batch, err := s.createBatch(txID, ns, config, collRWSet, expiryTime)
	if err != nil {
		return err
	}

	db, err := s.dbProvider.GetDB(s.channelID, collRWSet.CollectionName, ns)
	if err != nil {
		return err
	}

	err = db.Put(batch...)
	if err != nil {
		return errors.WithMessagef(err, "error persisting to [%s:%s]", ns, collRWSet.CollectionName)
	}

	if s.cache != nil {
		s.updateCache(ns, batch, collRWSet)
	}

	return nil
}

func (s *store) updateCache(ns string, batch []*api.KeyValue, collRWSet *rwsetutil.CollPvtRwSet) {
	for _, kv := range batch {
		if kv.Value != nil {
			logger.Debugf("[%s] Putting key [%s:%s:%s] in Tx [%s]", s.channelID, ns, collRWSet.CollectionName, kv.Key, kv.TxID)
			s.cache.Put(ns, collRWSet.CollectionName, kv.Key, kv.Value)
		} else {
			logger.Debugf("[%s] Deleting key [%s:%s:%s]", s.channelID, ns, collRWSet.CollectionName, kv.Key)
			s.cache.Delete(ns, collRWSet.CollectionName, kv.Key)
		}
	}
}

func (s *store) getData(txID, ns, coll, key string) (*storeapi.ExpiringValue, error) {

	var value *api.Value
	var err error

	if s.cache == nil {
		var db api.DB
		db, err = s.dbProvider.GetDB(s.channelID, coll, ns)
		if err != nil {
			return nil, errors.WithMessage(err, "error getting database")
		}
		value, err = db.Get(key)
		if err != nil {
			return nil, errors.WithMessage(err, "error loading value")
		}
	} else {
		value, err = s.cache.Get(ns, coll, key)
		if err != nil {
			return nil, err
		}
	}

	if value == nil {
		return nil, nil
	}

	logger.Debugf("[%s] Got value for key [%s:%s:%s] which was persisted in transaction [%s]. Current tx [%s]", s.channelID, ns, coll, key, value.TxID, txID)
	if value.TxID == txID {
		logger.Debugf("[%s] Key [%s:%s:%s] was persisted in same transaction [%s] as caller. Returning nil.", s.channelID, ns, coll, key, txID)
		return nil, nil
	}

	return &storeapi.ExpiringValue{Value: value.Value, Expiry: value.ExpiryTime}, nil
}

func (s *store) getDataMultipleKeys(txID, ns, coll string, keys ...string) (storeapi.ExpiringValues, error) {

	var values []*api.Value
	var err error

	if s.cache == nil {
		var db api.DB
		db, err = s.dbProvider.GetDB(s.channelID, coll, ns)
		if err != nil {
			return nil, errors.WithMessage(err, "error getting database")
		}

		values, err = db.GetMultiple(keys...)
		if err != nil {
			return nil, errors.WithMessage(err, "error loading values")
		}
	} else {
		values, err = s.cache.GetMultiple(ns, coll, keys...)
		if err != nil {
			return nil, err
		}
	}

	if len(values) != len(keys) {
		return nil, errors.New("not all of the values were returned for the set of keys")
	}

	var ret storeapi.ExpiringValues
	for i, value := range values {
		var v *storeapi.ExpiringValue
		if value != nil {
			logger.Debugf("[%s] Got value for key [%s:%s:%s] which was persisted in transaction [%s]. Current tx [%s]", s.channelID, ns, coll, keys[i], value.TxID, txID)
			if value.TxID == txID {
				logger.Debugf("[%s] Key [%s:%s:%s] was persisted in same transaction [%s] as caller. Returning nil.", s.channelID, ns, coll, keys[i], txID)
			} else {
				v = &storeapi.ExpiringValue{Value: value.Value, Expiry: value.ExpiryTime}
			}
		}
		ret = append(ret, v)
	}

	return ret, nil
}

func (s *store) createBatch(txID, ns string, config *cb.StaticCollectionConfig, collRWSet *rwsetutil.CollPvtRwSet, expiryTime time.Time) ([]*api.KeyValue, error) {
	var batch []*api.KeyValue
	for _, w := range collRWSet.KvRwSet.Writes {
		kv, err := s.newKeyValue(txID, ns, config, expiryTime, w)
		if err != nil {
			return nil, err
		}
		batch = append(batch, kv)
	}
	return batch, nil
}

func (s *store) newKeyValue(txID, ns string, config *cb.StaticCollectionConfig, expiryTime time.Time, w *kvrwset.KVWrite) (*api.KeyValue, error) {
	key := storeapi.NewKey(txID, ns, config.Name, w.Key)
	if w.IsDelete {
		dKey, err := s.beforeLoad(config, key)
		if err != nil {
			return nil, err
		}
		return &api.KeyValue{Key: dKey.Key}, nil
	}

	dKey, value, err := s.beforeSave(config, key,
		&storeapi.ExpiringValue{
			Value:  w.Value,
			Expiry: expiryTime,
		},
	)
	if err != nil {
		return nil, err
	}
	return api.NewKeyValue(dKey.Key, value.Value, txID, value.Expiry), nil
}

func (s *store) isAuthorized(ns string, config *common.StaticCollectionConfig) (bool, error) {
	policy, err := s.loadPolicy(ns, config)
	if err != nil {
		logger.Errorf("[%s] Error loading policy for collection [%s:%s]: %s", s.channelID, ns, config.Name, err)
		return false, err
	}

	localMSPID, err := getLocalMSPID()
	if err != nil {
		logger.Errorf("[%s] Error getting local MSP ID: %s", s.channelID, err)
		return false, err
	}
	for _, mspID := range policy.MemberOrgs() {
		if mspID == localMSPID {
			return true, nil
		}
	}
	return false, nil
}

// TODO: Consider caching policies to avoid marshalling every time
func (s *store) loadPolicy(ns string, config *common.StaticCollectionConfig) (privdata.CollectionAccessPolicy, error) {
	logger.Debugf("[%s] Loading collection policy for [%s:%s]", s.channelID, ns, config.Name)

	colAP := &privdata.SimpleCollection{}
	err := colAP.Setup(config, mspmgmt.GetIdentityDeserializer(s.channelID))
	if err != nil {
		return nil, errors.Wrapf(err, "error setting up collection policy %s", config.Name)
	}

	return colAP, nil
}

func (s *store) getExpirationTime(config *common.StaticCollectionConfig) (time.Time, error) {
	var expiryTime time.Time
	if config.TimeToLive == "" {
		return expiryTime, nil
	}
	ttl, e := time.ParseDuration(config.TimeToLive)
	if e != nil {
		// This shouldn't happen since the config was validated before being persisted
		return expiryTime, errors.Wrapf(e, "error parsing time-to-live for collection [%s]", config.Name)
	}
	return time.Now().Add(ttl), nil
}

func (s *store) getCollectionConfig(collConfigPkgs map[string]*common.CollectionConfigPackage, namespace, collName string) (*common.StaticCollectionConfig, bool) {
	collConfigPkg, ok := collConfigPkgs[namespace]
	if !ok {
		return nil, false
	}

	for _, collConfig := range collConfigPkg.Config {
		config := collConfig.GetStaticCollectionConfig()
		if config != nil && config.Name == collName && s.collTypeSupported(config.Type) {
			return config, true
		}
	}

	return nil, false
}

func (s *store) beforeSave(config *cb.StaticCollectionConfig, key *storeapi.Key, value *storeapi.ExpiringValue) (*storeapi.Key, *storeapi.ExpiringValue, error) {
	cfg, ok := s.collConfigs[config.Type]
	if !ok || cfg.decorator == nil {
		return key, value, nil
	}
	return cfg.decorator.BeforeSave(key, value)
}

func (s *store) beforeLoad(config *cb.StaticCollectionConfig, key *storeapi.Key) (*storeapi.Key, error) {
	cfg, ok := s.collConfigs[config.Type]
	if !ok || cfg.decorator == nil {
		return key, nil
	}
	return cfg.decorator.BeforeLoad(key)
}

func (s *store) collTypeSupported(collType cb.CollectionType) bool {
	_, ok := s.collConfigs[collType]
	return ok
}

// getLocalMSPID returns the MSP ID of the local peer. This variable may be overridden by unit tests.
var getLocalMSPID = func() (string, error) {
	return mspmgmt.GetLocalMSP().GetIdentifier()
}
