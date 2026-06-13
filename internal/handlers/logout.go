package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Logout(c *gin.Context) {

	session := sessions.Default(c)

	session.Clear()

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
