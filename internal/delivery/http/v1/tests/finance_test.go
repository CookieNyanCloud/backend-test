package tests

import (
	"bytes"
	"net/http/httptest"
	"testing"
	"time"

	mock_redis "github.com/cookienyancloud/avito-backend-test/internal/cache/redis/mocks"
	"github.com/cookienyancloud/avito-backend-test/internal/delivery/http/v1"
	"github.com/cookienyancloud/avito-backend-test/internal/domain"
	mock_service "github.com/cookienyancloud/avito-backend-test/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTransaction(t *testing.T) {

	type mockBehavior func(s *mock_service.MockIService, inp *domain.TransactionInput)

	tt := []struct {
		name          string
		inpBody       string
		input         *domain.TransactionInput
		mockB         mockBehavior
		expStatusCode int
		expReqBody    string
	}{
		{
			name:    "ok",
			inpBody: `{"id":"1993f8f2-d580-4fb1-bd8e-5bdfb7ddd7e4","sum":10,"description":"test ok"}`,
			input: &domain.TransactionInput{
				Id:          uuid.MustParse("1993f8f2-d580-4fb1-bd8e-5bdfb7ddd7e4"),
				Sum:         10,
				Description: "test ok",
			},
			mockB: func(s *mock_service.MockIService, inp *domain.TransactionInput) {
				s.
					EXPECT().
					MakeTransaction(gomock.Any(), inp).
					Return(nil)
			},
			expStatusCode: 200,
			expReqBody:    `{"message":"удачная транзакция"}`,
		},
		{
			name:    "wrong id",
			inpBody: `{"id":"1993f8f2-d580-4fb1","sum":10,"description":"test ok"}`,
			input: &domain.TransactionInput{
				Id:          uuid.Nil,
				Sum:         10,
				Description: "test wrong id",
			},
			mockB: func(s *mock_service.MockIService, inp *domain.TransactionInput) {
			},
			expStatusCode: 400,
			expReqBody:    `{"message":"неверные данные"}`,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			finService := mock_service.NewMockIService(c)
			subService := mock_service.NewMockICurrency(c)
			cache := mock_redis.NewMockICache(c)
			w := httptest.NewRecorder()
			r := gin.New()
			tc.mockB(finService, tc.input)
			handler := v1.NewHandler(finService, subService, cache)
			r.POST("/transaction", handler.Transaction)
			req := httptest.NewRequest("POST", "/transaction",
				bytes.NewBufferString(tc.inpBody))
			r.ServeHTTP(w, req)
			assert.Equal(t, tc.expStatusCode, w.Code)
			assert.Equal(t, tc.expReqBody, w.Body.String())

		})
	}

}

func TestRemittance(t *testing.T) {

	type mockBehavior func(s *mock_service.MockIService, inp *domain.RemittanceInput)

	tt := []struct {
		name          string
		inpBody       string
		input         *domain.RemittanceInput
		mockB         mockBehavior
		expStatusCode int
		expReqBody    string
	}{
		{
			name:    "ok",
			inpBody: `{"id_from":"1993f8f2-d580-4fb1-bd8e-5bdfb7ddd7e4","id_to":"2993f8f2-d580-4fb1-bd8e-5bdfb7ddd7e4","sum":10,"description":"test ok"}`,
			input: &domain.RemittanceInput{
				IdFrom:      uuid.MustParse("1993f8f2-d580-4fb1-bd8e-5bdfb7ddd7e4"),
				IdTo:        uuid.MustParse("2993f8f2-d580-4fb1-bd8e-5bdfb7ddd7e4"),
				Sum:         10,
				Description: "test ok",
			},
			mockB: func(s *mock_service.MockIService, inp *domain.RemittanceInput) {
				s.
					EXPECT().
					MakeRemittance(gomock.Any(), inp).
					Return(nil)
			},
			expStatusCode: 200,
			expReqBody:    `{"message":"удачная транзакция"}`,
		},
		{
			name:    "wrong id",
			inpBody: `{"id_from":"1993f8f2-d580-4fb1-bd8e-5","id_to":"2993f8f2-d580-4fb1-bd8e-5bdfb7ddd7e4","sum":10,"description":"test ok"}`,
			input: &domain.RemittanceInput{
				IdFrom:      uuid.Nil,
				IdTo:        uuid.MustParse("2993f8f2-d580-4fb1-bd8e-5bdfb7ddd7e4"),
				Sum:         10,
				Description: "test wrong id",
			},
			mockB: func(s *mock_service.MockIService, inp *domain.RemittanceInput) {
			},
			expStatusCode: 400,
			expReqBody:    `{"message":"неверные данные"}`,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			finService := mock_service.NewMockIService(c)
			subService := mock_service.NewMockICurrency(c)
			cache := mock_redis.NewMockICache(c)
			w := httptest.NewRecorder()
			r := gin.New()
			tc.mockB(finService, tc.input)
			handler := v1.NewHandler(finService, subService, cache)
			r.POST("/remittance", handler.Remittance)
			req := httptest.NewRequest("POST", "/remittance",
				bytes.NewBufferString(tc.inpBody))
			r.ServeHTTP(w, req)
			assert.Equal(t, tc.expStatusCode, w.Code)
			assert.Equal(t, tc.expReqBody, w.Body.String())

		})
	}

}

