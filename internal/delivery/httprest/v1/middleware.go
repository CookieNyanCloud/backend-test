package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//check request in cache by key
func (h *handler) CheckCache(c *gin.Context) {
	keyStr := c.GetHeader("Idempotence-Key")
	if keyStr == "" {
		h.newResponse(c, http.StatusBadRequest, keyFail, nil)
		return
	}
	key, err := uuid.Parse(keyStr)
	if err != nil {
		h.newResponse(c, http.StatusBadRequest, keyFail, err)
		return
	}
	state, err := h.cache.CheckKey(c, key)
	if err != nil {
		h.newResponse(c, http.StatusBadRequest, cacheFail, err)
		return
	}
	if state {
		h.newResponse(c, http.StatusConflict, duplicate, nil)
		return
	}
	if err := h.cache.CacheKey(c, key); err != nil {
		h.newResponse(c, http.StatusBadRequest, cacheFail, err)
		return
	}
}
