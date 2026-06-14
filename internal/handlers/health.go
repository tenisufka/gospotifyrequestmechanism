package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Health(c *gin.Context) {
	client := h.spotifyClient()

	if client == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "spotify not connected",
			"reason": "host has not authenticated",
		})
		return
	}

	device, err := client.ActiveDevice(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "spotify unavailable",
			"reason": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"device": device.Name,
		"type":   device.Type,
		"active": device.Active,
	})
}
