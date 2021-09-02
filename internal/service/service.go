package service

import (
	"context"
	"github.com/cookienyancloud/avito-backend-test/internal/repository"
	"github.com/google/uuid"
)

type FinanceService struct {
	repo repository.Finance
}

func NewFinanceService(
	repo repository.Finance,
) *FinanceService {
	return &FinanceService{
		repo,
	}
}

type TransactionInput struct {
	id  uuid.UUID
	sum string
}

type RemittanceInput struct {
	idFrom uuid.UUID
	idTo   uuid.UUID
	sum    string
}

type BalanceInput struct {
	id  uuid.UUID
	sum string
}

type Finance interface {
	Transaction(ctx context.Context, input TransactionInput) error
	Remittance(ctx context.Context, input RemittanceInput) error
	Balance(ctx context.Context, input BalanceInput) (string, error)
}

func (s *FinanceService) Transaction(ctx context.Context, input TransactionInput) error {
	return  nil
}

func (s *FinanceService) Remittance(ctx context.Context, input RemittanceInput) error {
	return  nil
}

func (s *FinanceService) Balance(ctx context.Context, input BalanceInput) (string, error) {
	return  "",nil
}
















