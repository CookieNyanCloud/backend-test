package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	financeTable = "finance"
)

type FinanceRepo struct {
	db *sqlx.DB
}

func NewFinanceRepo(db *sqlx.DB) *FinanceRepo {
	return &FinanceRepo{db: db}
}

type Finance interface {
	Transaction(ctx context.Context, id uuid.UUID, sum float64) error
	Remittance(ctx context.Context, idFrom uuid.UUID, idTo uuid.UUID, sum float64) error
	Balance(ctx context.Context, id uuid.UUID, cur string) (string, error)
}

func (r *FinanceRepo) NewFinanceRepo(db *sqlx.DB) *FinanceRepo {
	return &FinanceRepo{db: db}
}

func (r *FinanceRepo) Transaction(ctx context.Context, id uuid.UUID, sum float64) error {
	return nil
}

func (r *FinanceRepo) Remittance(ctx context.Context, idFrom uuid.UUID, idTo uuid.UUID, sum float64) error {
	return nil
}

func (r *FinanceRepo) Balance(ctx context.Context, id uuid.UUID, cur string) (string, error) {
	return "", nil
}
