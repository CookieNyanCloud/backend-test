package repository

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

//struct for communication with database
type FinanceRepo struct {
	db *sqlx.DB
}

//new struct
func NewFinanceRepo(db *sqlx.DB) *FinanceRepo {
	return &FinanceRepo{db: db}
}

type IRepo interface {
	IRepoMain
	IRepoSub
}

const (
	financeTable     = "userbalance"
	transactionTable = "transactions"
)

var (
	noBalance        = errors.New("недостаточно средств")
	unknownOperation = errors.New("неизвестная операция")
)
