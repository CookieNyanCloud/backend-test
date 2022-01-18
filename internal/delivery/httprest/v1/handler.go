package v1

import (
	"context"

	"github.com/cookienyancloud/avito-backend-test/internal/domain"
	"github.com/cookienyancloud/avito-backend-test/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//go:generate mockgen -source=handler.go -destination=mocks/serviceMock.go

//service layer
type IService interface {
	MakeTransaction(ctx context.Context, inp *domain.TransactionInput) error
	MakeRemittance(ctx context.Context, inp *domain.RemittanceInput) error
	GetBalance(ctx context.Context, inp *domain.BalanceInput) (float64, error)
	GetTransactionsList(ctx context.Context, inp *domain.TransactionsListInput) ([]domain.TransactionsListResponse, error)
}

//currency service layer
type ICurrency interface {
	GetCur(cur string, sum float64) (string, error)
}

//cache for middleware
type ICache interface {
	CacheKey(ctx context.Context, key uuid.UUID) error
	CheckKey(ctx context.Context, key uuid.UUID) (bool, error)
}

type handler struct {
	services   IService
	curService ICurrency
	cache      ICache
}

//new handler instance
func NewHandler(services IService, curService ICurrency, cache ICache) *handler {
	return &handler{
		services:   services,
		curService: curService,
		cache:      cache,
	}
}

func (h *handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initFinanceRoutes(v1)
	}
}

func (h *handler) newResponse(c *gin.Context, statusCode int, message string, err error) {
	if err != nil {
		logger.Error(err)
	}
	c.AbortWithStatusJSON(statusCode, domain.Response{Message: message})
}
