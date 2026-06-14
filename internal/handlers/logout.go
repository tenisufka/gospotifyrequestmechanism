package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Logout(c *gin.Context) {
	h.tokenStore.Set(nil)

	c.Redirect(
		http.StatusTemporaryRedirect,
		"/",
	)
}
