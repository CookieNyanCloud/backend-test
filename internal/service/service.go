package service

import (
	"github.com/cookienyancloud/avito-backend-test/internal/repository"
	"github.com/google/uuid"
)

type FinanceService struct {
	repo repository.FinanceOperations
}

func NewFinanceService(repo repository.FinanceOperations) *FinanceService {
	return &FinanceService{repo}
}

type Finance interface {
	MakeTransaction(id uuid.UUID, sum float64, description string) error
	MakeRemittance(idFrom uuid.UUID, idTo uuid.UUID, sum string, description string) error
	GetBalance(id int) (float64, error)
	GetTransactionsList(id uuid.UUID, sort string, dir string, page int) ([]repository.TransactionsList, error)
}

func (s *FinanceService) MakeTransaction(id uuid.UUID, sum float64, description string) error {
	return s.repo.MakeTransaction(id, sum, description)
}

func (s *FinanceService) MakeRemittance(idFrom uuid.UUID, idTo uuid.UUID, sum float64, description string) error {
	return s.repo.MakeRemittance(idFrom, idTo, sum, description)
}

func (s *FinanceService) GetBalance(id uuid.UUID) (float64, error) {
	return s.repo.GetBalance(id)
}

func (s *FinanceService) GetTransactionsList(id uuid.UUID, sort string, dir string, page int) ([]repository.TransactionsList, error) {
	return s.repo.GetTransactionsList(id, sort, dir, page)
}
