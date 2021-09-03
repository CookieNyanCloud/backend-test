package v1

import (
	"github.com/cookienyancloud/avito-backend-test/pkg/logger"
	"github.com/gin-gonic/gin"
)

type response struct {
	Message string `json:"message"`
}

type BalanceResponse struct {
	Balance string `json:"balanceResponse"`
	Cur     string `json:"cur"`
}

func newResponse(c *gin.Context, statusCode int, message string) {
	logger.Error(message)
	c.AbortWithStatusJSON(statusCode, response{message})
}
