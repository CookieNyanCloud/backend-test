package service

import (
	"github.com/cookienyancloud/avito-backend-test/internal/repository"
	"github.com/google/uuid"
)

//прослойка для связи с базой данных


type FinanceService struct {
	repo repository.Finance
}

func NewFinanceService(repo repository.Finance) *FinanceService {
	return &FinanceService{repo}
}

type Finance interface {
	Transaction(id uuid.UUID, sum float64, description string) error
	Remittance(idFrom uuid.UUID, idTo uuid.UUID, sum string, description string) error
	Balance(id int) (float64, error)
	GetTransactionsList(id uuid.UUID, sort string, dir string, page int) ([]repository.TransactionsList, error)
}

func (s *FinanceService) Transaction(id uuid.UUID, sum float64, description string) error {
	return s.repo.Transaction(id, sum, description)
}

func (s *FinanceService) Remittance(idFrom uuid.UUID, idTo uuid.UUID, sum float64, description string) error {
	return s.repo.Remittance(idFrom, idTo, sum, description)
}

func (s *FinanceService) Balance(id uuid.UUID) (float64, error) {
	return s.repo.Balance(id)
}

func (s *FinanceService) GetTransactionsList(id uuid.UUID, sort string, dir string, page int) ([]repository.TransactionsList, error) {
	return s.repo.GetTransactionsList(id, sort, dir, page)
}
