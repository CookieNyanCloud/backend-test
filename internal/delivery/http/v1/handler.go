package v1

import (
	delivery "github.com/cookienyancloud/avito-backend-test/internal/delivery/http"
	"github.com/cookienyancloud/avito-backend-test/internal/domain"
	"github.com/cookienyancloud/avito-backend-test/pkg/logger"
	"github.com/gin-gonic/gin"
)

type handler struct {
	services   delivery.IService
	curService delivery.ICurrency
	cache      delivery.ICache
}

//new handler instance
func NewHandler(services delivery.IService, curService delivery.ICurrency, cache delivery.ICache) *handler {
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
	c.AbortWithStatusJSON(statusCode, domain.Response{message})
}
