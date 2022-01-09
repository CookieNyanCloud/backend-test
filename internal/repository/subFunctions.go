package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

//go:generate mockgen -source=subFunctions.go -destination=mocks/subFunctionsMock.go

type IRepoSub interface {
	CreateNewUser(ctx context.Context, id uuid.UUID) error
	CreateNewTransaction(ctx context.Context, idFrom uuid.UUID, operation string, sum float64, idTo uuid.UUID, description string) error
}

const (
	transaction = "transaction"
	remittance  = "remittance"
)

func (r *FinanceRepo) CreateNewUser(ctx context.Context, id uuid.UUID) error {
	query := fmt.Sprintf("INSERT INTO %s (id) values ($1)",
		financeTable)
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *FinanceRepo) CreateNewTransaction(ctx context.Context, idFrom uuid.UUID, operation string, sum float64, idTo uuid.UUID, description string) error {
	switch operation {
	case remittance:
		query := fmt.Sprintf("INSERT INTO %s (user_id, operation, sum, user_to, description) values ($1, $2, $3, $4, $5)", transactionTable)
		_, err := r.db.Exec(ctx, query, idFrom, operation, sum, idTo, description)
		if err != nil {
			return err
		}

	case transaction:
		query := fmt.Sprintf("INSERT INTO %s (user_id, operation, sum, description) values ($1, $2, $3, $4)", transactionTable)
		_, err := r.db.Exec(ctx, query, idFrom, operation, sum, description)
		if err != nil {
			return err
		}
	default:
		return unknownOperation
	}
	return nil
}
