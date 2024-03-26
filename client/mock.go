// Code generated by MockGen. DO NOT EDIT.
// Source: ./client/interface.go
//
// Generated by this command:
//
//	mockgen -source=./client/interface.go -destination=./client/mock.go -package=client
//

// Package client is a generated GoMock package.
package client

import (
	context "context"
	reflect "reflect"

	pactus "github.com/pactus-project/pactus/www/grpc/gen/go"
	gomock "go.uber.org/mock/gomock"
)

// MockIClient is a mock of IClient interface.
type MockIClient struct {
	ctrl     *gomock.Controller
	recorder *MockIClientMockRecorder
}

// MockIClientMockRecorder is the mock recorder for MockIClient.
type MockIClientMockRecorder struct {
	mock *MockIClient
}

// NewMockIClient creates a new mock instance.
func NewMockIClient(ctrl *gomock.Controller) *MockIClient {
	mock := &MockIClient{ctrl: ctrl}
	mock.recorder = &MockIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIClient) EXPECT() *MockIClientMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockIClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockIClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIClient)(nil).Close))
}

// GetBalance mocks base method.
func (m *MockIClient) GetBalance(arg0 context.Context, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockIClientMockRecorder) GetBalance(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockIClient)(nil).GetBalance), arg0, arg1)
}

// GetBlockchainHeight mocks base method.
func (m *MockIClient) GetBlockchainHeight(arg0 context.Context) (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockchainHeight", arg0)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockchainHeight indicates an expected call of GetBlockchainHeight.
func (mr *MockIClientMockRecorder) GetBlockchainHeight(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockchainHeight", reflect.TypeOf((*MockIClient)(nil).GetBlockchainHeight), arg0)
}

// GetBlockchainInfo mocks base method.
func (m *MockIClient) GetBlockchainInfo(arg0 context.Context) (*pactus.GetBlockchainInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockchainInfo", arg0)
	ret0, _ := ret[0].(*pactus.GetBlockchainInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockchainInfo indicates an expected call of GetBlockchainInfo.
func (mr *MockIClientMockRecorder) GetBlockchainInfo(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockchainInfo", reflect.TypeOf((*MockIClient)(nil).GetBlockchainInfo), arg0)
}

// GetFee mocks base method.
func (m *MockIClient) GetFee(arg0 context.Context, arg1 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFee", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFee indicates an expected call of GetFee.
func (mr *MockIClientMockRecorder) GetFee(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFee", reflect.TypeOf((*MockIClient)(nil).GetFee), arg0, arg1)
}

// GetNetworkInfo mocks base method.
func (m *MockIClient) GetNetworkInfo(arg0 context.Context) (*pactus.GetNetworkInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkInfo", arg0)
	ret0, _ := ret[0].(*pactus.GetNetworkInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetworkInfo indicates an expected call of GetNetworkInfo.
func (mr *MockIClientMockRecorder) GetNetworkInfo(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkInfo", reflect.TypeOf((*MockIClient)(nil).GetNetworkInfo), arg0)
}

// GetTransactionData mocks base method.
func (m *MockIClient) GetTransactionData(arg0 context.Context, arg1 string) (*pactus.GetTransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionData", arg0, arg1)
	ret0, _ := ret[0].(*pactus.GetTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionData indicates an expected call of GetTransactionData.
func (mr *MockIClientMockRecorder) GetTransactionData(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionData", reflect.TypeOf((*MockIClient)(nil).GetTransactionData), arg0, arg1)
}

// GetValidatorInfo mocks base method.
func (m *MockIClient) GetValidatorInfo(arg0 context.Context, arg1 string) (*pactus.GetValidatorResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorInfo", arg0, arg1)
	ret0, _ := ret[0].(*pactus.GetValidatorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValidatorInfo indicates an expected call of GetValidatorInfo.
func (mr *MockIClientMockRecorder) GetValidatorInfo(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorInfo", reflect.TypeOf((*MockIClient)(nil).GetValidatorInfo), arg0, arg1)
}

// GetValidatorInfoByNumber mocks base method.
func (m *MockIClient) GetValidatorInfoByNumber(arg0 context.Context, arg1 int32) (*pactus.GetValidatorResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorInfoByNumber", arg0, arg1)
	ret0, _ := ret[0].(*pactus.GetValidatorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValidatorInfoByNumber indicates an expected call of GetValidatorInfoByNumber.
func (mr *MockIClientMockRecorder) GetValidatorInfoByNumber(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorInfoByNumber", reflect.TypeOf((*MockIClient)(nil).GetValidatorInfoByNumber), arg0, arg1)
}

// LastBlockTime mocks base method.
func (m *MockIClient) LastBlockTime(arg0 context.Context) (uint32, uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastBlockTime", arg0)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(uint32)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LastBlockTime indicates an expected call of LastBlockTime.
func (mr *MockIClientMockRecorder) LastBlockTime(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastBlockTime", reflect.TypeOf((*MockIClient)(nil).LastBlockTime), arg0)
}
