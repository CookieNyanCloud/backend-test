package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cookienyancloud/avito-backend-test/internal/domain"
	"github.com/cookienyancloud/avito-backend-test/internal/repository/postgres"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestMakeTransaction(t *testing.T) {
	type args struct {
		id  uuid.UUID
		sum float64
	}
	type mockBehavior func(comm string, args args)

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println(err)
	}
	defer mockDB.Close()

	tt := []struct {
		name   string
		exec   string
		inp    *domain.TransactionInput
		mockB  mockBehavior
		expErr bool
	}{
		{
			name: "ok",
			exec: "INSERT INTO userbalance",
			inp: &domain.TransactionInput{
				Id:          uuid.MustParse("11c52c81-1b31-4c19-b911-791dc6a94f12"),
				Sum:         10,
				Description: "test",
			},
			mockB: func(comm string, args args) {
				mock.
					ExpectExec(comm).
					WithArgs(args.id, args.sum, args.sum).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expErr: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
			tc.mockB(tc.exec, args{
				id:  tc.inp.Id,
				sum: tc.inp.Sum,
			})
			repo := postgres.NewFinanceRepo(sqlxDB)
			err := repo.MakeTransaction(context.Background(), tc.inp)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
			if tc.expErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

		})
	}

}

func TestMakeRemittance(t *testing.T) {

	type args struct {
		idFrom uuid.UUID
		idTo   uuid.UUID
		sum    float64
	}
	type mockBehavior func(comm string, args args)

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println(err)
	}
	defer mockDB.Close()

	tt := []struct {
		name   string
		exec   string
		inp    *domain.RemittanceInput
		mockB  mockBehavior
		expErr bool
	}{
		{
			name: "ok",
			exec: "UPDATE userbalance",
			inp: &domain.RemittanceInput{
				IdFrom:      uuid.MustParse("11c52c81-1b31-4c19-b911-791dc6a94f12"),
				IdTo:        uuid.MustParse("21c52c81-1b31-4c19-b911-791dc6a94f12"),
				Sum:         10,
				Description: "test",
			},
			mockB: func(comm string, args args) {
				mock.ExpectBegin()
				mock.
					ExpectExec(comm).
					WithArgs(args.sum, args.idFrom).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.
					ExpectExec(comm).
					WithArgs(args.sum, args.idTo).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expErr: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
			tc.mockB(tc.exec, args{
				idFrom: tc.inp.IdFrom,
				idTo:   tc.inp.IdTo,
				sum:    tc.inp.Sum,
			})
			repo := postgres.NewFinanceRepo(sqlxDB)
			err := repo.MakeRemittance(context.Background(), tc.inp)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
			if tc.expErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}

func TestGetBalance(t *testing.T) {

	type args struct {
		id uuid.UUID
	}
	type mockBehavior func(comm string, args args, res float64)

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println(err)
	}
	defer mockDB.Close()

	tt := []struct {
		name   string
		exec   string
		inp    *domain.BalanceInput
		mockB  mockBehavior
		expRes float64
		expErr bool
	}{
		{
			name: "ok",
			exec: "SELECT balance FROM userbalance",
			inp: &domain.BalanceInput{
				Id: uuid.MustParse("11c52c81-1b31-4c19-b911-791dc6a94f12"),
			},
			mockB: func(comm string, args args, res float64) {
				mock.
					ExpectQuery(comm).
					WithArgs(args.id).
					WillReturnRows(sqlmock.NewRows([]string{"balance"}).AddRow(res)).
					WillReturnError(nil)
			},
			expRes: 100.1,
			expErr: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
			tc.mockB(tc.exec, args{id: tc.inp.Id}, tc.expRes)
			repo := postgres.NewFinanceRepo(sqlxDB)
			res, err := repo.GetBalance(context.Background(), tc.inp)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
			if tc.expErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expRes, res)
			}
		})
	}

}

func TestGetTransactionsList(t *testing.T) {

	type args struct {
		id uuid.UUID
	}
	type mockBehavior func(comm string, args args, res []domain.TransactionsList)

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println(err)
	}
	defer mockDB.Close()

	tt := []struct {
		name   string
		exec   string
		inp    *domain.TransactionsListInput
		mockB  mockBehavior
		expRes []domain.TransactionsList
		expErr bool
	}{
		{
			name: "ok",
			exec: "SELECT * FROM transactions",
			inp: &domain.TransactionsListInput{
				Id:   uuid.MustParse("11c52c81-1b31-4c19-b911-791dc6a94f12"),
				Sort: "sum",
				Dir:  "asc",
				Page: 1,
			},
			mockB: func(comm string, args args, res []domain.TransactionsList) {
				mock.
					ExpectQuery(comm).
					WithArgs(args.id, args.id).
					WillReturnRows(sqlmock.NewRows(
						[]string{"user_id ", "operation", "sum", "date", "description", "user_to"}).
						AddRow(res[0].Id, res[0].Operation, res[0].Sum, res[0].Date, res[0].Description, res[0].IdTo)).
					WillReturnError(nil)
			},
			expRes: []domain.TransactionsList{
				{
					Id:          uuid.MustParse("11c52c81-1b31-4c19-b911-791dc6a94f12"),
					Operation:   "transaction",
					Sum:         10,
					Date:        time.Now(),
					Description: "test",
					IdTo:        uuid.Nil,
				},
			},
			expErr: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
			tc.mockB(tc.exec, args{id: tc.inp.Id}, tc.expRes)
			repo := postgres.NewFinanceRepo(sqlxDB)
			res, err := repo.GetTransactionsList(context.Background(), tc.inp)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
			if tc.expErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expRes, res)
			}
		})
	}

}

func TestCreateNewTransaction(t *testing.T) {
	type args struct {
		idFrom      uuid.UUID
		operation   string
		sum         float64
		idTo        uuid.UUID
		description string
	}
	type mockBehavior func(comm string, args args)
	execStr := func(in string) string {
		return fmt.Sprintf("INSERT INTO %s", in)
	}

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println(err)
	}
	defer mockDB.Close()

	tt := []struct {
		name   string
		exec   string
		inp    args
		mockB  mockBehavior
		expErr bool
	}{
		{
			name: "ok",
			exec: execStr("transactions"),
			inp: args{
				idFrom:      uuid.MustParse("11c52c81-1b31-4c19-b911-791dc6a94f12"),
				operation:   "transaction",
				sum:         10,
				idTo:        uuid.MustParse("21c52c81-1b31-4c19-b911-791dc6a94f12"),
				description: "test",
			},
			mockB: func(comm string, args args) {
				mock.
					ExpectExec(comm).
					WithArgs(args.idFrom, args.operation, args.sum, args.description).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expErr: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
			tc.mockB(tc.exec, args{
				idFrom:      tc.inp.idFrom,
				operation:   tc.inp.operation,
				sum:         tc.inp.sum,
				idTo:        tc.inp.idTo,
				description: tc.inp.description,
			})
			repo := postgres.NewFinanceRepo(sqlxDB)
			err := repo.CreateNewTransaction(context.Background(), tc.inp.idFrom, tc.inp.operation, tc.inp.sum, tc.inp.idTo, tc.inp.description)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
			if tc.expErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

		})
	}

}
