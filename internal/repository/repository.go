package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

//struct for communication with database
type FinanceRepo struct {
	db *pgxpool.Pool
}

//new struct
func NewFinanceRepo(db *pgxpool.Pool) *FinanceRepo {
	return &FinanceRepo{db: db}
}

const (
	financeTable     = "userbalance"
	transactionTable = "transactions"
	transaction      = "transaction"
	remittance       = "remittance"
)
