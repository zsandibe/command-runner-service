// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/zsandibe/command-runner-service/internal/domain"
	entity "github.com/zsandibe/command-runner-service/internal/entity"
)

// MockCommand is a mock of Command interface.
type MockCommand struct {
	ctrl     *gomock.Controller
	recorder *MockCommandMockRecorder
}

// MockCommandMockRecorder is the mock recorder for MockCommand.
type MockCommandMockRecorder struct {
	mock *MockCommand
}

// NewMockCommand creates a new mock instance.
func NewMockCommand(ctrl *gomock.Controller) *MockCommand {
	mock := &MockCommand{ctrl: ctrl}
	mock.recorder = &MockCommandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommand) EXPECT() *MockCommandMockRecorder {
	return m.recorder
}

// CreateCommand mocks base method.
func (m *MockCommand) CreateCommand(ctx context.Context, inp domain.CreateCommandInput) (*entity.Command, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCommand", ctx, inp)
	ret0, _ := ret[0].(*entity.Command)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCommand indicates an expected call of CreateCommand.
func (mr *MockCommandMockRecorder) CreateCommand(ctx, inp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCommand", reflect.TypeOf((*MockCommand)(nil).CreateCommand), ctx, inp)
}

// CreateCommandOutput mocks base method.
func (m *MockCommand) CreateCommandOutput(ctx context.Context, id int64, output string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCommandOutput", ctx, id, output)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCommandOutput indicates an expected call of CreateCommandOutput.
func (mr *MockCommandMockRecorder) CreateCommandOutput(ctx, id, output interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCommandOutput", reflect.TypeOf((*MockCommand)(nil).CreateCommandOutput), ctx, id, output)
}

// DeleteCommandById mocks base method.
func (m *MockCommand) DeleteCommandById(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCommandById", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCommandById indicates an expected call of DeleteCommandById.
func (mr *MockCommandMockRecorder) DeleteCommandById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCommandById", reflect.TypeOf((*MockCommand)(nil).DeleteCommandById), ctx, id)
}

// GetAllCommands mocks base method.
func (m *MockCommand) GetAllCommands(ctx context.Context) (*[]entity.Command, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCommands", ctx)
	ret0, _ := ret[0].(*[]entity.Command)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCommands indicates an expected call of GetAllCommands.
func (mr *MockCommandMockRecorder) GetAllCommands(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCommands", reflect.TypeOf((*MockCommand)(nil).GetAllCommands), ctx)
}

// GetCommandById mocks base method.
func (m *MockCommand) GetCommandById(ctx context.Context, id int64) (*entity.Command, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommandById", ctx, id)
	ret0, _ := ret[0].(*entity.Command)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommandById indicates an expected call of GetCommandById.
func (mr *MockCommandMockRecorder) GetCommandById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommandById", reflect.TypeOf((*MockCommand)(nil).GetCommandById), ctx, id)
}

// GetCommands mocks base method.
func (m *MockCommand) GetCommands(ctx context.Context, ids []int64) (*[]entity.Command, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommands", ctx, ids)
	ret0, _ := ret[0].(*[]entity.Command)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommands indicates an expected call of GetCommands.
func (mr *MockCommandMockRecorder) GetCommands(ctx, ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommands", reflect.TypeOf((*MockCommand)(nil).GetCommands), ctx, ids)
}

// MockCache is a mock of Cache interface.
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
}

// MockCacheMockRecorder is the mock recorder for MockCache.
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance.
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockCache) Delete(key int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCacheMockRecorder) Delete(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCache)(nil).Delete), key)
}

// Get mocks base method.
func (m *MockCache) Get(key int64) (any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCacheMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCache)(nil).Get), key)
}

// GetAllKeys mocks base method.
func (m *MockCache) GetAllKeys() ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllKeys")
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllKeys indicates an expected call of GetAllKeys.
func (mr *MockCacheMockRecorder) GetAllKeys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllKeys", reflect.TypeOf((*MockCache)(nil).GetAllKeys))
}

// GetLen mocks base method.
func (m *MockCache) GetLen() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLen")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLen indicates an expected call of GetLen.
func (mr *MockCacheMockRecorder) GetLen() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLen", reflect.TypeOf((*MockCache)(nil).GetLen))
}

// Set mocks base method.
func (m *MockCache) Set(key int64, value any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockCacheMockRecorder) Set(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockCache)(nil).Set), key, value)
}
