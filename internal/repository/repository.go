package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
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
	Description string    `json:"description,omitempty" db:"description"`
	IdTo        string    `json:"id_to,omitempty" db:"user_to"`
}

type listToValidate struct {
	Id          int           `json:"id" db:"user_id"`
	Operation   string        `json:"operation"db:"operation"`
	Sum         float64       `json:"sum" db:"sum"`
	Date        time.Time     `json:"date" db:"date"`
	Description string        `json:"description,omitempty" db:"description"`
	IdTo        sql.NullInt64 `json:"id_to,omitempty" db:"user_to"`
}

type Finance interface {
	//основные методы
	Transaction(id int, sum float64, description string) error
	Remittance(idFrom int, idTo int, sum float64, description string) error
	Balance(id int) (float64, error)
	GetTransactionsList(id int, sort string, dir string, page int) ([]TransactionsList, error)
	//вспомогательные
	NewUser(id int, sum float64) error
	NewTransaction(idFrom int, operation string, sum float64, idTo int, description string) error
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

func (r *FinanceRepo) Transaction(id int, sum float64, description string) error {
	//получение баланса
	currentBalance, err := r.Balance(id)
	if err != nil {
		//проверка на причину ошибка
		if err.Error() == noUserTxt && sum > 0 {
			//нет пользователя, создаем и начисляем
			err = r.NewUser(id, sum)
			if err != nil {
				return err
			}
			//начисление записываем
			err = r.NewTransaction(id, transaction, sum, -1, description)
			if err != nil {
				return err
			}
		}
		//иная ошибка
		return err
	}
	//новый баланс и его проверка
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

func (r *FinanceRepo) Remittance(idFrom int, idTo int, sum float64, description string) error {
	//проверка баланса отправителя
	currentBalanceFrom, err := r.Balance(idFrom)
	if err != nil {
		if err.Error() == noUser {
			//пользователя нет
			return errors.New(noSense)
		}
		//иная ошибка
		return err
	}
	//проверка баланса получателя
	currentBalanceTo, err := r.Balance(idTo)
	if err != nil && err.Error() != noUserTxt {
		//иная ошибка
		return err
	}
	if err != nil && err.Error() == noUserTxt {
		//создание получателя
		err = r.NewUser(idTo, 0)
		currentBalanceTo = 0
		if err != nil {
			return err
		}
	}
	//новые балансы и проверка
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

		err = r.NewTransaction(idFrom, remittance, sum, idTo, description)
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
			//сообщение об отсутствии пользователя
			return -1, errors.New(noUserTxt)
		}
		return -1, err
	}
	return currentBalance, nil
}

func (r *FinanceRepo) NewUser(id int, sum float64) error {
	//создание пользователя
	query := fmt.Sprintf("INSERT INTO %s (id, balance) values ($1, $2)",
		financeTable)
	_, err := r.db.Exec(query, id, sum)
	if err != nil {
		return err
	}

	return nil
}

func (r *FinanceRepo) NewTransaction(idFrom int, operation string, sum float64, idTo int, description string) error {
	//проверка на наличие получателя, создание соответствующей записи
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

func (r *FinanceRepo) GetTransactionsList(id int, sort string, dir string, page int) ([]TransactionsList, error) {
	limit := 5
	var toVal []listToValidate
	//пагинация
	offset := limit * (page - 1)
	if offset < 0 {
		offset = 0
	}
	//создание списка, ограничение непустыми записями
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
		if toVal[i].IdTo.Valid {
			list[i].IdTo = strconv.FormatInt(toVal[i].IdTo.Int64, 10)
		} else {
			list[i].IdTo = ""
		}
	}

	return list, nil
}
