// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/klaytn/klaytn/node/cn/filters (interfaces: Backend)

// Package cn is a generated GoMock package.
package cn

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	blockchain "github.com/klaytn/klaytn/blockchain"
	bloombits "github.com/klaytn/klaytn/blockchain/bloombits"
	types "github.com/klaytn/klaytn/blockchain/types"
	common "github.com/klaytn/klaytn/common"
	event "github.com/klaytn/klaytn/event"
	rpc "github.com/klaytn/klaytn/networks/rpc"
	database "github.com/klaytn/klaytn/storage/database"
)

// MockBackend is a mock of Backend interface
type MockBackend struct {
	ctrl     *gomock.Controller
	recorder *MockBackendMockRecorder
}

// MockBackendMockRecorder is the mock recorder for MockBackend
type MockBackendMockRecorder struct {
	mock *MockBackend
}

// NewMockBackend creates a new mock instance
func NewMockBackend(ctrl *gomock.Controller) *MockBackend {
	mock := &MockBackend{ctrl: ctrl}
	mock.recorder = &MockBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBackend) EXPECT() *MockBackendMockRecorder {
	return m.recorder
}

// BloomStatus mocks base method
func (m *MockBackend) BloomStatus() (uint64, uint64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BloomStatus")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(uint64)
	return ret0, ret1
}

// BloomStatus indicates an expected call of BloomStatus
func (mr *MockBackendMockRecorder) BloomStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BloomStatus", reflect.TypeOf((*MockBackend)(nil).BloomStatus))
}

// ChainDB mocks base method
func (m *MockBackend) ChainDB() database.DBManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainDB")
	ret0, _ := ret[0].(database.DBManager)
	return ret0
}

// ChainDB indicates an expected call of ChainDB
func (mr *MockBackendMockRecorder) ChainDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainDB", reflect.TypeOf((*MockBackend)(nil).ChainDB))
}

// EventMux mocks base method
func (m *MockBackend) EventMux() *event.TypeMux {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventMux")
	ret0, _ := ret[0].(*event.TypeMux)
	return ret0
}

// EventMux indicates an expected call of EventMux
func (mr *MockBackendMockRecorder) EventMux() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventMux", reflect.TypeOf((*MockBackend)(nil).EventMux))
}

// GetBlockReceipts mocks base method
func (m *MockBackend) GetBlockReceipts(arg0 context.Context, arg1 common.Hash) types.Receipts {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockReceipts", arg0, arg1)
	ret0, _ := ret[0].(types.Receipts)
	return ret0
}

// GetBlockReceipts indicates an expected call of GetBlockReceipts
func (mr *MockBackendMockRecorder) GetBlockReceipts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockReceipts", reflect.TypeOf((*MockBackend)(nil).GetBlockReceipts), arg0, arg1)
}

// GetLogs mocks base method
func (m *MockBackend) GetLogs(arg0 context.Context, arg1 common.Hash) ([][]*types.Log, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogs", arg0, arg1)
	ret0, _ := ret[0].([][]*types.Log)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogs indicates an expected call of GetLogs
func (mr *MockBackendMockRecorder) GetLogs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogs", reflect.TypeOf((*MockBackend)(nil).GetLogs), arg0, arg1)
}

// HeaderByNumber mocks base method
func (m *MockBackend) HeaderByNumber(arg0 context.Context, arg1 rpc.BlockNumber) (*types.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeaderByNumber", arg0, arg1)
	ret0, _ := ret[0].(*types.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeaderByNumber indicates an expected call of HeaderByNumber
func (mr *MockBackendMockRecorder) HeaderByNumber(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeaderByNumber", reflect.TypeOf((*MockBackend)(nil).HeaderByNumber), arg0, arg1)
}

// ServiceFilter mocks base method
func (m *MockBackend) ServiceFilter(arg0 context.Context, arg1 *bloombits.MatcherSession) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ServiceFilter", arg0, arg1)
}

// ServiceFilter indicates an expected call of ServiceFilter
func (mr *MockBackendMockRecorder) ServiceFilter(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceFilter", reflect.TypeOf((*MockBackend)(nil).ServiceFilter), arg0, arg1)
}

// SubscribeChainEvent mocks base method
func (m *MockBackend) SubscribeChainEvent(arg0 chan<- blockchain.ChainEvent) event.Subscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeChainEvent", arg0)
	ret0, _ := ret[0].(event.Subscription)
	return ret0
}

// SubscribeChainEvent indicates an expected call of SubscribeChainEvent
func (mr *MockBackendMockRecorder) SubscribeChainEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeChainEvent", reflect.TypeOf((*MockBackend)(nil).SubscribeChainEvent), arg0)
}

// SubscribeLogsEvent mocks base method
func (m *MockBackend) SubscribeLogsEvent(arg0 chan<- []*types.Log) event.Subscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeLogsEvent", arg0)
	ret0, _ := ret[0].(event.Subscription)
	return ret0
}

// SubscribeLogsEvent indicates an expected call of SubscribeLogsEvent
func (mr *MockBackendMockRecorder) SubscribeLogsEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeLogsEvent", reflect.TypeOf((*MockBackend)(nil).SubscribeLogsEvent), arg0)
}

// SubscribeNewTxsEvent mocks base method
func (m *MockBackend) SubscribeNewTxsEvent(arg0 chan<- blockchain.NewTxsEvent) event.Subscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeNewTxsEvent", arg0)
	ret0, _ := ret[0].(event.Subscription)
	return ret0
}

// SubscribeNewTxsEvent indicates an expected call of SubscribeNewTxsEvent
func (mr *MockBackendMockRecorder) SubscribeNewTxsEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeNewTxsEvent", reflect.TypeOf((*MockBackend)(nil).SubscribeNewTxsEvent), arg0)
}

// SubscribeRemovedLogsEvent mocks base method
func (m *MockBackend) SubscribeRemovedLogsEvent(arg0 chan<- blockchain.RemovedLogsEvent) event.Subscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeRemovedLogsEvent", arg0)
	ret0, _ := ret[0].(event.Subscription)
	return ret0
}

// SubscribeRemovedLogsEvent indicates an expected call of SubscribeRemovedLogsEvent
func (mr *MockBackendMockRecorder) SubscribeRemovedLogsEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeRemovedLogsEvent", reflect.TypeOf((*MockBackend)(nil).SubscribeRemovedLogsEvent), arg0)
}