package v1

import (
	"github.com/gin-gonic/gin"
	//"github.com/google/uuid"
	"net/http"
)

//todo:uuid
func (h *Handler) initFinanceRoutes(api *gin.RouterGroup) {
	operation := api.Group("/operation")
	{
		operation.POST("/transaction", h.transaction)
		operation.POST("/remittance", h.remittance)
		operation.GET("/balance", h.balance)

	}
}

type TransactionInput struct {
	Id  int `json:"id" binding:"required"`
	Sum float64    `json:"sum" binding:"required"`
}

type RemittanceInput struct {
	IdFrom int `json:"id_from" binding:"required"`
	IdTo   int `json:"id_to" binding:"required"`
	Sum    string    `json:"sum" binding:"required"`
}

type BalanceInput struct {
	Id  int `json:"id" binding:"required"`
	Cur string    `json:"cur" binding:"required"`
}


func (h *Handler) transaction(c *gin.Context) {

	var inp TransactionInput
	//проверка данных для структуры
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "неверные данные")
		return
	}
	//передача данных
	var message string
	if err, message := h.services.Transaction(c, inp.Id, inp.Sum); err!= nil{
		newResponse(c, http.StatusInternalServerError, message)
		return
	}
	c.JSON(http.StatusOK, response{message})

}

func (h *Handler) remittance(c *gin.Context) {

}

func (h *Handler) balance(c *gin.Context) {

}
