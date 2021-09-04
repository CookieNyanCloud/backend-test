package service

import (
	"github.com/cookienyancloud/avito-backend-test/internal/repository"
)

type FinanceService struct {
	repo repository.Finance
}

func NewFinanceService(repo repository.Finance) *FinanceService {
	return &FinanceService{repo}
}

type Finance interface {
	Transaction(id int, sum float64 ) error
	Remittance(idFrom int, idTo int, sum string ) error
	Balance(id int ) (float64 , error)
}

func (s *FinanceService) Transaction(id int, sum float64 ) error {
	return s.repo.Transaction(id, sum)
}

func (s *FinanceService) Remittance(idFrom int, idTo int, sum float64 ) error {
	return s.repo.Remittance(idFrom,idTo,sum)
}

func (s *FinanceService) Balance(id int) (float64 , error) {
	return s.repo.Balance(id)
}
