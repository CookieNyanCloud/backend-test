package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
)

type FinanceRepo struct {
	db *sqlx.DB
}

func NewFinanceRepo(db *sqlx.DB) *FinanceRepo {
	return &FinanceRepo{db: db}
}

type TransactionsList struct {
	Id          uuid.UUID `json:"id" db:"user_id"`
	Operation   string    `json:"operation"db:"operation"`
	Sum         float64   `json:"sum" db:"sum"`
	Date        time.Time `json:"date" db:"date"`
	Description string    `json:"description,omitempty" db:"description"`
	IdTo        string    `json:"id_to,omitempty" db:"user_to"`
}

type listToValidate struct {
	Id          uuid.UUID `json:"id" db:"user_id"`
	Operation   string    `json:"operation"db:"operation"`
	Sum         float64   `json:"sum" db:"sum"`
	Date        time.Time `json:"date" db:"date"`
	Description string    `json:"description,omitempty" db:"description"`
	IdTo        uuid.UUID `json:"id_to,omitempty" db:"user_to"`
}

const (
	financeTable     = "userbalance"
	transactionTable = "transactions"
)

const (
	Minus     = "недостаточно средств"
	noUser    = "sql: no rows in result set"
	noUserTxt = "пользователя нет"
	noSense   = "создание пользователя с заранее отрицательным балансом"
)

const (
	transaction = "transaction"
	remittance  = "remittance"
)






