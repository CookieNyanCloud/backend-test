package v1

import (
	"github.com/cookienyancloud/avito-backend-test/internal/service"
	"github.com/gin-gonic/gin"
)

//хэндлер с двумя сервисами, работы с финансами и работы с курсом
type Handler struct {
	services   *service.FinanceService
	curService service.CurService
}

func NewHandler(services *service.FinanceService, curService service.CurService) *Handler {
	return &Handler{
		services:   services,
		curService: curService,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initFinanceRoutes(v1)
	}
}
