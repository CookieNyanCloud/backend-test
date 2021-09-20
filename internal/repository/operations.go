package repository

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type FinanceOperations interface {
	MakeTransaction(id uuid.UUID, sum float64, description string) error
	MakeRemittance(idFrom uuid.UUID, idTo uuid.UUID, sum float64, description string) error
	GetBalance(id uuid.UUID) (float64, error)
	GetTransactionsList(id uuid.UUID, sort string, dir string, page int) ([]TransactionsList, error)
}

func (r *FinanceRepo) MakeTransaction(id uuid.UUID, sum float64, description string) error {
	var noUserError = errors.New(noUser)
	currentBalance, err := r.GetBalance(id)
	if err != nil {
		//check err type
		if err== noUserError && sum > 0 {
			//no user
			err = r.CreateNewUser(id, sum)
			if err != nil {
				return err
			}
			//make task
			err = r.CreateNewTransaction(id, transaction, sum, uuid.Nil, description)
			if err != nil {
				return err
			}
		}
		//other err
		return err
	}
	newBalance := currentBalance + sum
	if newBalance >= 0 {
		query := fmt.Sprintf("UPDATE %s SET balance = $1  WHERE id = $2", financeTable)
		_, err = r.db.Exec(query, newBalance, id)
		if err != nil {
			return err
		}
		err = r.CreateNewTransaction(id, transaction, sum, uuid.Nil, description)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New(Minus)
}

func (r *FinanceRepo) MakeRemittance(idFrom uuid.UUID, idTo uuid.UUID, sum float64, description string) error {
	//check
	var noUserError = errors.New(noUser)
	currentBalanceFrom, err := r.GetBalance(idFrom)
	if err != nil {
		if err == noUserError {
			return errors.New(noSense)
		}
		//other
		return err
	}
	//check
	currentBalanceTo, err := r.GetBalance(idTo)
	if err != nil && err != noUserError {
		//other
		return err
	}
	if err != nil && err != noUserError {
		err = r.CreateNewUser(idTo, 0)
		currentBalanceTo = 0
		if err != nil {
			return err
		}
	}
	newBalanceFrom := currentBalanceFrom - sum
	newBalanceTo := currentBalanceTo + sum
	if newBalanceFrom >= 0 {
		query := fmt.Sprintf("UPDATE %s SET balance = $1  WHERE id = $2",
			financeTable)
		_, err = r.db.Exec(query, newBalanceFrom, idFrom)
		if err != nil {
			return err
		}

		query = fmt.Sprintf("UPDATE %s SET balance = $1  WHERE id = $2",
			financeTable)
		_, err = r.db.Exec(query, newBalanceTo, idTo)
		if err != nil {
			return err
		}

		err = r.CreateNewTransaction(idFrom, remittance, sum, idTo, description)
		if err != nil {
			return err
		}

		return nil
	}
	return errors.New(Minus)
}

func (r *FinanceRepo) GetBalance(id uuid.UUID) (float64, error) {
	var currentBalance float64
	query := fmt.Sprintf(`SELECT balance FROM %s WHERE id=$1`, financeTable)
	err := r.db.Get(&currentBalance, query, id)
	if err != nil {
		return -1, err
	}
	return currentBalance, nil
}

func (r *FinanceRepo) GetTransactionsList(id uuid.UUID, sort string, dir string, page int) ([]TransactionsList, error) {
	limit := 5
	var toVal []listToValidate
	//pagination
	offset := limit * (page - 1)
	if offset < 0 {
		offset = 0
	}
	query := fmt.Sprintf(`SELECT * FROM %s WHERE user_id=$1 OR user_to=$2 ORDER BY "%s"  %s LIMIT %d OFFSET %d`, transactionTable, sort, dir, limit, offset)
	err := r.db.Select(&toVal, query, id, id)
	if err != nil || len(toVal) == 0 {
		return []TransactionsList{}, err
	}
	list := make([]TransactionsList, len(toVal))
	for i := 0; i < len(toVal); i++ {
		list[i].Id = toVal[i].Id
		list[i].Operation = toVal[i].Operation
		list[i].Sum = toVal[i].Sum
		list[i].Date = toVal[i].Date
		list[i].Description = toVal[i].Description
		if toVal[i].IdTo == uuid.Nil {
			list[i].IdTo = ""
		} else {
			list[i].IdTo = toVal[i].IdTo.String()
		}
	}

	return list, nil
}
