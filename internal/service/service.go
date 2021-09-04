package service

import (
	"context"
	"github.com/cookienyancloud/avito-backend-test/internal/repository"
)
//todo:uuid

type FinanceService struct {
	repo repository.Finance
}

func NewFinanceService(repo repository.Finance) *FinanceService {
	return &FinanceService{repo}
}

type Finance interface {
	Transaction(ctx context.Context, id int, sum float64 ) error
	Remittance(ctx context.Context, idFrom int, idTo int, sum string ) error
	Balance(ctx context.Context, id int ) (float64 , error)
}

func (s *FinanceService) Transaction(ctx context.Context, id int, sum float64 ) error {
	return s.repo.Transaction(ctx, id, sum)
}

func (s *FinanceService) Remittance(ctx context.Context, idFrom int, idTo int, sum float64 ) error {
	return s.repo.Remittance(ctx,idFrom,idTo,sum)
}

func (s *FinanceService) Balance(ctx context.Context, id int) (float64 , error) {
	return s.repo.Balance(ctx,id)
}
