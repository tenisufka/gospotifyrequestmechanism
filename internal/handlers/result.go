package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify"

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

	// Spotify client (już powinien być autoryzowany middlewarem)
	token := c.GetString("spotify_token")
	if token == "" {
		c.HTML(http.StatusUnauthorized, "wrong.html", gin.H{
			"error": "missing spotify token",
		})
		return
	}

	// Jeśli klient NIE jest jeszcze ustawiony, tworzysz go:
	// (jeśli masz już h.spotify globalnie - pomiń)
	h.spotify = spotify.NewAuthenticator("", nil).NewClient(&spotify.Token{
		AccessToken: token,
	})

	// 1. Search track
	results, err := h.spotify.Search(song, spotify.SearchTypeTrack)
	if err != nil || results.Tracks == nil || len(results.Tracks.Tracks) == 0 {
		c.HTML(http.StatusNotFound, "wrong.html", gin.H{
			"error": "track not found",
		})
		return
	}

	track := results.Tracks.Tracks[0]

	// 2. Validation
	if track.Explicit {
		c.HTML(http.StatusForbidden, "wrong.html", gin.H{
			"error": "explicit track",
		})
		return
	}

	if track.Duration > 300000000000 { // duration in nanoseconds in zmb3/spotify
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

	// 3. Lyrics (blocking call - bez goroutine, bo i tak czekasz)
	artist := track.Artists[0].Name
	lyricsText, err := h.lyrics.GetLyrics(ctx, artist, track.Name)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "wrong.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	// 4. Filter lyrics
	if filters.ContainsBadWords(lyricsText) {
		c.HTML(http.StatusForbidden, "wrong.html", gin.H{
			"error": "lyrics contain forbidden words",
		})
		return
	}

	// 5. Get devices
	devices, err := h.spotify.PlayerDevices()
	if err != nil {
		c.HTML(http.StatusBadRequest, "wrong.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(devices) == 0 {
		c.HTML(http.StatusBadRequest, "wrong.html", gin.H{
			"error": "no active devices",
		})
		return
	}

	deviceID := devices[0].ID

	// 6. Add to queue (IMPORTANT: spotify.URI, nie string)
	err = h.spotify.AddToQueue(deviceID, track.URI)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "wrong.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	// 7. Success
	c.HTML(http.StatusOK, "result.html", gin.H{
		"track":  track.Name,
		"lyrics": lyricsText,
	})
}
