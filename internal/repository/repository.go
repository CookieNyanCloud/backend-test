package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

//todo:uuid
//todo: no init users
//todo:money types

type FinanceRepo struct {
	db *sqlx.DB
}

func NewFinanceRepo(db *sqlx.DB) *FinanceRepo {
	return &FinanceRepo{db: db}
}

type Finance interface {
	Transaction(ctx context.Context, id int, sum float64) error
	Remittance(ctx context.Context, idFrom int, idTo int, sum float64) error
	Balance(ctc context.Context, id int) (float64,error)

	NewTransaction(ctx context.Context, idFrom int, operation string, sum float64, idTo int) error
}

func (r *FinanceRepo) NewFinanceRepo(db *sqlx.DB) *FinanceRepo {
	return &FinanceRepo{db: db}
}

const (
	financeTable = "userbalance"
	transactionTable = "transactions"
)

const (
	Minus   = "недостаточно средств"
)

const (
	transaction = "transaction"
	remittance = "remittance"
	balance = "balance"
)

const (
	rub = "rub"
	usd = "usd"
	eur = "eur"
)

func (r *FinanceRepo) NewTransaction(ctx context.Context, idFrom int, operation string,sum float64, idTo int) error {
	if idTo >0 {
		query := fmt.Sprintf("INSERT INTO %s (user_id, operation,sum, user_to) values ($1, $2, $3, $4) RETURNING id",
			transactionTable)
		_, err := r.db.Exec(query,idFrom,operation,sum, idTo)
		if err != nil {
			return err
		}
	}else {
		query := fmt.Sprintf("INSERT INTO %s (user_id, operation, sum, user_to) values ($1, $2, $3, NULL)",
			transactionTable)
		_, err := r.db.Exec(query,idFrom,operation,sum)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *FinanceRepo) Balance (ctx context.Context, id int) (float64, error)  {
	var currentBalance float64
	query := fmt.Sprintf(`SELECT balance FROM %s WHERE id=$1`, financeTable)
	err := r.db.Get(&currentBalance, query, id)
	if err != nil {
		return -1, err
	}
	return currentBalance, nil
}

func (r *FinanceRepo) Transaction(ctx context.Context, id int, sum float64) error {

	currentBalance, err:= r.Balance(ctx,id)
	if err != nil {
		return err
	}

	newBalance := currentBalance + sum
	if newBalance >= 0 {
		query := fmt.Sprintf("UPDATE %s SET balance = $1  WHERE id = $2", financeTable)
		_, err = r.db.Exec(query, newBalance, id)
		if err != nil {
			return err
		}
		err = r.NewTransaction(ctx, id, transaction, sum,-1)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New(Minus)
}

func (r *FinanceRepo) Remittance(ctx context.Context, idFrom int, idTo int, sum float64) error {
	currentBalanceFrom, err := r.Balance(ctx,idFrom)
	if err != nil {
		return err
	}

	currentBalanceTo, err:= r.Balance(ctx, idTo)
	if err != nil {
		return err
	}

	newBalanceFrom:= currentBalanceFrom - sum
	newBalanceTo:= currentBalanceTo + sum
	if newBalanceFrom >= 0 {
		query := fmt.Sprintf("UPDATE %s SET balance = $1  WHERE id = $2",
			financeTable)
		_, err = r.db.Exec(query, newBalanceFrom, idFrom)
		if err != nil {
			return err
		}

		query = fmt.Sprintf("UPDATE %s SET balance = $1  WHERE id = $2",
			financeTable)
		_, err = r.db.Exec(query, newBalanceTo, idTo)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New(Minus)
}

