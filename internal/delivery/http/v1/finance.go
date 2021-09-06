package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) initFinanceRoutes(api *gin.RouterGroup) {
	//пути к операциям
	operation := api.Group("/operation")
	{
		operation.POST("/transaction", h.transaction)
		operation.POST("/remittance", h.remittance)
		operation.GET("/balance", h.balance)
		operation.GET("/transactionsList", h.transactionsList)

	}
}

const (
	Success = "удачная транзакция"
)

type transactionInput struct {
	Id          int     `json:"id" binding:"required,gt=0"`
	Sum         float64 `json:"sum" binding:"required"`
	Description string  `json:"description" binding:"max=20"`
}

type remittanceInput struct {
	IdFrom      int     `json:"id_from" binding:"required,gt=0"`
	IdTo        int     `json:"id_to" binding:"required,gt=0"`
	Sum         float64 `json:"sum" binding:"required,gt=0"`
	Description string  `json:"description" binding:"max=20"`
}

type balanceInput struct {
	Id int `json:"id" binding:"required,gt=0"`
}

type transactionsListInput struct {
	Id int `json:"id" binding:"required,gt=0"`
}

func (h *Handler) transaction(c *gin.Context) {

	var inp transactionInput
	//проверка данных для структуры
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "неверные данные")
		return
	}

	//передача данных
	if err := h.services.Transaction(inp.Id, inp.Sum, inp.Description); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, response{Success})

}

func (h *Handler) remittance(c *gin.Context) {
	var inp remittanceInput
	//проверка данных для структуры
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "неверные данные")
		return
	}

	//передача данных
	if err := h.services.Remittance(inp.IdFrom, inp.IdTo, inp.Sum, inp.Description); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, response{Success})

}

func (h *Handler) balance(c *gin.Context) {
	cur := c.DefaultQuery("currency", "RUB")
	var inp balanceInput
	//проверка данных для структуры
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "неверные данные")
		return
	}
	//передача данных
	balance, err := h.services.Balance(inp.Id)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if cur == "RUB" {
		c.JSON(http.StatusOK, BalanceResponse{
			Balance: fmt.Sprintf("₽%.2f", balance),
			Cur:     cur,
		})
		return
	}

	balanceInCur, err := h.curService.GetCur(cur, balance)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, BalanceResponse{
		Balance: balanceInCur,
		Cur:     cur,
	})
}

func (h *Handler) transactionsList(c *gin.Context) {

	sort := c.DefaultQuery("sort", "date")
	dir := c.DefaultQuery("dir", "ASC")
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var inp transactionsListInput
	//проверка данных для структуры
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "неверные данные")
		return
	}
	//передача данных
	list, err := h.services.GetTransactionsList(inp.Id, sort, dir, page)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}
