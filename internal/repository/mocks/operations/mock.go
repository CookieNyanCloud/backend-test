// Code generated by MockGen. DO NOT EDIT.
// Source: operations.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	domain "github.com/cookienyancloud/avito-backend-test/internal/domain"
	repository "github.com/cookienyancloud/avito-backend-test/internal/repository"
	gomock "github.com/golang/mock/gomock"
)

// MockFinanceOperations is a mock of FinanceOperations interface.
type MockFinanceOperations struct {
	ctrl     *gomock.Controller
	recorder *MockFinanceOperationsMockRecorder
}

// MockFinanceOperationsMockRecorder is the mock recorder for MockFinanceOperations.
type MockFinanceOperationsMockRecorder struct {
	mock *MockFinanceOperations
}

// NewMockFinanceOperations creates a new mock instance.
func NewMockFinanceOperations(ctrl *gomock.Controller) *MockFinanceOperations {
	mock := &MockFinanceOperations{ctrl: ctrl}
	mock.recorder = &MockFinanceOperationsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFinanceOperations) EXPECT() *MockFinanceOperationsMockRecorder {
	return m.recorder
}

// GetBalance mocks base method.
func (m *MockFinanceOperations) GetBalance(ctx context.Context, inp *domain.BalanceInput) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", ctx, inp)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockFinanceOperationsMockRecorder) GetBalance(ctx, inp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockFinanceOperations)(nil).GetBalance), ctx, inp)
}

// GetTransactionsList mocks base method.
func (m *MockFinanceOperations) GetTransactionsList(ctx context.Context, inp *domain.TransactionsListInput) ([]repository.TransactionsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionsList", ctx, inp)
	ret0, _ := ret[0].([]repository.TransactionsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionsList indicates an expected call of GetTransactionsList.
func (mr *MockFinanceOperationsMockRecorder) GetTransactionsList(ctx, inp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionsList", reflect.TypeOf((*MockFinanceOperations)(nil).GetTransactionsList), ctx, inp)
}

// MakeRemittance mocks base method.
func (m *MockFinanceOperations) MakeRemittance(ctx context.Context, inp *domain.RemittanceInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeRemittance", ctx, inp)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeRemittance indicates an expected call of MakeRemittance.
func (mr *MockFinanceOperationsMockRecorder) MakeRemittance(ctx, inp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeRemittance", reflect.TypeOf((*MockFinanceOperations)(nil).MakeRemittance), ctx, inp)
}

// MakeTransaction mocks base method.
func (m *MockFinanceOperations) MakeTransaction(ctx context.Context, inp *domain.TransactionInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeTransaction", ctx, inp)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeTransaction indicates an expected call of MakeTransaction.
func (mr *MockFinanceOperationsMockRecorder) MakeTransaction(ctx, inp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeTransaction", reflect.TypeOf((*MockFinanceOperations)(nil).MakeTransaction), ctx, inp)
}