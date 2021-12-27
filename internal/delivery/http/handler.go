package http

import (
	"github.com/cookienyancloud/avito-backend-test/internal/cache/redis"
	"github.com/cookienyancloud/avito-backend-test/internal/config"
	v1 "github.com/cookienyancloud/avito-backend-test/internal/delivery/http/v1"
	"github.com/cookienyancloud/avito-backend-test/internal/service"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service    service.IFinance
	curService service.ICurrency
	cache      redis.ICache
}

//new handler struct
func NewHandler(service service.IFinance, curService service.ICurrency, cache redis.ICache) *handler {
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
