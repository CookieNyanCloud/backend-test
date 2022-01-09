package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cookienyancloud/avito-backend-test/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *handler) initFinanceRoutes(api *gin.RouterGroup) {
	operation := api.Group("/operation")
	{
		toBeCached := operation.Group("/", h.CheckCache)
		{
			toBeCached.POST("/transaction", h.Transaction)
			toBeCached.POST("/remittance", h.Remittance)

		}
		operation.GET("/balance", h.Balance)
		operation.GET("/transactionsList", h.TransactionsList)

	}
}

const (
	success    = "удачная транзакция"
	userFail   = "неверные данные"
	keyFail    = "ошибка ключа"
	cacheFail  = "ошибка на стороне кеша"
	duplicate  = "повторный запрос"
	cacheState = "cache-state"
)

//handle user transactions request
func (h *handler) Transaction(c *gin.Context) {
	state := c.GetBool("cache-state")
	if !state {
		return
	}
	var inp domain.TransactionInput
	if err := c.BindJSON(&inp); err != nil {
		h.newResponse(c, http.StatusBadRequest, userFail, err)
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
	state := c.GetBool("cache-state")
	if !state {
		return
	}
	var inp domain.RemittanceInput
	if err := c.BindJSON(&inp); err != nil {
		h.newResponse(c, http.StatusBadRequest, userFail, err)
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

	if ok := listInputCheck(&inp); !ok {
		h.newResponse(c, http.StatusBadRequest, userFail, nil)
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

func listInputCheck(inp *domain.TransactionsListInput) bool {
	switch inp.Sort {
	case "sum", "date":
	default:
		return false
	}

	switch inp.Dir {
	case "asc", "desc":
	default:
		return false
	}
	return true
}
