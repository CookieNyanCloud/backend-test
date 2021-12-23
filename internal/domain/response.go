package domain

import (
	"database/sql"
	"time"
)

type Response struct {
	Message string `json:"message"`
}

type BalanceResponse struct {
	Balance string `json:"balanceResponse"`
	Cur     string `json:"cur"`
}
type TransactionsListResponse struct {
	Id          int           `json:"id" db:"user_id"`
	Operation   string        `json:"operation"db:"operation"`
	Sum         string        `json:"sum" db:"sum"`
	Date        time.Time     `json:"date" db:"date"`
	Description string        `json:"description,omitempty" db:"description"`
	IdTo        sql.NullInt64 `json:"id_to,omitempty" db:"user_to"`
}
