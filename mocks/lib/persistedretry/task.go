// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/uber/kraken/lib/persistedretry (interfaces: Task)

package mockpersistedretry

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockTask is a mock of Task interface
type MockTask struct {
	ctrl     *gomock.Controller
	recorder *MockTaskMockRecorder
}

// MockTaskMockRecorder is the mock recorder for MockTask
type MockTaskMockRecorder struct {
	mock *MockTask
}

// NewMockTask creates a new mock instance
func NewMockTask(ctrl *gomock.Controller) *MockTask {
	mock := &MockTask{ctrl: ctrl}
	mock.recorder = &MockTaskMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockTask) EXPECT() *MockTaskMockRecorder {
	return _m.recorder
}

// GetFailures mocks base method
func (_m *MockTask) GetFailures() int {
	ret := _m.ctrl.Call(_m, "GetFailures")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetFailures indicates an expected call of GetFailures
func (_mr *MockTaskMockRecorder) GetFailures() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetFailures", reflect.TypeOf((*MockTask)(nil).GetFailures))
}

// GetLastAttempt mocks base method
func (_m *MockTask) GetLastAttempt() time.Time {
	ret := _m.ctrl.Call(_m, "GetLastAttempt")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetLastAttempt indicates an expected call of GetLastAttempt
func (_mr *MockTaskMockRecorder) GetLastAttempt() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetLastAttempt", reflect.TypeOf((*MockTask)(nil).GetLastAttempt))
}

// Ready mocks base method
func (_m *MockTask) Ready() bool {
	ret := _m.ctrl.Call(_m, "Ready")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Ready indicates an expected call of Ready
func (_mr *MockTaskMockRecorder) Ready() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Ready", reflect.TypeOf((*MockTask)(nil).Ready))
}
