package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cookienyancloud/avito-backend-test/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *handler) initFinanceRoutes(api *gin.RouterGroup) {
	operation := api.Group("/operation")
	{
		operation.POST("/transaction", h.Transaction)
		operation.POST("/remittance", h.Remittance)
		operation.GET("/balance", h.Balance)
		operation.GET("/transactionsList", h.TransactionsList)

	}
}

const (
	success   = "удачная транзакция"
	userFail  = "неверные данные"
	cacheFail = "ошибка на стороне кеша"
	duplicate = "повторный запрос"
)

//handle user transactions request
func (h *handler) Transaction(c *gin.Context) {
	var inp domain.TransactionInput
	if err := c.BindJSON(&inp); err != nil {
		h.newResponse(c, http.StatusBadRequest, userFail, err)
		return
	}
	if is := h.CheckCache(c, inp.IdempotencyKey); is {
		return
	}
	if err := h.services.MakeTransaction(c, &inp); err != nil {
		h.newResponse(c, http.StatusBadRequest, userFail, err)
		return
	}
	c.JSON(http.StatusOK, domain.Response{success})

}

//handle transactions request from user to user
func (h *handler) Remittance(c *gin.Context) {
	var inp domain.RemittanceInput
	if err := c.BindJSON(&inp); err != nil {
		h.newResponse(c, http.StatusBadRequest, userFail, err)
		return
	}

	if is := h.CheckCache(c, inp.IdempotencyKey); is {
		return
	}
	if err := h.services.MakeRemittance(c, &inp); err != nil {
		h.newResponse(c, http.StatusBadRequest, userFail, err)
		return
	}
	c.JSON(http.StatusOK, domain.Response{success})

}

//handle check balance
func (h *handler) Balance(c *gin.Context) {
	cur := c.DefaultQuery("currency", "RUB")
	var inp domain.BalanceInput
	if err := c.BindJSON(&inp); err != nil {
		h.newResponse(c, http.StatusBadRequest, userFail, err)
		return
	}

	balance, err := h.services.GetBalance(c, &inp)
	if err != nil {
		h.newResponse(c, http.StatusNotFound, userFail, err)
		return
	}
	if cur == "RUB" {
		c.JSON(http.StatusOK, domain.BalanceResponse{
			Balance: fmt.Sprintf("₽%.2f", balance),
			Cur:     cur,
		})
		return
	}

	balanceInCur, err := h.curService.GetCur(cur, balance)
	if err != nil {
		h.newResponse(c, http.StatusBadRequest, userFail, err)
		return
	}

	c.JSON(http.StatusOK, domain.BalanceResponse{
		Balance: balanceInCur,
		Cur:     cur,
	})
}

//handle check all transactions by query
func (h *handler) TransactionsList(c *gin.Context) {

	var inp domain.TransactionsListInput
	if err := c.BindJSON(&inp); err != nil {
		h.newResponse(c, http.StatusBadRequest, userFail, err)
		return
	}

	inp.Sort = c.DefaultQuery("sort", "date")
	inp.Dir = c.DefaultQuery("dir", "ASC")
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		h.newResponse(c, http.StatusInternalServerError, userFail, err)
		return
	}
	inp.Page = page
	list, err := h.services.GetTransactionsList(c, &inp)
	if err != nil {
		h.newResponse(c, http.StatusBadRequest, userFail, err)
		return
	}

	c.JSON(http.StatusOK, list)
}

//check request in cache by key
func (h *handler) CheckCache(c *gin.Context, key uuid.UUID) bool {
	state, err := h.cache.CheckKey(c, key)
	if err != nil {
		h.newResponse(c, http.StatusInternalServerError, cacheFail, err)
		return true
	}
	if state == true {
		h.newResponse(c, http.StatusConflict, duplicate, nil)
		return true
	} else {
		if err := h.cache.CacheKey(c, key); err != nil {
			h.newResponse(c, http.StatusInternalServerError, cacheFail, err)
			return true
		}
	}

	return false
}
