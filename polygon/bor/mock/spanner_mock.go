// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ledgerwatch/erigon/polygon/bor (interfaces: Spanner)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	common "github.com/ledgerwatch/erigon-lib/common"
	consensus "github.com/ledgerwatch/erigon/consensus"
	valset "github.com/ledgerwatch/erigon/polygon/bor/valset"
	span "github.com/ledgerwatch/erigon/polygon/heimdall/span"
)

// MockSpanner is a mock of Spanner interface.
type MockSpanner struct {
	ctrl     *gomock.Controller
	recorder *MockSpannerMockRecorder
}

// MockSpannerMockRecorder is the mock recorder for MockSpanner.
type MockSpannerMockRecorder struct {
	mock *MockSpanner
}

// NewMockSpanner creates a new mock instance.
func NewMockSpanner(ctrl *gomock.Controller) *MockSpanner {
	mock := &MockSpanner{ctrl: ctrl}
	mock.recorder = &MockSpannerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpanner) EXPECT() *MockSpannerMockRecorder {
	return m.recorder
}

// CommitSpan mocks base method.
func (m *MockSpanner) CommitSpan(arg0 span.HeimdallSpan, arg1 consensus.SystemCall) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommitSpan", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CommitSpan indicates an expected call of CommitSpan.
func (mr *MockSpannerMockRecorder) CommitSpan(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitSpan", reflect.TypeOf((*MockSpanner)(nil).CommitSpan), arg0, arg1)
}

// GetCurrentProducers mocks base method.
func (m *MockSpanner) GetCurrentProducers(arg0 uint64, arg1 common.Address, arg2 consensus.ChainHeaderReader) ([]*valset.Validator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentProducers", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*valset.Validator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentProducers indicates an expected call of GetCurrentProducers.
func (mr *MockSpannerMockRecorder) GetCurrentProducers(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentProducers", reflect.TypeOf((*MockSpanner)(nil).GetCurrentProducers), arg0, arg1, arg2)
}

// GetCurrentSpan mocks base method.
func (m *MockSpanner) GetCurrentSpan(arg0 consensus.SystemCall) (*span.Span, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentSpan", arg0)
	ret0, _ := ret[0].(*span.Span)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentSpan indicates an expected call of GetCurrentSpan.
func (mr *MockSpannerMockRecorder) GetCurrentSpan(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentSpan", reflect.TypeOf((*MockSpanner)(nil).GetCurrentSpan), arg0)
}

// GetCurrentValidators mocks base method.
func (m *MockSpanner) GetCurrentValidators(arg0 uint64, arg1 common.Address, arg2 consensus.ChainHeaderReader) ([]*valset.Validator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentValidators", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*valset.Validator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentValidators indicates an expected call of GetCurrentValidators.
func (mr *MockSpannerMockRecorder) GetCurrentValidators(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentValidators", reflect.TypeOf((*MockSpanner)(nil).GetCurrentValidators), arg0, arg1, arg2)
}
