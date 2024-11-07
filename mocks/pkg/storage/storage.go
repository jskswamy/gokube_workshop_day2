// Code generated by MockGen. DO NOT EDIT.
// Source: storage.go
//
// Generated by this command:
//
//	mockgen -destination /Users/subramk/source/github.com/unmeshjoshi/gokube/mocks/pkg/storage/storage.go -package storage -source storage.go
//

// Package storage is a generated GoMock package.
package storage

import (
	context "context"
	runtime "gokube/pkg/runtime"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
	isgomock struct{}
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockStorage) Create(ctx context.Context, key string, obj runtime.Object) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, key, obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockStorageMockRecorder) Create(ctx, key, obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStorage)(nil).Create), ctx, key, obj)
}

// Delete mocks base method.
func (m *MockStorage) Delete(ctx context.Context, key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockStorageMockRecorder) Delete(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStorage)(nil).Delete), ctx, key)
}

// DeletePrefix mocks base method.
func (m *MockStorage) DeletePrefix(ctx context.Context, prefix string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePrefix", ctx, prefix)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePrefix indicates an expected call of DeletePrefix.
func (mr *MockStorageMockRecorder) DeletePrefix(ctx, prefix any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePrefix", reflect.TypeOf((*MockStorage)(nil).DeletePrefix), ctx, prefix)
}

// Get mocks base method.
func (m *MockStorage) Get(ctx context.Context, key string, obj runtime.Object) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key, obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockStorageMockRecorder) Get(ctx, key, obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStorage)(nil).Get), ctx, key, obj)
}

// List mocks base method.
func (m *MockStorage) List(ctx context.Context, prefix string, listObj any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, prefix, listObj)
	ret0, _ := ret[0].(error)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockStorageMockRecorder) List(ctx, prefix, listObj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStorage)(nil).List), ctx, prefix, listObj)
}

// Update mocks base method.
func (m *MockStorage) Update(ctx context.Context, key string, obj runtime.Object) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, key, obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockStorageMockRecorder) Update(ctx, key, obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStorage)(nil).Update), ctx, key, obj)
}