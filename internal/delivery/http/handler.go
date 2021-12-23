package http

import (
	"github.com/cookienyancloud/avito-backend-test/internal/cache/redis"
	"github.com/cookienyancloud/avito-backend-test/internal/config"
	v1 "github.com/cookienyancloud/avito-backend-test/internal/delivery/http/v1"
	"github.com/cookienyancloud/avito-backend-test/internal/service"
	"github.com/gin-gonic/gin"
)

//services handler
type Handler struct {
	service    service.IFinance
	curService service.ICurrency
	cache      redis.ICache
}

func NewHandler(service service.IFinance, curService service.ICurrency, cache redis.ICache) *Handler {
	return &Handler{
		service:    service,
		curService: curService,
		cache:      cache,
	}
}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	router := gin.Default()
	router.Use(
		corsMiddleware,
	)
	h.initAPI(router)
	return router
}

func (h *Handler) initAPI(router *gin.Engine) {

	handlerV1 := v1.NewHandler(h.service, h.curService, h.cache)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
