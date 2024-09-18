// Code generated by MockGen. DO NOT EDIT.
// Source: task.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	domain "github.com/Batyr2024/todoGolang/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockinterfaceHandler is a mock of interfaceHandler interface.
type MockinterfaceHandler struct {
	ctrl     *gomock.Controller
	recorder *MockinterfaceHandlerMockRecorder
}

// MockinterfaceHandlerMockRecorder is the mock recorder for MockinterfaceHandler.
type MockinterfaceHandlerMockRecorder struct {
	mock *MockinterfaceHandler
}

// NewMockinterfaceHandler creates a new mock instance.
func NewMockinterfaceHandler(ctrl *gomock.Controller) *MockinterfaceHandler {
	mock := &MockinterfaceHandler{ctrl: ctrl}
	mock.recorder = &MockinterfaceHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockinterfaceHandler) EXPECT() *MockinterfaceHandlerMockRecorder {
	return m.recorder
}

// ChangeCheckedAll mocks base method.
func (m *MockinterfaceHandler) ChangeCheckedAll(ctx context.Context, checked bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeCheckedAll", ctx, checked)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeCheckedAll indicates an expected call of ChangeCheckedAll.
func (mr *MockinterfaceHandlerMockRecorder) ChangeCheckedAll(ctx, checked interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeCheckedAll", reflect.TypeOf((*MockinterfaceHandler)(nil).ChangeCheckedAll), ctx, checked)
}

// ChangeCheckedByID mocks base method.
func (m *MockinterfaceHandler) ChangeCheckedByID(ctx context.Context, id int, checked bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeCheckedByID", ctx, id, checked)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeCheckedByID indicates an expected call of ChangeCheckedByID.
func (mr *MockinterfaceHandlerMockRecorder) ChangeCheckedByID(ctx, id, checked interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeCheckedByID", reflect.TypeOf((*MockinterfaceHandler)(nil).ChangeCheckedByID), ctx, id, checked)
}

// ChangeText mocks base method.
func (m *MockinterfaceHandler) ChangeText(ctx context.Context, id int, text string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeText", ctx, id, text)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeText indicates an expected call of ChangeText.
func (mr *MockinterfaceHandlerMockRecorder) ChangeText(ctx, id, text interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeText", reflect.TypeOf((*MockinterfaceHandler)(nil).ChangeText), ctx, id, text)
}

// Create mocks base method.
func (m *MockinterfaceHandler) Create(ctx context.Context, task domain.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, task)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockinterfaceHandlerMockRecorder) Create(ctx, task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockinterfaceHandler)(nil).Create), ctx, task)
}

// DeleteAll mocks base method.
func (m *MockinterfaceHandler) DeleteAll(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockinterfaceHandlerMockRecorder) DeleteAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockinterfaceHandler)(nil).DeleteAll), ctx)
}

// DeleteByID mocks base method.
func (m *MockinterfaceHandler) DeleteByID(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockinterfaceHandlerMockRecorder) DeleteByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockinterfaceHandler)(nil).DeleteByID), ctx, id)
}

// FindAll mocks base method.
func (m *MockinterfaceHandler) FindAll(ctx context.Context) ([]*domain.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx)
	ret0, _ := ret[0].([]*domain.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockinterfaceHandlerMockRecorder) FindAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockinterfaceHandler)(nil).FindAll), ctx)
}
