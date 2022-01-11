package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cookienyancloud/avito-backend-test/internal/domain"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

//go:generate mockgen -source=operations.go -destination=mocks/operationsMock.go

type IRepo interface {
	//main
	MakeTransaction(ctx context.Context, inp *domain.TransactionInput) error
	MakeRemittance(ctx context.Context, inp *domain.RemittanceInput) error
	GetBalance(ctx context.Context, inp *domain.BalanceInput) (float64, error)
	GetTransactionsList(ctx context.Context, inp *domain.TransactionsListInput) ([]domain.TransactionsList, error)
	//sub
	CreateNewTransaction(ctx context.Context, idFrom uuid.UUID, operation string, sum float64, idTo uuid.UUID, description string) error
}

//transaction from user
func (r *FinanceRepo) MakeTransaction(ctx context.Context, inp *domain.TransactionInput) error {
	query := fmt.Sprintf(`
INSERT INTO %s (id, balance)
values ($1, $2) 
ON CONFLICT (id) DO UPDATE 
SET balance = %s.balance + $3`, financeTable, financeTable)
	_, err := r.db.ExecContext(ctx, query, inp.Id, inp.Sum, inp.Sum)
	if err != nil {
		return errors.Wrap(err, "exec")
	}
	return nil
}

//transaction from user to user
func (r *FinanceRepo) MakeRemittance(ctx context.Context, inp *domain.RemittanceInput) error {
	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: 0,
		ReadOnly:  false,
	})
	if err != nil {
		return errors.Wrap(err, "begin transaction")
	}
	defer tx.Rollback()

	query := fmt.Sprintf("UPDATE %s SET balance = balance - $1  WHERE id = $2", financeTable)
	_, err = r.db.ExecContext(ctx, query, inp.Sum, inp.IdFrom)
	if err != nil {
		return errors.Wrap(err, "update first")
	}

	query = fmt.Sprintf("UPDATE %s SET balance = balance + $1  WHERE id = $2", financeTable)
	_, err = r.db.ExecContext(ctx, query, inp.Sum, inp.IdTo)
	if err != nil {
		return errors.Wrap(err, "update second")
	}
	if err = tx.Commit(); err != nil {
		return errors.Wrap(err, "commit")
	}
	return nil
}

//user balance
func (r *FinanceRepo) GetBalance(ctx context.Context, inp *domain.BalanceInput) (float64, error) {
	var currentBalance float64
	query := fmt.Sprintf(`SELECT balance FROM %s WHERE id=$1`, financeTable)
	if err := r.db.QueryRowContext(ctx, query, inp.Id).Scan(&currentBalance); err != nil {
		return 0, errors.Wrap(err, "scanning")
	}
	return currentBalance, nil
}

//list of all transactions  by query
func (r *FinanceRepo) GetTransactionsList(ctx context.Context, inp *domain.TransactionsListInput) ([]domain.TransactionsList, error) {
	limit := 5
	var list []domain.TransactionsList
	//pagination
	offset := limit * (inp.Page - 1)
	if offset < 0 {
		offset = 0
	}
	query := fmt.Sprintf(`SELECT * FROM %s WHERE user_id= $1 OR user_to= $2 ORDER BY %s %s LIMIT %d OFFSET %d`, transactionTable, inp.Sort, inp.Dir, limit, offset)
	err := r.db.SelectContext(ctx, &list, query, inp.Id, inp.Id)
	if err != nil {
		return nil, errors.Wrap(err, "select")
	}
	return list, nil
}

//list set of transactions
func (r *FinanceRepo) CreateNewTransaction(ctx context.Context, idFrom uuid.UUID, operation string, sum float64, idTo uuid.UUID, description string) error {
	switch operation {
	case remittance:
		query := fmt.Sprintf("INSERT INTO %s (user_id, operation, sum, user_to, description) values ($1, $2, $3, $4, $5)", transactionTable)
		_, err := r.db.ExecContext(ctx, query, idFrom, operation, sum, idTo, description)
		if err != nil {
			return errors.Wrap(err, "exec remittance")
		}

	case transaction:
		query := fmt.Sprintf("INSERT INTO %s (user_id, operation, sum, description) values ($1, $2, $3, $4)", transactionTable)
		_, err := r.db.ExecContext(ctx, query, idFrom, operation, sum, description)
		if err != nil {
			return errors.Wrap(err, "exec transaction")
		}
	default:
		return errors.New("неизвестная операция")
	}
	return nil
}
