package v1

import (
	"github.com/cookienyancloud/avito-backend-test/internal/cache/redis"
	"github.com/cookienyancloud/avito-backend-test/internal/domain"
	"github.com/cookienyancloud/avito-backend-test/internal/service"
	"github.com/cookienyancloud/avito-backend-test/pkg/logger"
	"github.com/gin-gonic/gin"
)

type handler struct {
	services   service.IService
	curService service.ICurrency
	cache      redis.ICache
}

//new handler instance
func NewHandler(services service.IService, curService service.ICurrency, cache redis.ICache) *handler {
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
