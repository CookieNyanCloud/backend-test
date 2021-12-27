package domain

import (
	"database/sql"
	"time"
)

//general response struct
type Response struct {
	Message string `json:"message"`
}

//response balance and currency
type BalanceResponse struct {
	Balance string `json:"balanceResponse"`
	Cur     string `json:"cur"`
}

//response list of transactions
type TransactionsListResponse struct {
	Id          int           `json:"id" db:"user_id"`
	Operation   string        `json:"operation"db:"operation"`
	Sum         string        `json:"sum" db:"sum"`
	Date        time.Time     `json:"date" db:"date"`
	Description string        `json:"description,omitempty" db:"description"`
	IdTo        sql.NullInt64 `json:"id_to,omitempty" db:"user_to"`
}

//currency api struct
type CurrencyResponse struct {
	Success   bool                   `json:"success"`
	Timestamp int64                  `json:"timestamp"`
	Base      string                 `json:"base"`
	Date      string                 `json:"date"`
	Rates     map[string]interface{} `json:"rates"`
}
