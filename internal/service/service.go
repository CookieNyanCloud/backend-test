package service

import (
	"context"
	"github.com/cookienyancloud/avito-backend-test/internal/repository"
	"github.com/google/uuid"
)

type UserTransactionInput struct {
	id  uuid.UUID
	sum string
}

type UserRemittanceInput struct {
	idFrom uuid.UUID
	idTo   uuid.UUID
	sum    string
}

type UserBalanceInput struct {
	id  uuid.UUID
	sum string
}

type Users interface {
	Transaction(ctx context.Context, input UserTransactionInput) error
	Remittance(ctx context.Context, input UserRemittanceInput) error
	Balance(ctx context.Context, input UserBalanceInput) (string, error)
}

type Services struct {
	Users
}


func NewServices(repos *repository.Repositories) *Services {
	usersService:=NewUsersService(repos)
	return &Services{
		usersService,
	}
}












