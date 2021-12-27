package repository

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

//struct for communication with database
type FinanceRepo struct {
	db *sqlx.DB
}

type IFinanceRepo interface {
	FinanceOperations
	FinanceSubFunctions
}

//new struct
func NewFinanceRepo(db *sqlx.DB) IFinanceRepo {
	return &FinanceRepo{db: db}
}

//struct for  transactions list request
type TransactionsList struct {
	Id             uuid.UUID `json:"id" db:"user_id"`
	Operation      string    `json:"operation"db:"operation"`
	Sum            float64   `json:"sum" db:"sum"`
	Date           time.Time `json:"date" db:"date"`
	Description    string    `json:"description,omitempty" db:"description"`
	IdTo           uuid.UUID `json:"id_to,omitempty" db:"user_to"`
	IdempotencyKey uuid.UUID `json:"idempotency_key" db:"idempotency_key"`
}

const (
	financeTable     = "userbalance"
	transactionTable = "transactions"
)

const (
	transaction = "transaction"
	remittance  = "remittance"
)

var (
	noBalance        = errors.New("недостаточно средств")
	unknownOperation = errors.New("неизвестная операция")
)
