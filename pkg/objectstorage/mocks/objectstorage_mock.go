// Code generated by MockGen. DO NOT EDIT.
// Source: objectstorage.go

// Package mocks is a generated GoMock package.
package mocks

import (
	io "io"
	reflect "reflect"

	objectstorage "d7y.io/dragonfly/v2/pkg/objectstorage"
	gomock "github.com/golang/mock/gomock"
)

// MockObjectStorage is a mock of ObjectStorage interface.
type MockObjectStorage struct {
	ctrl     *gomock.Controller
	recorder *MockObjectStorageMockRecorder
}

// MockObjectStorageMockRecorder is the mock recorder for MockObjectStorage.
type MockObjectStorageMockRecorder struct {
	mock *MockObjectStorage
}

// NewMockObjectStorage creates a new mock instance.
func NewMockObjectStorage(ctrl *gomock.Controller) *MockObjectStorage {
	mock := &MockObjectStorage{ctrl: ctrl}
	mock.recorder = &MockObjectStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockObjectStorage) EXPECT() *MockObjectStorageMockRecorder {
	return m.recorder
}

// CreateBucket mocks base method.
func (m *MockObjectStorage) CreateBucket(bucketName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBucket", bucketName)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBucket indicates an expected call of CreateBucket.
func (mr *MockObjectStorageMockRecorder) CreateBucket(bucketName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBucket", reflect.TypeOf((*MockObjectStorage)(nil).CreateBucket), bucketName)
}

// CreateObject mocks base method.
func (m *MockObjectStorage) CreateObject(bucketName, objectKey string, reader io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateObject", bucketName, objectKey, reader)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateObject indicates an expected call of CreateObject.
func (mr *MockObjectStorageMockRecorder) CreateObject(bucketName, objectKey, reader interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateObject", reflect.TypeOf((*MockObjectStorage)(nil).CreateObject), bucketName, objectKey, reader)
}

// DeleteBucket mocks base method.
func (m *MockObjectStorage) DeleteBucket(bucketName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBucket", bucketName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBucket indicates an expected call of DeleteBucket.
func (mr *MockObjectStorageMockRecorder) DeleteBucket(bucketName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBucket", reflect.TypeOf((*MockObjectStorage)(nil).DeleteBucket), bucketName)
}

// DeleteObject mocks base method.
func (m *MockObjectStorage) DeleteObject(bucketName, objectKey string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObject", bucketName, objectKey)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteObject indicates an expected call of DeleteObject.
func (mr *MockObjectStorageMockRecorder) DeleteObject(bucketName, objectKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObject", reflect.TypeOf((*MockObjectStorage)(nil).DeleteObject), bucketName, objectKey)
}

// GetBucketMetadata mocks base method.
func (m *MockObjectStorage) GetBucketMetadata(bucketName string) (*objectstorage.BucketMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBucketMetadata", bucketName)
	ret0, _ := ret[0].(*objectstorage.BucketMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketMetadata indicates an expected call of GetBucketMetadata.
func (mr *MockObjectStorageMockRecorder) GetBucketMetadata(bucketName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketMetadata", reflect.TypeOf((*MockObjectStorage)(nil).GetBucketMetadata), bucketName)
}

// GetObjectMetadata mocks base method.
func (m *MockObjectStorage) GetObjectMetadata(bucketName, objectKey string) (*objectstorage.ObjectMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjectMetadata", bucketName, objectKey)
	ret0, _ := ret[0].(*objectstorage.ObjectMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjectMetadata indicates an expected call of GetObjectMetadata.
func (mr *MockObjectStorageMockRecorder) GetObjectMetadata(bucketName, objectKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectMetadata", reflect.TypeOf((*MockObjectStorage)(nil).GetObjectMetadata), bucketName, objectKey)
}

// GetOject mocks base method.
func (m *MockObjectStorage) GetOject(bucketName, objectKey string) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOject", bucketName, objectKey)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOject indicates an expected call of GetOject.
func (mr *MockObjectStorageMockRecorder) GetOject(bucketName, objectKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOject", reflect.TypeOf((*MockObjectStorage)(nil).GetOject), bucketName, objectKey)
}

// ListBucketMetadatas mocks base method.
func (m *MockObjectStorage) ListBucketMetadatas() ([]*objectstorage.BucketMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBucketMetadatas")
	ret0, _ := ret[0].([]*objectstorage.BucketMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBucketMetadatas indicates an expected call of ListBucketMetadatas.
func (mr *MockObjectStorageMockRecorder) ListBucketMetadatas() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBucketMetadatas", reflect.TypeOf((*MockObjectStorage)(nil).ListBucketMetadatas))
}

// ListObjectMetadatas mocks base method.
func (m *MockObjectStorage) ListObjectMetadatas(bucketName, prefix, marker string, limit int64) ([]*objectstorage.ObjectMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListObjectMetadatas", bucketName, prefix, marker, limit)
	ret0, _ := ret[0].([]*objectstorage.ObjectMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListObjectMetadatas indicates an expected call of ListObjectMetadatas.
func (mr *MockObjectStorageMockRecorder) ListObjectMetadatas(bucketName, prefix, marker, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObjectMetadatas", reflect.TypeOf((*MockObjectStorage)(nil).ListObjectMetadatas), bucketName, prefix, marker, limit)
}