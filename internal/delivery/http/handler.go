package http

import (
	"context"

	"github.com/cookienyancloud/avito-backend-test/internal/config"
	v1 "github.com/cookienyancloud/avito-backend-test/internal/delivery/http/v1"
	"github.com/cookienyancloud/avito-backend-test/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//go:generate mockgen -source=handler.go -destination=mocks/ServiceMock.go

type IService interface {
	MakeTransaction(ctx context.Context, inp *domain.TransactionInput) error
	MakeRemittance(ctx context.Context, inp *domain.RemittanceInput) error
	GetBalance(ctx context.Context, inp *domain.BalanceInput) (float64, error)
	GetTransactionsList(ctx context.Context, inp *domain.TransactionsListInput) ([]domain.TransactionsListResponse, error)
}

//currency interface
type ICurrency interface {
	GetCur(cur string, sum float64) (string, error)
}

type ICache interface {
	CacheKey(ctx context.Context, key uuid.UUID) error
	CheckKey(ctx context.Context, key uuid.UUID) (bool, error)
}

type handler struct {
	service    IService
	curService ICurrency
	cache      ICache
}

//new handler struct
func NewHandler(service IService, curService ICurrency, cache ICache) *handler {
	return &handler{
		service:    service,
		curService: curService,
		cache:      cache,
	}
}

//initiate gin
func (h *handler) Init(cfg *config.Config) *gin.Engine {
	router := gin.Default()
	router.Use(
		corsMiddleware,
	)
	h.initAPI(router)
	return router
}

func (h *handler) initAPI(router *gin.Engine) {

	handlerV1 := v1.NewHandler(h.service, h.curService, h.cache)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
