package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SpotifyClientID     string
	SpotifyClientSecret string
	SpotifyRedirectURI  string
	ServerPort          string
	SessionSecret       string // Nowe pole dla klucza sesji
}

func Load() (*Config, error) {
	// Plik .env powinien znajdować się w głównym katalogu projektu: <root projektu>/.env
	// Ignorujemy błąd, jeśli plik .env nie istnieje (np. w środowisku produkcyjnym zmienne są wstrzykiwane bezpośrednio)
	_ = godotenv.Load()

	cfg := &Config{
		SpotifyClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		SpotifyClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		SpotifyRedirectURI:  os.Getenv("SPOTIFY_REDIRECT_URI"),
		ServerPort:          os.Getenv("SERVER_PORT"),
		SessionSecret:       os.Getenv("SESSION_SECRET"), // Pobieranie klucza sesji
	}

	// Walidacja wymaganych zmiennych dla Spotify API
	if cfg.SpotifyClientID == "" {
		return nil, fmt.Errorf("missing SPOTIFY_CLIENT_ID")
	}

	if cfg.SpotifyClientSecret == "" {
		return nil, fmt.Errorf("missing SPOTIFY_CLIENT_SECRET")
	}

	// Walidacja nowego pola SessionSecret (dobra praktyka bezpieczeństwa)
	if cfg.SessionSecret == "" {
		return nil, fmt.Errorf("missing SESSION_SECRET")
	}

	// Ustawienie domyślnego portu, jeśli nie został zdefiniowany
	if cfg.ServerPort == "" {
		cfg.ServerPort = ":8080"
	}

	return cfg, nil
}
