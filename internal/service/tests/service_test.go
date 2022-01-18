package tests

import (
	"context"
	"testing"

	"github.com/cookienyancloud/avito-backend-test/internal/domain"
	service "github.com/cookienyancloud/avito-backend-test/internal/service"
	mock_service "github.com/cookienyancloud/avito-backend-test/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/magiconair/properties/assert"
)

func TestMakeTransaction(t *testing.T) {
	type mockBehavior func(m *mock_service.MockIRepo, inp *domain.TransactionInput)

	tt := []struct {
		name  string
		comm  string
		inp   *domain.TransactionInput
		mockB mockBehavior
		exp   error
	}{
		{
			name: "ok",
			comm: "transaction",
			inp: &domain.TransactionInput{
				Id:          uuid.MustParse("1993f8f2-d580-4fb1-bd8e-5bdfb7ddd7e4"),
				Sum:         10,
				Description: "test",
			},
			mockB: func(r *mock_service.MockIRepo, inp *domain.TransactionInput) {
				r.
					EXPECT().
					MakeTransaction(gomock.Any(), inp).
					Return(nil)
				r.
					EXPECT().
					CreateNewTransaction(gomock.Any(), inp.Id, "transaction", inp.Sum, uuid.Nil, inp.Description).
					Return(nil)
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)

			defer c.Finish()
			dbMain := mock_service.NewMockIRepo(c)
			tc.mockB(dbMain, tc.inp)
			fs := service.NewFinanceService(dbMain)
			err := fs.MakeTransaction(context.Background(), tc.inp)
			assert.Equal(t, err, tc.exp)
		})
	}
}

func TestMakeRemittance(t *testing.T) {
	type mockBehavior func(m *mock_service.MockIRepo, inp *domain.RemittanceInput)

	tt := []struct {
		name  string
		comm  string
		inp   *domain.RemittanceInput
		mockB mockBehavior
		exp   error
	}{
		{
			name: "ok",
			comm: "remittance",
			inp: &domain.RemittanceInput{
				IdFrom:      uuid.MustParse("1993f8f2-d580-4fb1-bd8e-5bdfb7ddd7e4"),
				IdTo:        uuid.MustParse("2993f8f2-d580-4fb1-bd8e-5bdfb7ddd7e4"),
				Sum:         10,
				Description: "test",
			},
			mockB: func(r *mock_service.MockIRepo, inp *domain.RemittanceInput) {
				r.
					EXPECT().
					MakeRemittance(gomock.Any(), inp).
					Return(nil)
				r.
					EXPECT().
					CreateNewTransaction(gomock.Any(), inp.IdFrom, "remittance", inp.Sum, inp.IdTo, inp.Description).
					Return(nil)
			},
			exp: nil,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)

			defer c.Finish()
			dbMain := mock_service.NewMockIRepo(c)
			tc.mockB(dbMain, tc.inp)
			fs := service.NewFinanceService(dbMain)
			err := fs.MakeRemittance(context.Background(), tc.inp)
			assert.Equal(t, err, tc.exp)
		})
	}
}

func TestGetBalance(t *testing.T) {
	type mockBehavior func(m *mock_service.MockIRepo, inp *domain.BalanceInput)

	tt := []struct {
		name   string
		inp    *domain.BalanceInput
		mockB  mockBehavior
		expErr error
		expRes float64
	}{
		{
			name: "ok",
			inp: &domain.BalanceInput{
				Id: uuid.MustParse("1993f8f2-d580-4fb1-bd8e-5bdfb7ddd7e4"),
			},
			mockB: func(r *mock_service.MockIRepo, inp *domain.BalanceInput) {
				r.
					EXPECT().
					GetBalance(gomock.Any(), inp).
					Return(100.1, nil)
			},
			expErr: nil,
			expRes: 100.1,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)

			defer c.Finish()
			dbMain := mock_service.NewMockIRepo(c)
			tc.mockB(dbMain, tc.inp)
			fs := service.NewFinanceService(dbMain)
			res, err := fs.GetBalance(context.Background(), tc.inp)
			assert.Equal(t, err, tc.expErr)
			assert.Equal(t, res, tc.expRes)
		})
	}
}

func TestGetTransactionsList(t *testing.T) {
	type mockBehavior func(m *mock_service.MockIRepo, inp *domain.TransactionsListInput)

	tt := []struct {
		name   string
		comm   string
		inp    *domain.TransactionsListInput
		mockB  mockBehavior
		expErr error
		expRes []domain.TransactionsListResponse
	}{
		{
			name: "ok",
			comm: "remittance",
			inp: &domain.TransactionsListInput{
				Id:   uuid.MustParse("1993f8f2-d580-4fb1-bd8e-5bdfb7ddd7e4"),
				Sort: "sum",
				Dir:  "asc",
				Page: 1,
			},
			mockB: func(r *mock_service.MockIRepo, inp *domain.TransactionsListInput) {
				r.
					EXPECT().
					GetTransactionsList(gomock.Any(), inp).
					Return([]domain.TransactionsList{}, nil)
			},
			expErr: nil,
			expRes: []domain.TransactionsListResponse{},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)

			defer c.Finish()
			dbMain := mock_service.NewMockIRepo(c)
			tc.mockB(dbMain, tc.inp)
			fs := service.NewFinanceService(dbMain)
			res, err := fs.GetTransactionsList(context.Background(), tc.inp)
			assert.Equal(t, err, tc.expErr)
			assert.Equal(t, res, tc.expRes)
		})
	}
}
