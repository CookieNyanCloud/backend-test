package repository

import (
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

const (
	financeTable     = "userbalance"
	transactionTable = "transactions"
	transaction      = "transaction"
	remittance       = "remittance"
)
