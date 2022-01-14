package domain

import (
	"time"

	"github.com/google/uuid"
)

//user input for transaction
type TransactionInput struct {
	Id          uuid.UUID `json:"id" binding:"required"`
	Sum         float64   `json:"sum" binding:"required"`
	Description string    `json:"description" binding:"max=20"`
}

//user input for remittance
type RemittanceInput struct {
	IdFrom      uuid.UUID `json:"id_from" binding:"required"`
	IdTo        uuid.UUID `json:"id_to" binding:"required"`
	Sum         float64   `json:"sum" binding:"required,gt=0"`
	Description string    `json:"description" binding:"max=20"`
	//IdempotencyKey uuid.UUID `json:"idempotency_key" binding:"required"`
}

//user input for balance
type BalanceInput struct {
	Id uuid.UUID `json:"id" binding:"required"`
	//Cur string    `json:"-"`
}

//user input for list of transactions
type TransactionsListInput struct {
	Id   uuid.UUID `json:"id" binding:"required"`
	Sort string    `json:"-"`
	Dir  string    `json:"-"`
	Page int       `json:"-"`
}

//struct for  transactions list request
type TransactionsList struct {
	Id          uuid.UUID `json:"id" db:"user_id"`
	Operation   string    `json:"operation"db:"operation"`
	Sum         float64   `json:"sum" db:"sum"`
	Date        time.Time `json:"date" db:"date"`
	Description string    `json:"description,omitempty" db:"description"`
	IdTo        uuid.UUID `json:"id_to,omitempty" db:"user_to"`
}
