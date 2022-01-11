package service

import (
	"context"

	"github.com/cookienyancloud/avito-backend-test/internal/domain"
	"github.com/cookienyancloud/avito-backend-test/internal/repository"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

//go:generate mockgen -source=service.go -destination=mocks/serviceMock.go

type FinanceService struct {
	repo repository.IRepo
}

func NewFinanceService(repo repository.IRepo) *FinanceService {
	return &FinanceService{repo}
}

type IService interface {
	MakeTransaction(ctx context.Context, inp *domain.TransactionInput) error
	MakeRemittance(ctx context.Context, inp *domain.RemittanceInput) error
	GetBalance(ctx context.Context, inp *domain.BalanceInput) (float64, error)
	GetTransactionsList(ctx context.Context, inp *domain.TransactionsListInput) ([]domain.TransactionsListResponse, error)
}

func (f *FinanceService) MakeTransaction(ctx context.Context, inp *domain.TransactionInput) error {
	if err := f.repo.MakeTransaction(ctx, inp); err != nil {
		return errors.Wrap(err, "transaction")
	}
	if err := f.repo.CreateNewTransaction(ctx, inp.Id, "transaction", inp.Sum, uuid.Nil, inp.Description); err != nil {
		return errors.Wrap(err, "create transaction")
	}
	return nil
}

func (f *FinanceService) MakeRemittance(ctx context.Context, inp *domain.RemittanceInput) error {
	if err := f.repo.MakeRemittance(ctx, inp); err != nil {
		return errors.Wrap(err, "remittance")
	}
	if err := f.repo.CreateNewTransaction(ctx, inp.IdFrom, "remittance", inp.Sum, inp.IdTo, inp.Description); err != nil {
		return errors.Wrap(err, "create transaction")
	}
	return nil
}

func (f *FinanceService) GetBalance(ctx context.Context, inp *domain.BalanceInput) (float64, error) {
	return f.repo.GetBalance(ctx, inp)
}

func (f *FinanceService) GetTransactionsList(ctx context.Context, inp *domain.TransactionsListInput) ([]domain.TransactionsListResponse, error) {
	list, err := f.repo.GetTransactionsList(ctx, inp)
	responses := make([]domain.TransactionsListResponse, len(list))
	if err != nil {
		return nil, err
	}
	for i := range list {
		if list[i].IdTo == uuid.Nil {
			responses[i].IdTo = ""
		}
		responses[i].Id = list[i].Id
		responses[i].Description = list[i].Description
		responses[i].Operation = list[i].Operation
		responses[i].Date = list[i].Date
		responses[i].Sum = list[i].Sum
	}
	return responses, nil
}
