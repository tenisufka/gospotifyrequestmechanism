package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
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
		c.HTML(
			http.StatusInternalServerError,
			"wrong.html",
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	// Zapis tokenu do sesji
	session := sessions.Default(c)

	session.Set(
		"spotify_token",
		token.AccessToken,
	)

	if err := session.Save(); err != nil {
		c.HTML(
			http.StatusInternalServerError,
			"wrong.html",
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.Redirect(
		http.StatusTemporaryRedirect,
		"/",
	)
}
