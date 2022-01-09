package repository

import (
	"errors"

	"github.com/jackc/pgx/v4"
)

//struct for communication with database
type FinanceRepo struct {
	db *pgx.Conn
}

//new struct
func NewFinanceRepo(db *pgx.Conn) *FinanceRepo {
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