func TestBalance(t *testing.T) {

	type mockBehavior func(s *mock_service.MockIService, c *mock_service.MockICurrency, inp *domain.BalanceInput)

	tt := []struct {
		name          string
		inpBody       string
		cur           string
		input         *domain.BalanceInput
		mockB         mockBehavior
		expStatusCode int
		expReqBody    string
	}{
		{
			name:    "ok rub",
			inpBody: `{"id":"1993f8f2-d580-4fb1-bd8e-5bdfb7ddd7e4"}`,
			cur:     "",
			input: &domain.BalanceInput{
				Id: uuid.MustParse("1993f8f2-d580-4fb1-bd8e-5bdfb7ddd7e4"),
			},
			mockB: func(s *mock_service.MockIService, c *mock_service.MockICurrency, inp *domain.BalanceInput) {
				s.
					EXPECT().
					GetBalance(gomock.Any(), inp).
					Return(100.1, nil)
			},
			expStatusCode: 200,
			expReqBody:    `{"balanceResponse":"₽100.10","cur":"RUB"}`,
		},
		{
			name:    "ok USD",
			inpBody: `{"id":"1993f8f2-d580-4fb1-bd8e-5bdfb7ddd7e4"}`,
			cur:     "?currency=USD",
			input: &domain.BalanceInput{
				Id: uuid.MustParse("1993f8f2-d580-4fb1-bd8e-5bdfb7ddd7e4"),
			},
			mockB: func(s *mock_service.MockIService, c *mock_service.MockICurrency, inp *domain.BalanceInput) {
				s.
					EXPECT().
					GetBalance(gomock.Any(), inp).
					Return(100.1, nil)
				c.
					EXPECT().
					GetCur("USD", 100.1).
					Return("$12", nil)
			},
			expStatusCode: 200,
			expReqBody:    `{"balanceResponse":"$12","cur":"USD"}`,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			finService := mock_service.NewMockIService(c)
			subService := mock_service.NewMockICurrency(c)
			cache := mock_redis.NewMockICache(c)
			w := httptest.NewRecorder()
			r := gin.New()
			tc.mockB(finService, subService, tc.input)
			handler := v1.NewHandler(finService, subService, cache)
			r.GET("/balance", handler.Balance)
			req := httptest.NewRequest("GET", "/balance"+tc.cur,
				bytes.NewBufferString(tc.inpBody))
			r.ServeHTTP(w, req)
			assert.Equal(t, tc.expStatusCode, w.Code)
			assert.Equal(t, tc.expReqBody, w.Body.String())

		})
	}

}

func TestTransactionsList(t *testing.T) {

	type mockBehavior func(s *mock_service.MockIService, inp *domain.TransactionsListInput)

	tt := []struct {
		name          string
		inpBody       string
		input         *domain.TransactionsListInput
		mockB         mockBehavior
		expStatusCode int
		expReqBody    string
	}{
		{
			name:    "ok",
			inpBody: `{"id":"1993f8f2-d580-4fb1-bd8e-5bdfb7ddd7e4"}`,
			input: &domain.TransactionsListInput{
				Id:   uuid.MustParse("1993f8f2-d580-4fb1-bd8e-5bdfb7ddd7e4"),
				Sort: "sum",
				Dir:  "asc",
				Page: 1,
			},
			mockB: func(s *mock_service.MockIService, inp *domain.TransactionsListInput) {
				parse, _ := time.Parse("2006-01-02T15:04:05.000Z", "0001-01-01T00:00:00Z")
				s.
					EXPECT().
					GetTransactionsList(gomock.Any(), inp).
					Return([]domain.TransactionsList{domain.TransactionsList{
						Id:          uuid.UUID{},
						Operation:   "remittance",
						Sum:         10,
						Date:        parse,
						Description: "",
						IdTo:        uuid.MustParse("1993f8f2-d580-4fb1-bd8e-5bdfb7ddd7e4"),
					}}, nil)
			},
			expStatusCode: 200,
			expReqBody:    `[{"id":"00000000-0000-0000-0000-000000000000","operation":"remittance","sum":10,"date":"0001-01-01T00:00:00Z","id_to":"1993f8f2-d580-4fb1-bd8e-5bdfb7ddd7e4"}]`,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			finService := mock_service.NewMockIService(c)
			subService := mock_service.NewMockICurrency(c)
			cache := mock_redis.NewMockICache(c)
			w := httptest.NewRecorder()
			r := gin.New()
			tc.mockB(finService, tc.input)
			handler := v1.NewHandler(finService, subService, cache)
			r.GET("/transactionsList", handler.TransactionsList)
			req := httptest.NewRequest("GET", "/transactionsList?sort=sum&dir=asc&page=1",
				bytes.NewBufferString(tc.inpBody))
			r.ServeHTTP(w, req)
			assert.Equal(t, tc.expStatusCode, w.Code)
			assert.Equal(t, tc.expReqBody, w.Body.String())

		})
	}

}
