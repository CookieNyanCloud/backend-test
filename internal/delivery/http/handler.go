package http

import (
	"github.com/cookienyancloud/avito-backend-test/internal/config"
	v1 "github.com/cookienyancloud/avito-backend-test/internal/delivery/http/v1"
	"github.com/cookienyancloud/avito-backend-test/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service     *service.FinanceService
	curService  service.CurService
}

func NewHandler(service *service.FinanceService, curService service.CurService) *Handler {
	return &Handler{
		service: service,
		curService:curService,
	}
}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	router.Use(
		//gin.Recovery(),
		//gin.Logger(),
		//limiter.Limit(cfg.Limiter.RPS, cfg.Limiter.Burst, cfg.Limiter.TTL),
		corsMiddleware,
	)

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.service,h.curService)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}