package v1

import (
	"github.com/cookienyancloud/avito-backend-test/pkg/logger"
	"github.com/gin-gonic/gin"
	"time"
)

type response struct {
	Message string `json:"message"`
}

type BalanceResponse struct {
	Balance string `json:"balanceResponse"`
	Cur     string `json:"cur"`
}
type transactionsListResponse struct {
	Id          int       `json:"id" db:"user_id"`
	Operation   string    `json:"operation"db:"operation"`
	Sum         float64   `json:"sum" db:"sum"`
	Date        time.Time `json:"date" db:"date"`
	Description string    `json:"description" db:"description"`
	IdTo        int       `json:"id_to" db:"user_to"`
}



func newResponse(c *gin.Context, statusCode int, message string) {
	logger.Error(message)
	c.AbortWithStatusJSON(statusCode, response{message})
}
