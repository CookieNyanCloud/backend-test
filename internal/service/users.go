package service

import (
	"context"
	"github.com/cookienyancloud/avito-backend-test/internal/repository"
)

type UsersService struct {
	repo repository.Users
}

func NewUsersService(
	repo repository.Users,
) *UsersService {
	return &UsersService{
		repo,
	}
}

func (s *UsersService) Transaction(ctx context.Context, input UserTransactionInput) error {
	return  nil
}

func (s *UsersService) Remittance(ctx context.Context, input UserRemittanceInput) error {
	return  nil
}

func (s *UsersService) Balance(ctx context.Context, input UserBalanceInput) (string, error) {
	return  "",nil
}

