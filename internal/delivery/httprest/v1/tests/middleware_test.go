package tests

import (
	"net/http/httptest"
	"testing"

	v1 "github.com/cookienyancloud/avito-backend-test/internal/delivery/httprest/v1"
	mock_service "github.com/cookienyancloud/avito-backend-test/internal/delivery/httprest/v1/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCheckCache(t *testing.T) {
	type mockBehavior func(c *mock_service.MockICache)

	tt := []struct {
		name          string
		headerName    string
		headerValue   string
		mockB         mockBehavior
		expStatusCode int
		expResBody    string
	}{
		{
			name:        "OK full way",
			headerName:  "Idempotence-Key",
			headerValue: "1993f8f2-d580-4fb1-bd8e-5bdfb7ddd7e4",
			mockB: func(c *mock_service.MockICache) {
				key := uuid.MustParse("1993f8f2-d580-4fb1-bd8e-5bdfb7ddd7e4")
				c.EXPECT().
					CheckKey(gomock.Any(), key).
					Return(false, nil)
				c.EXPECT().
					CacheKey(gomock.Any(), key).
					Return(nil)
			},
			expStatusCode: 200,
			expResBody:    "1993f8f2-d580-4fb1-bd8e-5bdfb7ddd7e4",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			finService := mock_service.NewMockIService(c)
			subService := mock_service.NewMockICurrency(c)
			cache := mock_service.NewMockICache(c)
			r := gin.New()
			tc.mockB(cache)
			handler := v1.NewHandler(finService, subService, cache)
			r.POST("/protected", handler.CheckCache, func(c *gin.Context) {
				keyStr := c.GetHeader("Idempotence-Key")
				c.String(200, keyStr)
			})
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/protected", nil)
			req.Header.Set(tc.headerName, tc.headerValue)
			r.ServeHTTP(w, req)
			assert.Equal(t, w.Code, tc.expStatusCode)
			assert.Equal(t, w.Body.String(), tc.expResBody)

		})
	}
}
