package httprest

import (
	"github.com/cookienyancloud/avito-backend-test/internal/config"
	v1 "github.com/cookienyancloud/avito-backend-test/internal/delivery/httprest/v1"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service    v1.IService
	curService v1.ICurrency
	cache      v1.ICache
}

//new handler struct
func NewHandler(service v1.IService, curService v1.ICurrency, cache v1.ICache) *handler {
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
