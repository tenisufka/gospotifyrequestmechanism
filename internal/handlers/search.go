package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Search(c *gin.Context) {

	c.HTML(
		http.StatusOK,
		"search.html",
		gin.H{
			"title": "Search",
		},
	)
}
