package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SearchResults(c *gin.Context) {

	spotifyClient := h.spotifyClient()

	if spotifyClient == nil {
		c.String(http.StatusOK, "spotifyClient == nil")
		return
	}

	song := c.PostForm("song")
	println("SEARCH QUERY:", song)

	if song == "" {
		c.String(http.StatusOK, "song is empty")
		return
	}

	tracks, err := spotifyClient.SearchTracks(
		c.Request.Context(),
		song,
	)
	println("TRACKS FOUND:", len(tracks))

	if err != nil {
		c.String(
			http.StatusOK,
			"ERROR: %s",
			err.Error(),
		)
		return
	}

	data := gin.H{
		"title": "Search Results",
	}

	if len(tracks) > 0 {

		image := ""

		if len(tracks[0].Album.Images) > 0 {
			image = tracks[0].Album.Images[0].URL
		}

		artist := ""

		if len(tracks[0].Artists) > 0 {
			artist = tracks[0].Artists[0].Name
		}

		data["Query1Name"] = tracks[0].Name
		data["Query1Artist"] = artist
		data["Query1Image"] = image
		data["Query1URL"] = tracks[0].ID.String()
	}

	if len(tracks) > 1 {

		image := ""

		if len(tracks[1].Album.Images) > 0 {
			image = tracks[1].Album.Images[0].URL
		}

		artist := ""

		if len(tracks[1].Artists) > 0 {
			artist = tracks[1].Artists[0].Name
		}

		data["Query2Name"] = tracks[1].Name
		data["Query2Artist"] = artist
		data["Query2Image"] = image
		data["Query2URL"] = tracks[1].ID.String()
	}

	if len(tracks) > 2 {

		image := ""

		if len(tracks[2].Album.Images) > 0 {
			image = tracks[2].Album.Images[0].URL
		}

		artist := ""

		if len(tracks[2].Artists) > 0 {
			artist = tracks[2].Artists[0].Name
		}

		data["Query3Name"] = tracks[2].Name
		data["Query3Artist"] = artist
		data["Query3Image"] = image
		data["Query3URL"] = tracks[2].ID.String()
	}

	c.HTML(
		http.StatusOK,
		"search.html",
		data,
	)
}
