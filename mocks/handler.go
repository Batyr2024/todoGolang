// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/api/handler/task.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
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
func (m *MockinterfaceHandler) ChangeCheckedAll(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ChangeCheckedAll", ctx)
}

// ChangeCheckedAll indicates an expected call of ChangeCheckedAll.
func (mr *MockinterfaceHandlerMockRecorder) ChangeCheckedAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeCheckedAll", reflect.TypeOf((*MockinterfaceHandler)(nil).ChangeCheckedAll), ctx)
}

// ChangeCheckedByID mocks base method.
func (m *MockinterfaceHandler) ChangeCheckedByID(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ChangeCheckedByID", ctx)
}

// ChangeCheckedByID indicates an expected call of ChangeCheckedByID.
func (mr *MockinterfaceHandlerMockRecorder) ChangeCheckedByID(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeCheckedByID", reflect.TypeOf((*MockinterfaceHandler)(nil).ChangeCheckedByID), ctx)
}

// ChangeText mocks base method.
func (m *MockinterfaceHandler) ChangeText(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ChangeText", ctx)
}

// ChangeText indicates an expected call of ChangeText.
func (mr *MockinterfaceHandlerMockRecorder) ChangeText(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeText", reflect.TypeOf((*MockinterfaceHandler)(nil).ChangeText), ctx)
}

// Create mocks base method.
func (m *MockinterfaceHandler) Create(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Create", ctx)
}

// Create indicates an expected call of Create.
func (mr *MockinterfaceHandlerMockRecorder) Create(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockinterfaceHandler)(nil).Create), ctx)
}

// DeleteAll mocks base method.
func (m *MockinterfaceHandler) DeleteAll(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteAll", ctx)
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockinterfaceHandlerMockRecorder) DeleteAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockinterfaceHandler)(nil).DeleteAll), ctx)
}

// DeleteByID mocks base method.
func (m *MockinterfaceHandler) DeleteByID(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteByID", ctx)
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockinterfaceHandlerMockRecorder) DeleteByID(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockinterfaceHandler)(nil).DeleteByID), ctx)
}

// FindAll mocks base method.
func (m *MockinterfaceHandler) FindAll(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FindAll", ctx)
}

// FindAll indicates an expected call of FindAll.
func (mr *MockinterfaceHandlerMockRecorder) FindAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockinterfaceHandler)(nil).FindAll), ctx)
}
