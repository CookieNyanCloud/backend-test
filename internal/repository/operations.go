package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/cookienyancloud/avito-backend-test/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

//go:generate mockgen -source=operations.go -destination=mocks/operationsMock.go

type IRepoMain interface {
	MakeTransaction(ctx context.Context, inp *domain.TransactionInput) error
	MakeRemittance(ctx context.Context, inp *domain.RemittanceInput) error
	GetBalance(ctx context.Context, inp *domain.BalanceInput) (float64, error)
	GetTransactionsList(ctx context.Context, inp *domain.TransactionsListInput) ([]domain.TransactionsList, error)
}

//transaction from user
func (r *FinanceRepo) MakeTransaction(ctx context.Context, inp *domain.TransactionInput) error {
	query := fmt.Sprintf(`
INSERT INTO %s (id, balance)
values ($1, $2) 
ON CONFLICT (id) DO UPDATE 
SET balance = %s.balance + $3`, financeTable, financeTable)
	_, err := r.db.Exec(ctx, query, inp.Id, inp.Sum, inp.Sum)
	if err != nil {
		//return err
		return errors.Wrap(err, "exec")
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
		return errors.Wrap(err, "begin transaction")
	}
	defer tx.Rollback(ctx)

	query := fmt.Sprintf("UPDATE %s SET balance = balance - $1  WHERE id = $2", financeTable)
	_, err = r.db.Exec(ctx, query, inp.Sum, inp.IdFrom)
	if err != nil {
		return errors.Wrap(err, "update first")
	}

	query = fmt.Sprintf("UPDATE %s SET balance = balance + $1  WHERE id = $2", financeTable)
	_, err = r.db.Exec(ctx, query, inp.Sum, inp.IdTo)
	if err != nil {
		return errors.Wrap(err, "update second")
	}
	if err = tx.Commit(ctx); err != nil {
		return errors.Wrap(err, "commit")
	}
	return nil
}

//user balance
func (r *FinanceRepo) GetBalance(ctx context.Context, inp *domain.BalanceInput) (float64, error) {
	var currentBalance float64
	query := fmt.Sprintf(`SELECT balance FROM %s WHERE id=$1`, financeTable)
	if err := r.db.QueryRow(ctx, query, inp.Id).Scan(&currentBalance); err != nil {
		return 0, errors.Wrap(err, "scanning")
	}
	return currentBalance, nil
}

//list of all transactions  by query
func (r *FinanceRepo) GetTransactionsList(ctx context.Context, inp *domain.TransactionsListInput) ([]domain.TransactionsList, error) {
	limit := 5
	list := make([]domain.TransactionsList, limit)
	//pagination
	offset := limit * (inp.Page - 1)
	if offset < 0 {
		offset = 0
	}
	query := fmt.Sprintf(`SELECT * FROM %s WHERE user_id= $1 OR user_to= $2 ORDER BY %s %s LIMIT %d OFFSET %d`, transactionTable, inp.Sort, inp.Dir, limit, offset)
	rows, err := r.db.Query(ctx, query, inp.Id, inp.Id)
	if err != nil {
		return nil, errors.Wrap(err, "query")
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
			return nil, errors.Wrap(err, "scan")
		}
		list = append(list, domain.TransactionsList{
			Id:          id,
			Operation:   op,
			Sum:         sum,
			Date:        date,
			Description: desc,
			IdTo:        idTo,
		})

	}
	if rows.Err() != nil {
		return nil, errors.Wrap(err, "rows")
	}
	return list, nil
}
