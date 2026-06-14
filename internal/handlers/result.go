package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"spotifysrmechanism/internal/filters"
)

func (h *Handler) Result(c *gin.Context) {

	spotifyClient := h.spotifyClient()

	if spotifyClient == nil {
		c.HTML(
			http.StatusServiceUnavailable,
			"wrong.html",
			gin.H{
				"error": "host is not connected to spotify",
			},
		)
		return
	}

	song := c.PostForm("song")

	if song == "" {
		c.HTML(
			http.StatusBadRequest,
			"wrong.html",
			gin.H{
				"error": "Nie podano utworu",
			},
		)
		return
	}

	ctx := c.Request.Context()

	track, err := spotifyClient.GetTrackByID(
		ctx,
		song,
	)
	if err != nil {
		c.HTML(
			http.StatusNotFound,
			"wrong.html",
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	imageURL := ""

	if len(track.Album.Images) > 0 {
		imageURL = track.Album.Images[0].URL
	}

	artist := "Unknown Artist"

	if len(track.Artists) > 0 {
		artist = track.Artists[0].Name
	}

	if track.Explicit {
		c.HTML(
			http.StatusForbidden,
			"wrong.html",
			gin.H{
				"error":  "Utwór oznaczony jako Explicit",
				"track":  track.Name,
				"artist": artist,
				"image":  imageURL,
			},
		)
		return
	}

	durationSeconds := int(track.Duration) / 1000

	if durationSeconds > 300 {
		c.HTML(
			http.StatusForbidden,
			"wrong.html",
			gin.H{
				"error":  "Utwór jest dłuższy niż 5 minut",
				"track":  track.Name,
				"artist": artist,
				"image":  imageURL,
			},
		)
		return
	}

	lyricsText, err := h.lyrics.GetLyrics(
		ctx,
		artist,
		track.Name,
	)

	if err != nil {
		// Nie udało się pobrać tekstu po 3 próbach.
		// Logujemy błąd, ale przepuszczamy utwór.
		println("Lyrics check skipped:", err.Error())
	} else {
		if filters.ContainsBadWords(lyricsText) {
			c.HTML(http.StatusForbidden, "wrong.html", gin.H{
				"error":  "Tekst zawiera niedozwolone słowa",
				"track":  track.Name,
				"artist": artist,
				"image":  imageURL,
			})
			return
		}
	}

	device, err := spotifyClient.ActiveDevice(ctx)

	if err != nil {
		c.HTML(
			http.StatusBadRequest,
			"wrong.html",
			gin.H{
				"error":  err.Error(),
				"track":  track.Name,
				"artist": artist,
				"image":  imageURL,
			},
		)
		return
	}

	err = spotifyClient.AddToQueue(
		ctx,
		track.ID,
		string(device.ID),
	)

	if err != nil {
		c.HTML(
			http.StatusInternalServerError,
			"wrong.html",
			gin.H{
				"error":  err.Error(),
				"track":  track.Name,
				"artist": artist,
				"image":  imageURL,
			},
		)
		return
	}

	c.HTML(
		http.StatusOK,
		"result.html",
		gin.H{
			"TrackName":   track.Name,
			"TrackArtist": artist,
			"TrackImage":  imageURL,
			"TrackEx":     track.Explicit,
			"lyrics":      lyricsText,
		},
	)
}
