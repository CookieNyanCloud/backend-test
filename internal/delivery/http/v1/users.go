package v1

import "github.com/gin-gonic/gin"

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	operation := api.Group("/operation")
	{
		operation.POST("/transaction", h.transaction)
		operation.POST("/remittance", h.remittance)
		operation.GET("/balance", h.balance)

	}
}

func (h *Handler) transaction(c *gin.Context) {

}

func (h *Handler) remittance(c *gin.Context) {

}

func (h *Handler) balance(c *gin.Context) {

}





