/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package mocks

import (
	"sync/atomic"

	"github.com/hyperledger/fabric/core/common/ccprovider"
	"github.com/hyperledger/fabric/extensions/gossip/api"
	cb "github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/ledger/rwset/kvrwset"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// MockBlockHandler is a mock block handler
type MockBlockHandler struct {
	numReads           int32
	numWrites          int32
	numCCEvents        int32
	numCCUpgradeEvents int32
	numConfigUpdates   int32
	numLSCCWrites      int32
	err                error
}

// NewMockBlockHandler returns a mock Block Handler
func NewMockBlockHandler() *MockBlockHandler {
	return &MockBlockHandler{}
}

// WithError sets an error
func (m *MockBlockHandler) WithError(err error) *MockBlockHandler {
	m.err = err
	return m
}

// NumReads returns the number of reads handled
func (m *MockBlockHandler) NumReads() int {
	return int(atomic.LoadInt32(&m.numReads))
}

// NumWrites returns the number of writes handled
func (m *MockBlockHandler) NumWrites() int {
	return int(atomic.LoadInt32(&m.numWrites))
}

// NumLSCCWrites returns the number of LSCC writes handled
func (m *MockBlockHandler) NumLSCCWrites() int {
	return int(atomic.LoadInt32(&m.numLSCCWrites))
}

// NumCCEvents returns the number of chaincode events handled
func (m *MockBlockHandler) NumCCEvents() int {
	return int(atomic.LoadInt32(&m.numCCEvents))
}

// NumCCUpgradeEvents returns the number of chaincode upgrades handled
func (m *MockBlockHandler) NumCCUpgradeEvents() int {
	return int(atomic.LoadInt32(&m.numCCUpgradeEvents))
}

// NumConfigUpdates returns the number of configuration updates handled
func (m *MockBlockHandler) NumConfigUpdates() int {
	return int(atomic.LoadInt32(&m.numConfigUpdates))
}

// HandleRead handles a read event by incrementing the read counter
func (m *MockBlockHandler) HandleRead(txMetadata api.TxMetadata, namespace string, kvRead *kvrwset.KVRead) error {
	atomic.AddInt32(&m.numReads, 1)
	return m.err
}

// HandleWrite handles a write event by incrementing the write counter
func (m *MockBlockHandler) HandleWrite(txMetadata api.TxMetadata, namespace string, kvWrite *kvrwset.KVWrite) error {
	atomic.AddInt32(&m.numWrites, 1)
	return m.err
}

// HandleChaincodeEvent handle a chaincode event by incrementing the CC event counter
func (m *MockBlockHandler) HandleChaincodeEvent(txMetadata api.TxMetadata, event *pb.ChaincodeEvent) error {
	atomic.AddInt32(&m.numCCEvents, 1)
	return m.err
}

// HandleChaincodeUpgradeEvent handles a chaincode upgrade event by incrementing the chaincode upgrade counter
func (m *MockBlockHandler) HandleChaincodeUpgradeEvent(txMetadata api.TxMetadata, chaincodeName string) error {
	atomic.AddInt32(&m.numCCUpgradeEvents, 1)
	return m.err
}

// HandleConfigUpdate handles a config update by incrementing the config update counter
func (m *MockBlockHandler) HandleConfigUpdate(blockNum uint64, configUpdate *cb.ConfigUpdate) error {
	atomic.AddInt32(&m.numConfigUpdates, 1)
	return m.err
}

// HandleLSCCWrite handles an LSCC write by incrementing the LSCC write counter
func (m *MockBlockHandler) HandleLSCCWrite(txMetadata api.TxMetadata, chaincodeName string, ccData *ccprovider.ChaincodeData, ccp *cb.CollectionConfigPackage) error {
	atomic.AddInt32(&m.numLSCCWrites, 1)
	return m.err
}
