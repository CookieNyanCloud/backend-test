// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	domain "github.com/cookienyancloud/avito-backend-test/internal/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockIService is a mock of IService interface.
type MockIService struct {
	ctrl     *gomock.Controller
	recorder *MockIServiceMockRecorder
}

// MockIServiceMockRecorder is the mock recorder for MockIService.
type MockIServiceMockRecorder struct {
	mock *MockIService
}

// NewMockIService creates a new mock instance.
func NewMockIService(ctrl *gomock.Controller) *MockIService {
	mock := &MockIService{ctrl: ctrl}
	mock.recorder = &MockIServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIService) EXPECT() *MockIServiceMockRecorder {
	return m.recorder
}

// GetBalance mocks base method.
func (m *MockIService) GetBalance(ctx context.Context, inp *domain.BalanceInput) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", ctx, inp)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockIServiceMockRecorder) GetBalance(ctx, inp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockIService)(nil).GetBalance), ctx, inp)
}

// GetTransactionsList mocks base method.
func (m *MockIService) GetTransactionsList(ctx context.Context, inp *domain.TransactionsListInput) ([]*domain.TransactionsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionsList", ctx, inp)
	ret0, _ := ret[0].([]*domain.TransactionsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionsList indicates an expected call of GetTransactionsList.
func (mr *MockIServiceMockRecorder) GetTransactionsList(ctx, inp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionsList", reflect.TypeOf((*MockIService)(nil).GetTransactionsList), ctx, inp)
}

// MakeRemittance mocks base method.
func (m *MockIService) MakeRemittance(ctx context.Context, inp *domain.RemittanceInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeRemittance", ctx, inp)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeRemittance indicates an expected call of MakeRemittance.
func (mr *MockIServiceMockRecorder) MakeRemittance(ctx, inp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeRemittance", reflect.TypeOf((*MockIService)(nil).MakeRemittance), ctx, inp)
}

// MakeTransaction mocks base method.
func (m *MockIService) MakeTransaction(ctx context.Context, inp *domain.TransactionInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeTransaction", ctx, inp)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeTransaction indicates an expected call of MakeTransaction.
func (mr *MockIServiceMockRecorder) MakeTransaction(ctx, inp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeTransaction", reflect.TypeOf((*MockIService)(nil).MakeTransaction), ctx, inp)
}
