package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

//go:generate mockgen -source=subFunctions.go -destination=mocks/subFunctions/mock.go

type FinanceSubFunctions interface {
	CreateNewUser(ctx context.Context, id uuid.UUID, sum float64) error
	CreateNewTransaction(ctx context.Context, idFrom uuid.UUID, operation string, sum float64, idTo uuid.UUID, description string, idempotencyKey uuid.UUID) error
}

func (r *FinanceRepo) CreateNewUser(ctx context.Context, id uuid.UUID, sum float64) error {
	println("CreateNewUser")
	query := fmt.Sprintf("INSERT INTO %s (id, balance) values ($1, $2)",
		financeTable)
	_, err := r.db.Exec(query, id, sum)
	if err != nil {
		return err
	}

	return nil
}

func (r *FinanceRepo) CreateNewTransaction(ctx context.Context, idFrom uuid.UUID, operation string, sum float64, idTo uuid.UUID, description string, idempotencyKey uuid.UUID) error {
	println("CreateNewTransaction")
	switch operation {
	case remittance:
		query := fmt.Sprintf("INSERT INTO %s (user_id, operation, sum, user_to, description, idempotency_key) values ($1, $2, $3, $4, $5, $6)", transactionTable)
		_, err := r.db.Exec(query, idFrom, operation, sum, idTo, description, idempotencyKey)
		if err != nil {
			return err
		}

	case transaction:
		query := fmt.Sprintf("INSERT INTO %s (user_id, operation, sum, description, idempotency_key) values ($1, $2, $3, $4, $5)", transactionTable)
		_, err := r.db.Exec(query, idFrom, operation, sum, description, idempotencyKey)
		if err != nil {
			return err
		}
	default:
		return UnknownOperation
	}
	return nil
}
