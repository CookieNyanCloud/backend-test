package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type FinanceRepo struct {
	db *sqlx.DB
}

func NewFinanceRepo(db *sqlx.DB) *FinanceRepo {
	return &FinanceRepo{db: db}
}

type TransactionsList struct {
	Id          int       `json:"id" db:"user_id"`
	Operation   string    `json:"operation"db:"operation"`
	Sum         float64   `json:"sum" db:"sum"`
	Date        time.Time `json:"date" db:"date"`
	Description string    `json:"description" db:"description"`
	IdTo        int       `json:"id_to" db:"user_to"`
}


type Finance interface {
	Transaction(id int, sum float64, description string) error
	Remittance(idFrom int, idTo int, sum float64, description string) error
	Balance(id int) (float64, error)

	NewUser(id int, sum float64) error
	NewTransaction(idFrom int, operation string, sum float64, idTo int, description string) error
	GetTransactionsList(id int, sort string,dir string, page int) ([]TransactionsList, error)
}

func (r *FinanceRepo) NewFinanceRepo(db *sqlx.DB) *FinanceRepo {
	return &FinanceRepo{db: db}
}

const (
	financeTable     = "userbalance"
	transactionTable = "transactions"
)

const (
	Minus     = "недостаточно средств"
	noUser    = "sql: no rows in result set"
	noUserTxt = "пользователя нет"
	noSense   = "создание пользователя с заранее отрицательным балансом"
)

const (
	transaction = "transaction"
	remittance  = "remittance"
)

func (r *FinanceRepo) Transaction(id int, sum float64,description string) error {

	currentBalance, err := r.Balance(id)
	if err != nil {
		if err.Error() == noUserTxt && sum > 0 {
			err = r.NewUser(id, sum)
			if err != nil {
				return err
			}
			err = r.NewTransaction(id, transaction, sum, -1,description)
			if err != nil {
				return err
			}
		}
		return err
	}

	newBalance := currentBalance + sum
	if newBalance >= 0 {
		query := fmt.Sprintf("UPDATE %s SET balance = $1  WHERE id = $2", financeTable)
		_, err = r.db.Exec(query, newBalance, id)
		if err != nil {
			return err
		}
		err = r.NewTransaction(id, transaction, sum, -1, description)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New(Minus)
}

func (r *FinanceRepo) Remittance(idFrom int, idTo int, sum float64,description string) error {
	currentBalanceFrom, err := r.Balance(idFrom)
	if err != nil {
		if err.Error() == noUser {
			return errors.New(noSense)
		}
		return err
	}

	currentBalanceTo, err := r.Balance(idTo)
	if err != nil && err.Error() != noUserTxt {
		println("DS")
		return err
	}
	if err != nil && err.Error() == noUserTxt {
		err = r.NewUser(idTo, 0)
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

		err = r.NewTransaction(idFrom, remittance, sum, idTo,description)
		if err != nil {
			return err
		}

		return nil
	}
	return errors.New(Minus)
}

func (r *FinanceRepo) Balance(id int) (float64, error) {
	var currentBalance float64
	query := fmt.Sprintf(`SELECT balance FROM %s WHERE id=$1`, financeTable)
	err := r.db.Get(&currentBalance, query, id)
	if err != nil {
		if err.Error() == noUser {
			return -1, errors.New(noUserTxt)
		}
		return -1, err
	}
	return currentBalance, nil
}

func (r *FinanceRepo) NewUser(id int, sum float64) error {
	query := fmt.Sprintf("INSERT INTO %s (id, balance) values ($1, $2)",
		financeTable)
	_, err := r.db.Exec(query, id, sum)
	if err != nil {
		return err
	}

	return nil
}

func (r *FinanceRepo) NewTransaction(idFrom int, operation string, sum float64, idTo int,description string) error {
	if idTo > 0 {
		query := fmt.Sprintf("INSERT INTO %s (user_id, operation,sum, user_to, description) values ($1, $2, $3, $4, $5)",
			transactionTable)
		_, err := r.db.Exec(query, idFrom, operation, sum, idTo, description)
		if err != nil {
			return err
		}
	} else {
		query := fmt.Sprintf("INSERT INTO %s (user_id, operation, sum, user_to, description) values ($1, $2, $3, -1,$4)",
			transactionTable)
		_, err := r.db.Exec(query, idFrom, operation, sum, description)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *FinanceRepo) GetTransactionsList(id int, sort string,dir string, page int) ([]TransactionsList, error) {
	var list []TransactionsList
	limit:= 5
	offset:=limit*(page-1)
	query := fmt.Sprintf(`SELECT * FROM %s WHERE user_id=$1 ORDER BY "%s"  %s LIMIT %d OFFSET %d`, transactionTable,sort, dir, limit, offset)
	err := r.db.Select(&list, query, id)
	if err != nil {
		return []TransactionsList{}, err
	}
	return list, nil
}
