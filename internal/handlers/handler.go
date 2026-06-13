package handlers

import (
	"github.com/zmb3/spotify/v2"

	"spotifysrmechanism/internal/config"
	"spotifysrmechanism/internal/lyrics"
	spotifyapi "spotifysrmechanism/internal/spotify"
)

type Handler struct {
	cfg        *config.Config
	spotify    *spotifyapi.Client
	lyrics     *lyrics.Client
	oauth      *spotifyapi.OAuth
	tokenStore *spotifyapi.TokenStore
}

func New(cfg *config.Config) *Handler {

	tokenStore := spotifyapi.NewTokenStore()

	oauthClient := spotifyapi.NewOAuth(
		cfg.SpotifyClientID,
		cfg.SpotifyClientSecret,
		cfg.SpotifyRedirectURI,
	)

	return &Handler{
		cfg: cfg,

		spotify: spotifyapi.New(
			spotify.New(nil),
		),

		lyrics: lyrics.New(),

		oauth: oauthClient,

		tokenStore: tokenStore,
	}
}

func (h *Handler) spotifyClient() *spotifyapi.Client {

	token := h.tokenStore.Get()

	if token == nil {
		return nil
	}

	return spotifyapi.NewAuthorized(
		token,
		h.cfg.SpotifyClientID,
		h.cfg.SpotifyClientSecret,
	)
}
