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

//type transactionsList struct {
//	Id          int
//	Operation   string
//	Sum         float64
//	Date        time.Time
//	Description string
//	IdTo        int
//}

type Finance interface {
	Transaction(id int, sum float64 ) error
	Remittance(idFrom int, idTo int, sum string ) error
	Balance(id int ) (float64 , error)
	GetTransactionsList(id int, sort string,dir string, page int) ([]repository.TransactionsList, error)
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

func (s *FinanceService) GetTransactionsList(id int, sort string,dir string, page int) ([]repository.TransactionsList, error){
	return s.repo.GetTransactionsList(id,sort,dir,page)
}
