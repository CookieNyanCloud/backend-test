package domain

import "github.com/google/uuid"

//user input for transaction
type TransactionInput struct {
	Id             uuid.UUID `json:"id" binding:"required"`
	Sum            float64   `json:"sum" binding:"required"`
	Description    string    `json:"description" binding:"max=20"`
	IdempotencyKey uuid.UUID `json:"idempotency_key" binding:"required"`
}

//user input for remittance
type RemittanceInput struct {
	IdFrom         uuid.UUID `json:"id_from" binding:"required"`
	IdTo           uuid.UUID `json:"id_to" binding:"required"`
	Sum            float64   `json:"sum" binding:"required,gt=0"`
	Description    string    `json:"description" binding:"max=20"`
	IdempotencyKey uuid.UUID `json:"idempotency_key" binding:"required"`
}

//user input for balance
type BalanceInput struct {
	Id uuid.UUID `json:"id" binding:"required"`
}

//user input for list of transactions
type TransactionsListInput struct {
	Id   uuid.UUID `json:"id" binding:"required"`
	Sort string    `json:"-"`
	Dir  string    `json:"-"`
	Page int       `json:"-"`
}
