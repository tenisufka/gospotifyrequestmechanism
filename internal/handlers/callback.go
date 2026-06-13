package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Callback(c *gin.Context) {

	code := c.Query("code")

	if code == "" {
		c.HTML(
			http.StatusBadRequest,
			"wrong.html",
			gin.H{
				"error": "spotify code missing",
			},
		)
		return
	}

	token, err := h.oauth.Exchange(
		c.Request.Context(),
		code,
	)
	if err != nil {
		c.HTML(http.StatusBadRequest, "wrong.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	h.tokenStore.Set(token)

	c.Redirect(
		http.StatusTemporaryRedirect,
		"/",
	)
}
