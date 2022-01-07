package repository

import (
	"context"
	"fmt"

	"github.com/cookienyancloud/avito-backend-test/internal/domain"
)

type IRepoMain interface {
	MakeTransaction(ctx context.Context, inp *domain.TransactionInput) error
	MakeRemittance(ctx context.Context, inp *domain.RemittanceInput) error
	GetBalance(ctx context.Context, inp *domain.BalanceInput) (float64, error)
	GetTransactionsList(ctx context.Context, inp *domain.TransactionsListInput) ([]*domain.TransactionsList, error)
}

//transaction from user
func (r *FinanceRepo) MakeTransaction(ctx context.Context, inp *domain.TransactionInput) error {
	query := fmt.Sprintf("UPDATE %s SET balance = balance + $1  WHERE id = $2", financeTable)
	_, err := r.db.Exec(query, inp.Sum, inp.Id)
	if err != nil {
		//return err
		return noBalance
	}
	return nil
}

//transaction from user to user
func (r *FinanceRepo) MakeRemittance(ctx context.Context, inp *domain.RemittanceInput) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := fmt.Sprintf("UPDATE %s SET balance = balance - $1  WHERE id = $2", financeTable)
	_, err = r.db.Exec(query, inp.Sum, inp.IdFrom)
	if err != nil {
		return noBalance
	}

	query = fmt.Sprintf("UPDATE %s SET balance = balance + $1  WHERE id = $2", financeTable)
	_, err = r.db.Exec(query, inp.Sum, inp.IdTo)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

//user balance
func (r *FinanceRepo) GetBalance(ctx context.Context, inp *domain.BalanceInput) (float64, error) {
	var currentBalance float64
	query := fmt.Sprintf(`SELECT balance FROM %s WHERE id=$1`, financeTable)
	err := r.db.Get(&currentBalance, query, inp.Id)
	if err != nil {
		return 0, err
	}
	return currentBalance, nil
}

//list of all transactions  by query
func (r *FinanceRepo) GetTransactionsList(ctx context.Context, inp *domain.TransactionsListInput) ([]*domain.TransactionsList, error) {
	limit := 5
	//todo:check
	list := make([]*domain.TransactionsList, limit)
	//pagination
	offset := limit * (inp.Page - 1)
	if offset < 0 {
		offset = 0
	}
	fmt.Println(inp)
	//query := fmt.Sprintf(`SELECT * FROM %s WHERE user_id=$1 OR user_to=$2 ORDER BY $3 $4 LIMIT %d OFFSET %d`, transactionTable, limit, offset)
	//err := r.db.Select(&list, query, inp.Id, inp.Id, inp.Sort, inp.Dir)
	query := fmt.Sprintf(`SELECT * FROM %s WHERE user_id=$1 OR user_to=$2 ORDER BY %s %s LIMIT %d OFFSET %d`, transactionTable, inp.Sort, inp.Dir, limit, offset)
	err := r.db.Select(&list, query, inp.Id, inp.Id)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list, nil
}
