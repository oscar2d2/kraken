// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/uber/kraken/build-index/tagstore (interfaces: FileStore)

package mocktagstore

import (
	gomock "github.com/golang/mock/gomock"
	base "github.com/uber/kraken/lib/store/base"
	metadata "github.com/uber/kraken/lib/store/metadata"
	io "io"
	reflect "reflect"
)

// MockFileStore is a mock of FileStore interface
type MockFileStore struct {
	ctrl     *gomock.Controller
	recorder *MockFileStoreMockRecorder
}

// MockFileStoreMockRecorder is the mock recorder for MockFileStore
type MockFileStoreMockRecorder struct {
	mock *MockFileStore
}

// NewMockFileStore creates a new mock instance
func NewMockFileStore(ctrl *gomock.Controller) *MockFileStore {
	mock := &MockFileStore{ctrl: ctrl}
	mock.recorder = &MockFileStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockFileStore) EXPECT() *MockFileStoreMockRecorder {
	return _m.recorder
}

// CreateCacheFile mocks base method
func (_m *MockFileStore) CreateCacheFile(_param0 string, _param1 io.Reader) error {
	ret := _m.ctrl.Call(_m, "CreateCacheFile", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCacheFile indicates an expected call of CreateCacheFile
func (_mr *MockFileStoreMockRecorder) CreateCacheFile(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CreateCacheFile", reflect.TypeOf((*MockFileStore)(nil).CreateCacheFile), arg0, arg1)
}

// GetCacheFileReader mocks base method
func (_m *MockFileStore) GetCacheFileReader(_param0 string) (base.FileReader, error) {
	ret := _m.ctrl.Call(_m, "GetCacheFileReader", _param0)
	ret0, _ := ret[0].(base.FileReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCacheFileReader indicates an expected call of GetCacheFileReader
func (_mr *MockFileStoreMockRecorder) GetCacheFileReader(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetCacheFileReader", reflect.TypeOf((*MockFileStore)(nil).GetCacheFileReader), arg0)
}

// SetCacheFileMetadata mocks base method
func (_m *MockFileStore) SetCacheFileMetadata(_param0 string, _param1 metadata.Metadata) (bool, error) {
	ret := _m.ctrl.Call(_m, "SetCacheFileMetadata", _param0, _param1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetCacheFileMetadata indicates an expected call of SetCacheFileMetadata
func (_mr *MockFileStoreMockRecorder) SetCacheFileMetadata(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetCacheFileMetadata", reflect.TypeOf((*MockFileStore)(nil).SetCacheFileMetadata), arg0, arg1)
}
