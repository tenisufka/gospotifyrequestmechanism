package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Index(c *gin.Context) {

	client := h.spotifyClient()

	if client == nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":          "Spotify Request Mechanism",
			"PlaybackName":   "",
			"PlaybackArtist": "",
			"PlaybackImage":  "",
			"PlaybackEx":     false,
		})
		return
	}

	current, err := client.CurrentlyPlaying(c.Request.Context())

	if err != nil || current == nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":          "Spotify Request Mechanism",
			"PlaybackName":   "",
			"PlaybackArtist": "",
			"PlaybackImage":  "",
			"PlaybackEx":     false,
		})
		return
	}

	image := ""
	if len(current.Album.Images) > 0 {
		image = current.Album.Images[0].URL
	}

	artist := ""
	if len(current.Artists) > 0 {
		artist = current.Artists[0].Name
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":          "Spotify Request Mechanism",
		"PlaybackName":   current.Name,
		"PlaybackArtist": artist,
		"PlaybackImage":  image,
		"PlaybackEx":     current.Explicit,
	})
}
