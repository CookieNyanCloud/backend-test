package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	//"github.com/google/uuid"
	"net/http"
)



func (h *Handler) initFinanceRoutes(api *gin.RouterGroup) {
	operation := api.Group("/operation")
	{
		operation.POST("/transaction", h.transaction)
		operation.POST("/remittance", h.remittance)
		operation.GET("/balance", h.balance)

	}
}

const (
	Success = "удачная транзакция"
)


type TransactionInput struct {
	Id  int `json:"id" binding:"required"`
	Sum float64    `json:"sum" binding:"required"`
}

type RemittanceInput struct {
	IdFrom int `json:"id_from" binding:"required"`
	IdTo   int `json:"id_to" binding:"required"`
	Sum    float64    `json:"sum" binding:"required"`
}

type BalanceInput struct {
	Id  int `json:"id" binding:"required"`
}


func (h *Handler) transaction(c *gin.Context) {

	var inp TransactionInput
	//проверка данных для структуры
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "неверные данные")
		return
	}

	//передача данных

	if err := h.services.Transaction(inp.Id, inp.Sum);err!= nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, response{Success})

}

func (h *Handler) remittance(c *gin.Context) {
	var inp RemittanceInput
	//проверка данных для структуры
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "неверные данные")
		return
	}

	//передача данных
	if err := h.services.Remittance(inp.IdFrom,inp.IdTo,inp.Sum);err!= nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, response{Success})

}

func (h *Handler) balance(c *gin.Context) {
	cur:=c.DefaultQuery("currency", "RUB")
	var inp BalanceInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "неверные данные")
		return
	}
	balance ,err:= h.services.Balance(inp.Id)
	if err!=nil{
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if cur == "RUB" {
		c.JSON(http.StatusOK, BalanceResponse{
			Balance: fmt.Sprintf("%.2f", balance),
			Cur:     cur,
		})
		return
	}

	balanceInCur, err:= h.curService.GetCur(cur,balance)
	if err!=nil{
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, BalanceResponse{
		Balance:balanceInCur,
		Cur:     cur,
	})
}
