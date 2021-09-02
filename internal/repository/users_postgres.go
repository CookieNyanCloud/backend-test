package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UsersRepo struct {
	db *sqlx.DB
}

func NewUsersRepo(db *sqlx.DB) *UsersRepo {
	return &UsersRepo{db: db}
}

func (r *UsersRepo) NewUsersRepo(db *sqlx.DB) *UsersRepo {
	return &UsersRepo{db: db}
}

func (r *UsersRepo) Transaction(ctx context.Context, id uuid.UUID, sum float64) error {
	return nil
}

func (r *UsersRepo) Remittance(ctx context.Context, idFrom uuid.UUID, idTo uuid.UUID, sum float64) error {
	return nil
}

func (r *UsersRepo) Balance(ctx context.Context, id uuid.UUID, cur string) (string, error) {
	return "", nil
}
