package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/cookienyancloud/avito-backend-test/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

//go:generate mockgen -source=operations.go -destination=mocks/operationsMock.go

type IRepoMain interface {
	MakeTransaction(ctx context.Context, inp *domain.TransactionInput) error
	MakeRemittance(ctx context.Context, inp *domain.RemittanceInput) error
	GetBalance(ctx context.Context, inp *domain.BalanceInput) (float64, error)
	GetTransactionsList(ctx context.Context, inp *domain.TransactionsListInput) ([]*domain.TransactionsList, error)
}

//transaction from user
func (r *FinanceRepo) MakeTransaction(ctx context.Context, inp *domain.TransactionInput) error {
	query := fmt.Sprintf("UPDATE %s SET balance = balance + $1  WHERE id = $2", financeTable)
	_, err := r.db.Exec(ctx, query, inp.Sum, inp.Id)
	if err != nil {
		//return err
		return noBalance
	}
	return nil
}

//transaction from user to user
func (r *FinanceRepo) MakeRemittance(ctx context.Context, inp *domain.RemittanceInput) error {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:       pgx.ReadCommitted,
		AccessMode:     pgx.ReadOnly,
		DeferrableMode: pgx.Deferrable,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := fmt.Sprintf("UPDATE %s SET balance = balance - $1  WHERE id = $2", financeTable)
	_, err = r.db.Exec(ctx, query, inp.Sum, inp.IdFrom)
	if err != nil {
		return noBalance
	}

	query = fmt.Sprintf("UPDATE %s SET balance = balance + $1  WHERE id = $2", financeTable)
	_, err = r.db.Exec(ctx, query, inp.Sum, inp.IdTo)
	if err != nil {
		return err
	}
	if err = tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}

//user balance
func (r *FinanceRepo) GetBalance(ctx context.Context, inp *domain.BalanceInput) (float64, error) {
	var currentBalance float64
	query := fmt.Sprintf(`SELECT balance FROM %s WHERE id=$1`, financeTable)
	err := r.db.QueryRow(ctx, query, inp.Id).Scan(&currentBalance)
	if err != nil {
		return 0, err
	}
	return currentBalance, nil
}

//list of all transactions  by query
func (r *FinanceRepo) GetTransactionsList(ctx context.Context, inp *domain.TransactionsListInput) ([]*domain.TransactionsList, error) {
	limit := 5
	//todo:check
	list := make([]*domain.TransactionsList, limit)
	//pagination
	offset := limit * (inp.Page - 1)
	if offset < 0 {
		offset = 0
	}
	query := fmt.Sprintf(`SELECT * FROM %s WHERE user_id=$1 OR user_to=$2 ORDER BY %s %s LIMIT %d OFFSET %d`, transactionTable, inp.Sort, inp.Dir, limit, offset)
	rows, err := r.db.Query(ctx, query, inp.Id, inp.Id)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	var id uuid.UUID
	var op string
	var sum float64
	var date time.Time
	var desc string
	var idTo uuid.UUID
	if rows.Next() {
		err := rows.Scan(&id, &op, &sum, &date, &desc, &idTo)
		if err != nil {
			return nil, err
		}
		list = append(list, &domain.TransactionsList{
			Id:          id,
			Operation:   op,
			Sum:         sum,
			Date:        date,
			Description: desc,
			IdTo:        idTo,
		})

	}
	return list, rows.Err()
}
