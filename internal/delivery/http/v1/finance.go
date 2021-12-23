package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cookienyancloud/avito-backend-test/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	success   = "удачная транзакция"
	userFail  = "неверные данные"
	cacheFail = "ошибка на стороне кеша"
	duplicate = "повторный запрос"
)

func (h *Handler) transaction(c *gin.Context) {

	var inp domain.TransactionInput
	if err := c.BindJSON(&inp); err != nil {
		h.newResponse(c, http.StatusBadRequest, userFail, err)
		return
	}
	if ok := h.CheckCache(c, inp.IdempotencyKey); !ok {
		return
	}
	if err := h.Services.MakeTransaction(c, &inp); err != nil {
		h.newResponse(c, http.StatusBadRequest, userFail, err)
		return
	}
	c.JSON(http.StatusOK, domain.Response{success})

}

func (h *Handler) remittance(c *gin.Context) {
	var inp domain.RemittanceInput
	if err := c.BindJSON(&inp); err != nil {
		h.newResponse(c, http.StatusBadRequest, userFail, err)
		return
	}

	if ok := h.CheckCache(c, inp.IdempotencyKey); !ok {
		return
	}
	if err := h.Services.MakeRemittance(c, &inp); err != nil {
		h.newResponse(c, http.StatusNotFound, userFail, err)
		return
	}
	c.JSON(http.StatusOK, domain.Response{success})

}

func (h *Handler) balance(c *gin.Context) {
	cur := c.DefaultQuery("currency", "RUB")
	var inp domain.BalanceInput
	if err := c.BindJSON(&inp); err != nil {
		h.newResponse(c, http.StatusBadRequest, userFail, err)
		return
	}

	balance, err := h.Services.GetBalance(c, &inp)
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

	balanceInCur, err := h.CurService.GetCur(cur, balance)
	if err != nil {
		h.newResponse(c, http.StatusBadRequest, userFail, err)
		return
	}

	c.JSON(http.StatusOK, domain.BalanceResponse{
		Balance: balanceInCur,
		Cur:     cur,
	})
}

func (h *Handler) transactionsList(c *gin.Context) {

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
	list, err := h.Services.GetTransactionsList(c, &inp)
	if err != nil {
		h.newResponse(c, http.StatusBadRequest, userFail, err)
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) CheckCache(c *gin.Context, key uuid.UUID) bool {
	state, err := h.cache.CheckKey(c, key)
	if err != nil {
		h.newResponse(c, http.StatusInternalServerError, cacheFail, err)
		return false
	}
	if state == true {
		h.newResponse(c, http.StatusConflict, duplicate, err)
		return false
	}
	if err := h.cache.CacheKey(c, key); err != nil {
		h.newResponse(c, http.StatusInternalServerError, cacheFail, err)
		return false
	}
	return true
}
