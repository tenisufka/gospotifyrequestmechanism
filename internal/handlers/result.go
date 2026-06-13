package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"spotifysrmechanism/internal/filters"
)

func (h *Handler) Result(c *gin.Context) {

	song := c.PostForm("song")

	if song == "" {
		c.HTML(http.StatusBadRequest, "wrong.html", gin.H{
			"error": "song is empty",
		})
		return
	}

	ctx := c.Request.Context()

	// Inicjalizacja i sprawdzenie klienta Spotify dla hosta
	spotifyClient := h.spotifyClient()
	if spotifyClient == nil {
		c.HTML(http.StatusForbidden, "wrong.html", gin.H{
			"error": "host is not logged into spotify",
		})
		return
	}

	// Wyszukanie utworu za pomocą pobranego klienta
	track, err := spotifyClient.SearchTrack(
		ctx,
		song,
	)

	if err != nil {
		c.HTML(http.StatusNotFound, "wrong.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	if track.Explicit {
		c.HTML(http.StatusForbidden, "wrong.html", gin.H{
			"error": "explicit track",
		})
		return
	}

	durationSeconds := int(track.Duration) / 1000

	if durationSeconds > 300 {
		c.HTML(http.StatusForbidden, "wrong.html", gin.H{
			"error": "track longer than 5 minutes",
		})
		return
	}

	if len(track.Artists) == 0 {
		c.HTML(http.StatusInternalServerError, "wrong.html", gin.H{
			"error": "track has no artists",
		})
		return
	}

	artist := track.Artists[0].Name

	lyricsText, err := h.lyrics.GetLyrics(
		ctx,
		artist,
		track.Name,
	)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "wrong.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	if filters.ContainsBadWords(lyricsText) {
		c.HTML(http.StatusForbidden, "wrong.html", gin.H{
			"error": "lyrics contain forbidden words",
		})
		return
	}

	// Pobranie aktywnego urządzenia za pomocą pobranego klienta
	device, err := spotifyClient.ActiveDevice(ctx)

	if err != nil {
		c.HTML(http.StatusBadRequest, "wrong.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	// Dodanie do kolejki za pomocą pobranego klienta
	err = spotifyClient.AddToQueue(
		ctx,
		track.ID,
		string(device.ID),
	)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "wrong.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "result.html", gin.H{
		"track":  track.Name,
		"lyrics": lyricsText,
	})
}
