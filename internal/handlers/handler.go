package handlers

import (
	"spotifysrmechanism/internal/config"
	"spotifysrmechanism/internal/lyrics"
	spotifyapi "spotifysrmechanism/internal/spotify"

	"github.com/zmb3/spotify/v2"
)

type Handler struct {
	cfg     *config.Config
	spotify *spotifyapi.Client
	lyrics  *lyrics.Client
	oauth   *spotifyapi.OAuth
}

func New(cfg *config.Config) *Handler {

	return &Handler{
		cfg:     cfg,
		spotify: spotifyapi.New(spotify.New(nil)),
		lyrics:  lyrics.New(),
	}
}

//jebac ten cały projekt
