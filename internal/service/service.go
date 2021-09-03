package service

import (
	"context"
	"github.com/cookienyancloud/avito-backend-test/internal/repository"
)
//todo:uuid

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

type Finance interface {
	Transaction(ctx context.Context, id int, sum float64 ) (error, string)
	Remittance(ctx context.Context, idFrom int, idTo int, sum string ) (error, string)
	BalanceBalance(ctx context.Context, id int, cur string) (string , error)
}



func (s *FinanceService) Transaction(ctx context.Context, id int, sum float64 ) (error, string) {
	return s.repo.Transaction(ctx, id, sum)
}

func (s *FinanceService) Remittance(ctx context.Context, idFrom int, idTo int, sum int64 ) (error, string) {
	return nil, ""
}

func (s *FinanceService) Balance(ctx context.Context, id int, cur string ) (int64 , error) {
	return 0, nil
}
