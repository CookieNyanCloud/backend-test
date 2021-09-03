package repository

import (
	"context"
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
	Transaction(ctx context.Context, id int, sum float64) (error, string)
	Remittance(ctx context.Context, idFrom int, idTo int, sum string) (error, string)
	Balance(ctx context.Context, id int, cur string) (int64, error)
}

const (
	financeTable = "userbalance"
	transactionTable = "transactions"
)

const (
	Unknown = "ошибка на стороне сервера"
	Minus   = "недостаточно средств"
	Success = "удачная транзакция"
)

const (
	rub = "rub"
	usd = "usd"
	eur = "eur"
)

func (r *FinanceRepo) NewFinanceRepo(db *sqlx.DB) *FinanceRepo {
	return &FinanceRepo{db: db}
}

func (r *FinanceRepo) Transaction(ctx context.Context, id int, sum float64) (error, string) {

	var currentBalance float64
	query := fmt.Sprintf(`SELECT balance FROM %s WHERE id=$1`, financeTable)
	err := r.db.Get(&currentBalance, query, id)
	println(currentBalance)
	newBalance:= currentBalance + sum
	if newBalance >=0 {
		query = fmt.Sprintf("UPDATE %s SET balance = $1  WHERE id = $2",
			financeTable)
		_, err = r.db.Exec(query, newBalance)
		if err != nil {
			return err, Unknown
		}
	}else{
		return nil, Minus
	}
	return err, Success
}

func (r *FinanceRepo) Remittance(ctx context.Context, idFrom int, idTo int, sum string) (error, string) {
	return nil, Success
}

func (r *FinanceRepo) Balance(ctx context.Context, id int, cur string) (int64, error) {
	return 0, nil
}
