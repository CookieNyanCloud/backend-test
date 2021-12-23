package v1

import (
	"github.com/cookienyancloud/avito-backend-test/internal/cache/redis"
	"github.com/cookienyancloud/avito-backend-test/internal/domain"
	"github.com/cookienyancloud/avito-backend-test/internal/service"
	"github.com/cookienyancloud/avito-backend-test/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Services   service.IFinance
	CurService service.ICurrency
	cache      redis.ICache
}

func NewHandler(services service.IFinance, curService service.ICurrency, cache redis.ICache) *Handler {
	return &Handler{
		Services:   services,
		CurService: curService,
		cache:      cache,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initFinanceRoutes(v1)
	}
}

func (h *Handler) newResponse(c *gin.Context, statusCode int, message string, err error) {
	logger.Error(err)
	c.AbortWithStatusJSON(statusCode, domain.Response{message})
}
