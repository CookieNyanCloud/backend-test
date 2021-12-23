package domain

import "github.com/google/uuid"

type TransactionInput struct {
	Id             uuid.UUID `json:"id" binding:"required"`
	Sum            float64   `json:"sum" binding:"required"`
	Description    string    `json:"description" binding:"max=20"`
	IdempotencyKey uuid.UUID `json:"idempotency_key" binding:"required"`
}

type RemittanceInput struct {
	IdFrom         uuid.UUID `json:"id_from" binding:"required"`
	IdTo           uuid.UUID `json:"id_to" binding:"required"`
	Sum            float64   `json:"sum" binding:"required,gt=0"`
	Description    string    `json:"description" binding:"max=20"`
	IdempotencyKey uuid.UUID `json:"idempotency_key" binding:"required"`
}

type BalanceInput struct {
	Id uuid.UUID `json:"id" binding:"required"`
}

type TransactionsListInput struct {
	Id   uuid.UUID `json:"id" binding:"required"`
	Sort string
	Dir  string
	Page int
}
