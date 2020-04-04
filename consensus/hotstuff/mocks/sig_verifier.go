// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dapperlabs/flow-go/consensus/hotstuff (interfaces: SigVerifier)

// Package mocks is a generated GoMock package.
package mocks

import (
	model "github.com/dapperlabs/flow-go/consensus/hotstuff/model"
	crypto "github.com/dapperlabs/flow-go/crypto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSigVerifier is a mock of SigVerifier interface
type MockSigVerifier struct {
	ctrl     *gomock.Controller
	recorder *MockSigVerifierMockRecorder
}

// MockSigVerifierMockRecorder is the mock recorder for MockSigVerifier
type MockSigVerifierMockRecorder struct {
	mock *MockSigVerifier
}

// NewMockSigVerifier creates a new mock instance
func NewMockSigVerifier(ctrl *gomock.Controller) *MockSigVerifier {
	mock := &MockSigVerifier{ctrl: ctrl}
	mock.recorder = &MockSigVerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSigVerifier) EXPECT() *MockSigVerifierMockRecorder {
	return m.recorder
}

// VerifyRandomBeaconSig mocks base method
func (m *MockSigVerifier) VerifyRandomBeaconSig(arg0 crypto.Signature, arg1 *model.Block, arg2 crypto.PublicKey) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyRandomBeaconSig", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyRandomBeaconSig indicates an expected call of VerifyRandomBeaconSig
func (mr *MockSigVerifierMockRecorder) VerifyRandomBeaconSig(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyRandomBeaconSig", reflect.TypeOf((*MockSigVerifier)(nil).VerifyRandomBeaconSig), arg0, arg1, arg2)
}

// VerifyRandomBeaconThresholdSig mocks base method
func (m *MockSigVerifier) VerifyRandomBeaconThresholdSig(arg0 crypto.Signature, arg1 *model.Block, arg2 crypto.PublicKey) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyRandomBeaconThresholdSig", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyRandomBeaconThresholdSig indicates an expected call of VerifyRandomBeaconThresholdSig
func (mr *MockSigVerifierMockRecorder) VerifyRandomBeaconThresholdSig(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyRandomBeaconThresholdSig", reflect.TypeOf((*MockSigVerifier)(nil).VerifyRandomBeaconThresholdSig), arg0, arg1, arg2)
}

// VerifyStakingAggregatedSig mocks base method
func (m *MockSigVerifier) VerifyStakingAggregatedSig(arg0 []crypto.Signature, arg1 *model.Block, arg2 []crypto.PublicKey) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyStakingAggregatedSig", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyStakingAggregatedSig indicates an expected call of VerifyStakingAggregatedSig
func (mr *MockSigVerifierMockRecorder) VerifyStakingAggregatedSig(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyStakingAggregatedSig", reflect.TypeOf((*MockSigVerifier)(nil).VerifyStakingAggregatedSig), arg0, arg1, arg2)
}

// VerifyStakingSig mocks base method
func (m *MockSigVerifier) VerifyStakingSig(arg0 crypto.Signature, arg1 *model.Block, arg2 crypto.PublicKey) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyStakingSig", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyStakingSig indicates an expected call of VerifyStakingSig
func (mr *MockSigVerifierMockRecorder) VerifyStakingSig(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyStakingSig", reflect.TypeOf((*MockSigVerifier)(nil).VerifyStakingSig), arg0, arg1, arg2)
}