// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

module github.com/trustbloc/fabric-peer-ext/mod/peer

replace github.com/hyperledger/fabric => github.com/trustbloc/fabric-mod v0.1.1-0.20191210140617-384d402d27ae

replace github.com/hyperledger/fabric/extensions => ./

replace github.com/trustbloc/fabric-peer-ext => ../..

replace github.com/hyperledger/fabric-protos-go => github.com/trustbloc/fabric-protos-go-ext v0.1.1-0.20191126151100-5a61374c2e1b

require (
	github.com/golang/protobuf v1.3.2
	github.com/hyperledger/fabric v2.0.0-alpha+incompatible
	github.com/hyperledger/fabric-chaincode-go v0.0.0-20191108205148-17c4b2760b56
	github.com/hyperledger/fabric-protos-go v0.0.0-20191121202242-f5500d5e3e85
	github.com/hyperledger/fabric/extensions v0.0.0
	github.com/magiconair/properties v1.8.1
	github.com/pkg/errors v0.8.1
	github.com/spf13/viper v1.3.2
	github.com/stretchr/testify v1.4.0
	github.com/trustbloc/fabric-peer-ext v0.0.0
)

go 1.13
