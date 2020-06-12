// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	pb "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel/invoke"
	fabApi "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
)

type TxnClient struct {
	InvokeHandlerStub        func(handler invoke.Handler, request channel.Request, options ...channel.RequestOption) (channel.Response, error)
	invokeHandlerMutex       sync.RWMutex
	invokeHandlerArgsForCall []struct {
		handler invoke.Handler
		request channel.Request
		options []channel.RequestOption
	}
	invokeHandlerReturns struct {
		result1 channel.Response
		result2 error
	}
	invokeHandlerReturnsOnCall map[int]struct {
		result1 channel.Response
		result2 error
	}
	ComputeTxnIDStub        func(nonce []byte) (string, error)
	computeTxnIDMutex       sync.RWMutex
	computeTxnIDArgsForCall []struct {
		nonce []byte
	}
	computeTxnIDReturns struct {
		result1 string
		result2 error
	}
	computeTxnIDReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	SigningIdentityStub        func() ([]byte, error)
	signingIdentityMutex       sync.RWMutex
	signingIdentityArgsForCall []struct{}
	signingIdentityReturns     struct {
		result1 []byte
		result2 error
	}
	signingIdentityReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	GetPeerStub        func(endpoint string) (fabApi.Peer, error)
	getPeerMutex       sync.RWMutex
	getPeerArgsForCall []struct {
		endpoint string
	}
	getPeerReturns struct {
		result1 fabApi.Peer
		result2 error
	}
	getPeerReturnsOnCall map[int]struct {
		result1 fabApi.Peer
		result2 error
	}
	VerifyProposalSignatureStub        func(signedProposal *pb.SignedProposal) error
	verifyProposalSignatureMutex       sync.RWMutex
	verifyProposalSignatureArgsForCall []struct {
		signedProposal *pb.SignedProposal
	}
	verifyProposalSignatureReturns struct {
		result1 error
	}
	verifyProposalSignatureReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *TxnClient) InvokeHandler(handler invoke.Handler, request channel.Request, options ...channel.RequestOption) (channel.Response, error) {
	fake.invokeHandlerMutex.Lock()
	ret, specificReturn := fake.invokeHandlerReturnsOnCall[len(fake.invokeHandlerArgsForCall)]
	fake.invokeHandlerArgsForCall = append(fake.invokeHandlerArgsForCall, struct {
		handler invoke.Handler
		request channel.Request
		options []channel.RequestOption
	}{handler, request, options})
	fake.recordInvocation("InvokeHandler", []interface{}{handler, request, options})
	fake.invokeHandlerMutex.Unlock()
	if fake.InvokeHandlerStub != nil {
		return fake.InvokeHandlerStub(handler, request, options...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.invokeHandlerReturns.result1, fake.invokeHandlerReturns.result2
}

func (fake *TxnClient) InvokeHandlerCallCount() int {
	fake.invokeHandlerMutex.RLock()
	defer fake.invokeHandlerMutex.RUnlock()
	return len(fake.invokeHandlerArgsForCall)
}

func (fake *TxnClient) InvokeHandlerArgsForCall(i int) (invoke.Handler, channel.Request, []channel.RequestOption) {
	fake.invokeHandlerMutex.RLock()
	defer fake.invokeHandlerMutex.RUnlock()
	return fake.invokeHandlerArgsForCall[i].handler, fake.invokeHandlerArgsForCall[i].request, fake.invokeHandlerArgsForCall[i].options
}

func (fake *TxnClient) InvokeHandlerReturns(result1 channel.Response, result2 error) {
	fake.InvokeHandlerStub = nil
	fake.invokeHandlerReturns = struct {
		result1 channel.Response
		result2 error
	}{result1, result2}
}

func (fake *TxnClient) InvokeHandlerReturnsOnCall(i int, result1 channel.Response, result2 error) {
	fake.InvokeHandlerStub = nil
	if fake.invokeHandlerReturnsOnCall == nil {
		fake.invokeHandlerReturnsOnCall = make(map[int]struct {
			result1 channel.Response
			result2 error
		})
	}
	fake.invokeHandlerReturnsOnCall[i] = struct {
		result1 channel.Response
		result2 error
	}{result1, result2}
}

func (fake *TxnClient) ComputeTxnID(nonce []byte) (string, error) {
	var nonceCopy []byte
	if nonce != nil {
		nonceCopy = make([]byte, len(nonce))
		copy(nonceCopy, nonce)
	}
	fake.computeTxnIDMutex.Lock()
	ret, specificReturn := fake.computeTxnIDReturnsOnCall[len(fake.computeTxnIDArgsForCall)]
	fake.computeTxnIDArgsForCall = append(fake.computeTxnIDArgsForCall, struct {
		nonce []byte
	}{nonceCopy})
	fake.recordInvocation("ComputeTxnID", []interface{}{nonceCopy})
	fake.computeTxnIDMutex.Unlock()
	if fake.ComputeTxnIDStub != nil {
		return fake.ComputeTxnIDStub(nonce)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.computeTxnIDReturns.result1, fake.computeTxnIDReturns.result2
}

func (fake *TxnClient) ComputeTxnIDCallCount() int {
	fake.computeTxnIDMutex.RLock()
	defer fake.computeTxnIDMutex.RUnlock()
	return len(fake.computeTxnIDArgsForCall)
}

func (fake *TxnClient) ComputeTxnIDArgsForCall(i int) []byte {
	fake.computeTxnIDMutex.RLock()
	defer fake.computeTxnIDMutex.RUnlock()
	return fake.computeTxnIDArgsForCall[i].nonce
}

func (fake *TxnClient) ComputeTxnIDReturns(result1 string, result2 error) {
	fake.ComputeTxnIDStub = nil
	fake.computeTxnIDReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *TxnClient) ComputeTxnIDReturnsOnCall(i int, result1 string, result2 error) {
	fake.ComputeTxnIDStub = nil
	if fake.computeTxnIDReturnsOnCall == nil {
		fake.computeTxnIDReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.computeTxnIDReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *TxnClient) SigningIdentity() ([]byte, error) {
	fake.signingIdentityMutex.Lock()
	ret, specificReturn := fake.signingIdentityReturnsOnCall[len(fake.signingIdentityArgsForCall)]
	fake.signingIdentityArgsForCall = append(fake.signingIdentityArgsForCall, struct{}{})
	fake.recordInvocation("SigningIdentity", []interface{}{})
	fake.signingIdentityMutex.Unlock()
	if fake.SigningIdentityStub != nil {
		return fake.SigningIdentityStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.signingIdentityReturns.result1, fake.signingIdentityReturns.result2
}

func (fake *TxnClient) SigningIdentityCallCount() int {
	fake.signingIdentityMutex.RLock()
	defer fake.signingIdentityMutex.RUnlock()
	return len(fake.signingIdentityArgsForCall)
}

func (fake *TxnClient) SigningIdentityReturns(result1 []byte, result2 error) {
	fake.SigningIdentityStub = nil
	fake.signingIdentityReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *TxnClient) SigningIdentityReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.SigningIdentityStub = nil
	if fake.signingIdentityReturnsOnCall == nil {
		fake.signingIdentityReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.signingIdentityReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *TxnClient) GetPeer(endpoint string) (fabApi.Peer, error) {
	fake.getPeerMutex.Lock()
	ret, specificReturn := fake.getPeerReturnsOnCall[len(fake.getPeerArgsForCall)]
	fake.getPeerArgsForCall = append(fake.getPeerArgsForCall, struct {
		endpoint string
	}{endpoint})
	fake.recordInvocation("GetPeer", []interface{}{endpoint})
	fake.getPeerMutex.Unlock()
	if fake.GetPeerStub != nil {
		return fake.GetPeerStub(endpoint)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getPeerReturns.result1, fake.getPeerReturns.result2
}

func (fake *TxnClient) GetPeerCallCount() int {
	fake.getPeerMutex.RLock()
	defer fake.getPeerMutex.RUnlock()
	return len(fake.getPeerArgsForCall)
}

func (fake *TxnClient) GetPeerArgsForCall(i int) string {
	fake.getPeerMutex.RLock()
	defer fake.getPeerMutex.RUnlock()
	return fake.getPeerArgsForCall[i].endpoint
}

func (fake *TxnClient) GetPeerReturns(result1 fabApi.Peer, result2 error) {
	fake.GetPeerStub = nil
	fake.getPeerReturns = struct {
		result1 fabApi.Peer
		result2 error
	}{result1, result2}
}

func (fake *TxnClient) GetPeerReturnsOnCall(i int, result1 fabApi.Peer, result2 error) {
	fake.GetPeerStub = nil
	if fake.getPeerReturnsOnCall == nil {
		fake.getPeerReturnsOnCall = make(map[int]struct {
			result1 fabApi.Peer
			result2 error
		})
	}
	fake.getPeerReturnsOnCall[i] = struct {
		result1 fabApi.Peer
		result2 error
	}{result1, result2}
}

func (fake *TxnClient) VerifyProposalSignature(signedProposal *pb.SignedProposal) error {
	fake.verifyProposalSignatureMutex.Lock()
	ret, specificReturn := fake.verifyProposalSignatureReturnsOnCall[len(fake.verifyProposalSignatureArgsForCall)]
	fake.verifyProposalSignatureArgsForCall = append(fake.verifyProposalSignatureArgsForCall, struct {
		signedProposal *pb.SignedProposal
	}{signedProposal})
	fake.recordInvocation("VerifyProposalSignature", []interface{}{signedProposal})
	fake.verifyProposalSignatureMutex.Unlock()
	if fake.VerifyProposalSignatureStub != nil {
		return fake.VerifyProposalSignatureStub(signedProposal)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.verifyProposalSignatureReturns.result1
}

func (fake *TxnClient) VerifyProposalSignatureCallCount() int {
	fake.verifyProposalSignatureMutex.RLock()
	defer fake.verifyProposalSignatureMutex.RUnlock()
	return len(fake.verifyProposalSignatureArgsForCall)
}

func (fake *TxnClient) VerifyProposalSignatureArgsForCall(i int) *pb.SignedProposal {
	fake.verifyProposalSignatureMutex.RLock()
	defer fake.verifyProposalSignatureMutex.RUnlock()
	return fake.verifyProposalSignatureArgsForCall[i].signedProposal
}

func (fake *TxnClient) VerifyProposalSignatureReturns(result1 error) {
	fake.VerifyProposalSignatureStub = nil
	fake.verifyProposalSignatureReturns = struct {
		result1 error
	}{result1}
}

func (fake *TxnClient) VerifyProposalSignatureReturnsOnCall(i int, result1 error) {
	fake.VerifyProposalSignatureStub = nil
	if fake.verifyProposalSignatureReturnsOnCall == nil {
		fake.verifyProposalSignatureReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.verifyProposalSignatureReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *TxnClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.invokeHandlerMutex.RLock()
	defer fake.invokeHandlerMutex.RUnlock()
	fake.computeTxnIDMutex.RLock()
	defer fake.computeTxnIDMutex.RUnlock()
	fake.signingIdentityMutex.RLock()
	defer fake.signingIdentityMutex.RUnlock()
	fake.getPeerMutex.RLock()
	defer fake.getPeerMutex.RUnlock()
	fake.verifyProposalSignatureMutex.RLock()
	defer fake.verifyProposalSignatureMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *TxnClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
