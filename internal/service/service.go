package service

import (
	"context"

	"github.com/cookienyancloud/avito-backend-test/internal/domain"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

//go:generate mockgen -source=service.go -destination=mocks/RepoMock.go

type IRepo interface {
	//main
	MakeTransaction(ctx context.Context, inp *domain.TransactionInput) error
	MakeRemittance(ctx context.Context, inp *domain.RemittanceInput) error
	GetBalance(ctx context.Context, inp *domain.BalanceInput) (float64, error)
	GetTransactionsList(ctx context.Context, inp *domain.TransactionsListInput) ([]domain.TransactionsList, error)
	//sub
	CreateNewTransaction(ctx context.Context, idFrom uuid.UUID, operation string, sum float64, idTo uuid.UUID, description string) error
	StartMigration(ctx context.Context, dir, dest string) error
	Close(ctx context.Context) error
}

type FinanceService struct {
	repo IRepo
}

func NewFinanceService(repo IRepo) *FinanceService {
	return &FinanceService{repo}
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
	if err != nil {
		return nil, errors.Wrap(err, "get list")
	}
	responses := make([]domain.TransactionsListResponse, len(list))
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

func (f *FinanceService) StartMigration(ctx context.Context, dir, dest string) error {
	return f.repo.StartMigration(ctx, dir, dest)
}

func (f *FinanceService) Close(ctx context.Context) error {
	return f.repo.Close(ctx)
}
