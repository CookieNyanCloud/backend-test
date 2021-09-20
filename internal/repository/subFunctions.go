package repository

import (
	"fmt"
	"github.com/google/uuid"
)

type FinanceSubFunctions interface {
	CreateNewUser(id uuid.UUID, sum float64) error
	CreateNewTransaction(idFrom uuid.UUID, operation string, sum float64, idTo uuid.UUID, description string) error
}

func (r *FinanceRepo) CreateNewUser(id uuid.UUID, sum float64) error {
	query := fmt.Sprintf("INSERT INTO %s (id, balance) values ($1, $2)",
		financeTable)
	_, err := r.db.Exec(query, id, sum)
	if err != nil {
		return err
	}

	return nil
}

func (r *FinanceRepo) CreateNewTransaction(idFrom uuid.UUID, operation string, sum float64, idTo uuid.UUID, description string) error {
	if operation == remittance {
		query := fmt.Sprintf("INSERT INTO %s (user_id, operation, sum, user_to, description) values ($1, $2, $3, $4, $5)",
			transactionTable)
		_, err := r.db.Exec(query, idFrom, operation, sum, idTo, description)
		if err != nil {
			return err
		}
	} else {
		query := fmt.Sprintf("INSERT INTO %s (user_id, operation, sum, description) values ($1, $2, $3, $4)",
			transactionTable)
		_, err := r.db.Exec(query, idFrom, operation, sum, description)
		if err != nil {
			return err
		}
	}
	return nil
}
