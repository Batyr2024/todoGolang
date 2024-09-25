// Code generated by MockGen. DO NOT EDIT.
// Source: ./db/sqlc/querier.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	sqlc "github.com/Batyr2024/todoGolang/db/sqlc"
	. "github.com/Batyr2024/todoGolang/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockQuerier is a mock of Querier interface.
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier.
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance.
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// ChangeCheckedAll mocks base method.
func (m *MockQuerier) ChangeCheckedAll(ctx context.Context, arg *sqlc.ChangeCheckedAllParams) ([]int32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeCheckedAll", ctx, arg)
	ret0, _ := ret[0].([]int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeCheckedAll indicates an expected call of ChangeCheckedAll.
func (mr *MockQuerierMockRecorder) ChangeCheckedAll(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeCheckedAll", reflect.TypeOf((*MockQuerier)(nil).ChangeCheckedAll), ctx, arg)
}

// ChangeCheckedByID mocks base method.
func (m *MockQuerier) ChangeCheckedByID(ctx context.Context, arg *sqlc.ChangeCheckedByIDParams) (*Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeCheckedByID", ctx, arg)
	ret0, _ := ret[0].(*Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeCheckedByID indicates an expected call of ChangeCheckedByID.
func (mr *MockQuerierMockRecorder) ChangeCheckedByID(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeCheckedByID", reflect.TypeOf((*MockQuerier)(nil).ChangeCheckedByID), ctx, arg)
}

// ChangeTextByID mocks base method.
func (m *MockQuerier) ChangeTextByID(ctx context.Context, arg *sqlc.ChangeTextByIDParams) (*Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeTextByID", ctx, arg)
	ret0, _ := ret[0].(*Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeTextByID indicates an expected call of ChangeTextByID.
func (mr *MockQuerierMockRecorder) ChangeTextByID(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeTextByID", reflect.TypeOf((*MockQuerier)(nil).ChangeTextByID), ctx, arg)
}

// Create mocks base method.
func (m *MockQuerier) Create(ctx context.Context, arg *sqlc.CreateParams) (*Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, arg)
	ret0, _ := ret[0].(*Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockQuerierMockRecorder) Create(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockQuerier)(nil).Create), ctx, arg)
}

// DeleteAll mocks base method.
func (m *MockQuerier) DeleteAll(ctx context.Context) ([]int32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx)
	ret0, _ := ret[0].([]int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockQuerierMockRecorder) DeleteAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockQuerier)(nil).DeleteAll), ctx)
}

// DeleteByID mocks base method.
func (m *MockQuerier) DeleteByID(ctx context.Context, arg *sqlc.DeleteByIDParams) (*Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, arg)
	ret0, _ := ret[0].(*Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockQuerierMockRecorder) DeleteByID(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockQuerier)(nil).DeleteByID), ctx, arg)
}

// GetAll mocks base method.
func (m *MockQuerier) GetAll(ctx context.Context) ([]*Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]*Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockQuerierMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockQuerier)(nil).GetAll), ctx)
}
