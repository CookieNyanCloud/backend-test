package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "users"
)

type Users interface {
	Transaction(ctx context.Context, id uuid.UUID, sum float64) error
	Remittance(ctx context.Context, idFrom uuid.UUID, idTo uuid.UUID, sum float64) error
	Balance(ctx context.Context, id uuid.UUID, cur string) (string, error)
}

type Repositories struct {
	Users
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		NewUsersRepo(db),
	}
}
