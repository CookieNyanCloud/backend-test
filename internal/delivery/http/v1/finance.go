package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	//"github.com/google/uuid"
	"net/http"
)



func (h *Handler) initFinanceRoutes(api *gin.RouterGroup) {
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
	Id  int `json:"id" binding:"required"`
	Sum float64    `json:"sum" binding:"required"`
}

type remittanceInput struct {
	IdFrom int `json:"id_from" binding:"required"`
	IdTo   int `json:"id_to" binding:"required"`
	Sum    float64    `json:"sum" binding:"required"`
}

type balanceInput struct {
	Id  int `json:"id" binding:"required"`
}

type transactionsListInput struct {
	Id  int `json:"id" binding:"required"`
}

//type transactionsList struct {
//	Id          int       `json:"id"`
//	Operation   string    `json:"operation"`
//	Sum         float64   `json:"sum"`
//	Date        time.Time `json:"date"`
//	Description string    `json:"description"`
//	IdTo        int       `json:"id_to"`
//}


func (h *Handler) transaction(c *gin.Context) {

	var inp transactionInput
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
	var inp remittanceInput
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
	var inp balanceInput
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
			Balance: fmt.Sprintf("₽%.2f", balance),
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

func (h *Handler) transactionsList(c *gin.Context) {


	sort:=c.DefaultQuery("sort", "date")
	dir:=c.DefaultQuery("dir", "ASC")
	page, err:=strconv.Atoi(c.DefaultQuery("page", "0"))
	if err!=nil{
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var inp transactionsListInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "неверные данные")
		return
	}
	list ,err:= h.services.GetTransactionsList(inp.Id,sort,dir,page)
	if err!=nil{
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}