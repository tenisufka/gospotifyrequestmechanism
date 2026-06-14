package handlers

import (
	"spotifysrmechanism/internal/config"
	"spotifysrmechanism/internal/lyrics"
	"spotifysrmechanism/internal/spotify"
)

type Handler struct {
	cfg        *config.Config
	lyrics     *lyrics.Client
	oauth      *spotify.OAuth
	tokenStore *spotify.TokenStore
}

func New(cfg *config.Config) *Handler {
	tokenStore := spotify.NewTokenStore()

	oauthClient := spotify.NewOAuth(
		cfg.SpotifyClientID,
		cfg.SpotifyClientSecret,
		cfg.SpotifyRedirectURI,
	)

	return &Handler{
		cfg:        cfg,
		lyrics:     lyrics.New(),
		oauth:      oauthClient,
		tokenStore: tokenStore,
	}
}

func (h *Handler) spotifyClient() *spotify.Client {
	token := h.tokenStore.Get()

	if token == nil {
		return nil
	}

	return spotify.NewAuthorized(
		token,
		h.cfg.SpotifyClientID,
		h.cfg.SpotifyClientSecret,
		h.tokenStore,
	)
}
