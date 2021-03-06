// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	domain "github.com/cookienyancloud/avito-backend-test/internal/domain"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockIRepo is a mock of IRepo interface.
type MockIRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIRepoMockRecorder
}

// MockIRepoMockRecorder is the mock recorder for MockIRepo.
type MockIRepoMockRecorder struct {
	mock *MockIRepo
}

// NewMockIRepo creates a new mock instance.
func NewMockIRepo(ctrl *gomock.Controller) *MockIRepo {
	mock := &MockIRepo{ctrl: ctrl}
	mock.recorder = &MockIRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRepo) EXPECT() *MockIRepoMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockIRepo) Close(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockIRepoMockRecorder) Close(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIRepo)(nil).Close), ctx)
}

// CreateNewTransaction mocks base method.
func (m *MockIRepo) CreateNewTransaction(ctx context.Context, idFrom uuid.UUID, operation string, sum float64, idTo uuid.UUID, description string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewTransaction", ctx, idFrom, operation, sum, idTo, description)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNewTransaction indicates an expected call of CreateNewTransaction.
func (mr *MockIRepoMockRecorder) CreateNewTransaction(ctx, idFrom, operation, sum, idTo, description interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewTransaction", reflect.TypeOf((*MockIRepo)(nil).CreateNewTransaction), ctx, idFrom, operation, sum, idTo, description)
}

// GetBalance mocks base method.
func (m *MockIRepo) GetBalance(ctx context.Context, inp *domain.BalanceInput) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", ctx, inp)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockIRepoMockRecorder) GetBalance(ctx, inp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockIRepo)(nil).GetBalance), ctx, inp)
}

// GetTransactionsList mocks base method.
func (m *MockIRepo) GetTransactionsList(ctx context.Context, inp *domain.TransactionsListInput) ([]domain.TransactionsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionsList", ctx, inp)
	ret0, _ := ret[0].([]domain.TransactionsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionsList indicates an expected call of GetTransactionsList.
func (mr *MockIRepoMockRecorder) GetTransactionsList(ctx, inp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionsList", reflect.TypeOf((*MockIRepo)(nil).GetTransactionsList), ctx, inp)
}

// MakeRemittance mocks base method.
func (m *MockIRepo) MakeRemittance(ctx context.Context, inp *domain.RemittanceInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeRemittance", ctx, inp)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeRemittance indicates an expected call of MakeRemittance.
func (mr *MockIRepoMockRecorder) MakeRemittance(ctx, inp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeRemittance", reflect.TypeOf((*MockIRepo)(nil).MakeRemittance), ctx, inp)
}

// MakeTransaction mocks base method.
func (m *MockIRepo) MakeTransaction(ctx context.Context, inp *domain.TransactionInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeTransaction", ctx, inp)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeTransaction indicates an expected call of MakeTransaction.
func (mr *MockIRepoMockRecorder) MakeTransaction(ctx, inp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeTransaction", reflect.TypeOf((*MockIRepo)(nil).MakeTransaction), ctx, inp)
}
