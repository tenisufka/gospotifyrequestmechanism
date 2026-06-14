package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"spotifysrmechanism/internal/config"
	"spotifysrmechanism/internal/filters"
	"spotifysrmechanism/internal/handlers"
)

func main() {
	// 1. Ładowanie konfiguracji środowiskowej
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("configuration error: %v", err)
	}

	// 2. Ładowanie pliku z zakazanymi słowami
	if err := filters.LoadBadWords("internal/filters/badwords.txt"); err != nil {
		log.Fatalf("failed to load bad words: %v", err)
	}

	// 3. Inicjalizacja routera Gin
	router := gin.Default()

	// -----------------------------
	// STATIC FILES
	// -----------------------------
	router.Static("/static", "./static")

	// -----------------------------
	// HTML TEMPLATES
	// -----------------------------
	router.LoadHTMLGlob("templates/*")

	// 4. Tworzenie instancji handlera z przekazaną konfiguracją
	h := handlers.New(cfg)

	// 5. Rejestracja endpointów aplikacji
	router.GET("/", h.Index)
	router.GET("/login", h.Login)
	router.GET("/callback", h.Callback)

	router.GET("/search", h.Search)
	router.POST("/search", h.SearchResults)

	router.POST("/result", h.Result)

	// Endpoint sprawdzania stanu aplikacji
	router.GET("/logout", h.Logout)
	router.GET("/health", h.Health)

	// 6. Uruchomienie serwera HTTP
	log.Printf("Spotify Request Mechanism started on %s", cfg.ServerPort)
	if err := router.Run(cfg.ServerPort); err != nil {
		log.Fatalf("server startup failed: %v", err)
	}
}
