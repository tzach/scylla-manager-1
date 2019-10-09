// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/scylladb/mermaid/restapi (interfaces: BackupService)

// Package restapi is a generated GoMock package.
package restapi

import (
	context "context"
	json "encoding/json"
	gomock "github.com/golang/mock/gomock"
	backup "github.com/scylladb/mermaid/service/backup"
	uuid "github.com/scylladb/mermaid/uuid"
	reflect "reflect"
)

// MockBackupService is a mock of BackupService interface
type MockBackupService struct {
	ctrl     *gomock.Controller
	recorder *MockBackupServiceMockRecorder
}

// MockBackupServiceMockRecorder is the mock recorder for MockBackupService
type MockBackupServiceMockRecorder struct {
	mock *MockBackupService
}

// NewMockBackupService creates a new mock instance
func NewMockBackupService(ctrl *gomock.Controller) *MockBackupService {
	mock := &MockBackupService{ctrl: ctrl}
	mock.recorder = &MockBackupServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBackupService) EXPECT() *MockBackupServiceMockRecorder {
	return m.recorder
}

// GetProgress mocks base method
func (m *MockBackupService) GetProgress(arg0 context.Context, arg1, arg2, arg3 uuid.UUID) (backup.Progress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProgress", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(backup.Progress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProgress indicates an expected call of GetProgress
func (mr *MockBackupServiceMockRecorder) GetProgress(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProgress", reflect.TypeOf((*MockBackupService)(nil).GetProgress), arg0, arg1, arg2, arg3)
}

// GetTarget mocks base method
func (m *MockBackupService) GetTarget(arg0 context.Context, arg1 uuid.UUID, arg2 json.RawMessage, arg3 bool) (backup.Target, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTarget", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(backup.Target)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTarget indicates an expected call of GetTarget
func (mr *MockBackupServiceMockRecorder) GetTarget(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTarget", reflect.TypeOf((*MockBackupService)(nil).GetTarget), arg0, arg1, arg2, arg3)
}

// GetTargetSize mocks base method
func (m *MockBackupService) GetTargetSize(arg0 context.Context, arg1 uuid.UUID, arg2 backup.Target) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTargetSize", arg0, arg1, arg2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTargetSize indicates an expected call of GetTargetSize
func (mr *MockBackupServiceMockRecorder) GetTargetSize(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTargetSize", reflect.TypeOf((*MockBackupService)(nil).GetTargetSize), arg0, arg1, arg2)
}

// List mocks base method
func (m *MockBackupService) List(arg0 context.Context, arg1 uuid.UUID, arg2 string, arg3 []backup.Location, arg4 backup.ListFilter) ([]backup.ListItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]backup.ListItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockBackupServiceMockRecorder) List(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockBackupService)(nil).List), arg0, arg1, arg2, arg3, arg4)
}
