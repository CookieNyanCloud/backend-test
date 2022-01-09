// Code generated by MockGen. DO NOT EDIT.
// Source: redisCache.go

// Package mock_redis is a generated GoMock package.
package mock_redis

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockICache is a mock of ICache interface.
type MockICache struct {
	ctrl     *gomock.Controller
	recorder *MockICacheMockRecorder
}

// MockICacheMockRecorder is the mock recorder for MockICache.
type MockICacheMockRecorder struct {
	mock *MockICache
}

// NewMockICache creates a new mock instance.
func NewMockICache(ctrl *gomock.Controller) *MockICache {
	mock := &MockICache{ctrl: ctrl}
	mock.recorder = &MockICacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICache) EXPECT() *MockICacheMockRecorder {
	return m.recorder
}

// CacheKey mocks base method.
func (m *MockICache) CacheKey(ctx context.Context, key uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CacheKey", ctx, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// CacheKey indicates an expected call of CacheKey.
func (mr *MockICacheMockRecorder) CacheKey(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CacheKey", reflect.TypeOf((*MockICache)(nil).CacheKey), ctx, key)
}

// CheckKey mocks base method.
func (m *MockICache) CheckKey(ctx context.Context, key uuid.UUID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckKey", ctx, key)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckKey indicates an expected call of CheckKey.
func (mr *MockICacheMockRecorder) CheckKey(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckKey", reflect.TypeOf((*MockICache)(nil).CheckKey), ctx, key)
}