package service

import (
	"context"

	"github.com/cookienyancloud/avito-backend-test/internal/domain"
	"github.com/cookienyancloud/avito-backend-test/internal/repository"
)

type FinanceService struct {
	repo repository.FinanceOperations
}

func NewFinanceService(repo repository.FinanceOperations) *FinanceService {
	return &FinanceService{repo}
}

type IFinance interface {
	MakeTransaction(ctx context.Context, inp *domain.TransactionInput) error
	MakeRemittance(ctx context.Context, inp *domain.RemittanceInput) error
	GetBalance(ctx context.Context, inp *domain.BalanceInput) (float64, error)
	GetTransactionsList(ctx context.Context, inp *domain.TransactionsListInput) ([]repository.TransactionsList, error)
}

func (s *FinanceService) MakeTransaction(ctx context.Context, inp *domain.TransactionInput) error {
	return s.repo.MakeTransaction(ctx, inp)
}

func (s *FinanceService) MakeRemittance(ctx context.Context, inp *domain.RemittanceInput) error {
	return s.repo.MakeRemittance(ctx, inp)
}

func (s *FinanceService) GetBalance(ctx context.Context, inp *domain.BalanceInput) (float64, error) {
	return s.repo.GetBalance(ctx, inp)
}

func (s *FinanceService) GetTransactionsList(ctx context.Context, inp *domain.TransactionsListInput) ([]repository.TransactionsList, error) {
	return s.repo.GetTransactionsList(ctx, inp)
}
