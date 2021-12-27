package repository

import (
	"context"
	"fmt"

	"github.com/cookienyancloud/avito-backend-test/internal/domain"
	"github.com/google/uuid"
)

//go:generate mockgen -source=operations.go -destination=mocks/operations/mock.go

type FinanceOperations interface {
	MakeTransaction(ctx context.Context, inp *domain.TransactionInput) error
	MakeRemittance(ctx context.Context, inp *domain.RemittanceInput) error
	GetBalance(ctx context.Context, inp *domain.BalanceInput) (float64, error)
	GetTransactionsList(ctx context.Context, inp *domain.TransactionsListInput) ([]TransactionsList, error)
}

func (r *FinanceRepo) MakeTransaction(ctx context.Context, inp *domain.TransactionInput) error {
	println("MakeTransaction")
	balance := &domain.BalanceInput{
		Id: inp.Id,
	}
	_, err := r.GetBalance(ctx, balance)
	if err != nil {
		//no user
		err = r.CreateNewUser(ctx, inp.Id, inp.Sum)
		if err != nil {
			return err
		}
		//make task
		err = r.CreateNewTransaction(ctx, inp.Id, transaction, inp.Sum, uuid.Nil, inp.Description, inp.IdempotencyKey)
		if err != nil {
			return err
		}
		return nil
	}

	query := fmt.Sprintf("UPDATE %s SET balance = balance + $1  WHERE id = $2", financeTable)
	_, err = r.db.Exec(query, inp.Sum, inp.Id)
	if err != nil {
		//return err
		return NoBalance
	}
	err = r.CreateNewTransaction(ctx, inp.Id, transaction, inp.Sum, uuid.Nil, inp.Description, inp.IdempotencyKey)
	if err != nil {
		return err
	}
	return nil
}

func (r *FinanceRepo) MakeRemittance(ctx context.Context, inp *domain.RemittanceInput) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = r.GetBalance(ctx, &domain.BalanceInput{Id: inp.IdFrom})
	if err != nil {
		return NoBalance
	}

	_, err = r.GetBalance(ctx, &domain.BalanceInput{Id: inp.IdTo})
	if err != nil {
		err = r.CreateNewUser(ctx, inp.IdTo, 0)
		if err != nil {
			return err
		}
	}

	query := fmt.Sprintf("UPDATE %s SET balance = balance - $1  WHERE id = $2", financeTable)
	_, err = r.db.Exec(query, inp.Sum, inp.IdFrom)
	if err != nil {
		return NoBalance
	}

	query = fmt.Sprintf("UPDATE %s SET balance = balance + $1  WHERE id = $2", financeTable)
	_, err = r.db.Exec(query, inp.Sum, inp.IdTo)
	if err != nil {
		return err
	}

	err = r.CreateNewTransaction(ctx, inp.IdFrom, remittance, inp.Sum, inp.IdTo, inp.Description, inp.IdempotencyKey)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *FinanceRepo) GetBalance(ctx context.Context, inp *domain.BalanceInput) (float64, error) {
	println("GetBalance")

	var currentBalance float64
	query := fmt.Sprintf(`SELECT balance FROM %s WHERE id=$1`, financeTable)
	err := r.db.Get(&currentBalance, query, inp.Id)
	if err != nil {
		return 0, NoBalance
	}
	return currentBalance, nil
}

func (r *FinanceRepo) GetTransactionsList(ctx context.Context, inp *domain.TransactionsListInput) ([]TransactionsList, error) {
	limit := 5
	var list []TransactionsList
	//pagination
	offset := limit * (inp.Page - 1)
	if offset < 0 {
		offset = 0
	}
	query := fmt.Sprintf(`SELECT * FROM %s WHERE user_id=$1 OR user_to=$2 ORDER BY $3  $4 LIMIT %d OFFSET %d`, transactionTable, limit, offset)
	err := r.db.Select(&list, query, inp.Id, inp.Id, inp.Sort, inp.Dir)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list, nil
}
