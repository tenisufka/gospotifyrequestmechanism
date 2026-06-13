package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {

	url := h.oauth.AuthURL()

	c.Redirect(
		http.StatusTemporaryRedirect,
		url,
	)
}
