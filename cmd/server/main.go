package main

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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
	// Plik powinieneś umieścić w: <root projektu>/internal/filters/badwords.txt
	if err := filters.LoadBadWords("internal/filters/badwords.txt"); err != nil {
		log.Fatalf("failed to load bad words: %v", err)
	}

	// 3. Inicjalizacja routera Gin
	router := gin.Default()

	// Konfiguracja magazynu sesji opartego na ciasteczkach (cookies)
	store := cookie.NewStore(
		[]byte(cfg.SessionSecret),
	)

	// Rejestracja middleware do obsługi sesji w routerze
	router.Use(
		sessions.Sessions(
			"spotify-session",
			store,
		),
	)

	// -----------------------------
	// STATIC FILES
	// -----------------------------
	// Serwowanie plików statycznych z katalogu <root projektu>/static
	router.Static("/static", "./static")

	// -----------------------------
	// HTML TEMPLATES
	// -----------------------------
	// Ładowanie szablonów HTML z katalogu <root projektu>/templates
	router.LoadHTMLGlob("templates/*")

	// 4. Tworzenie instancji handlera z przekazaną konfiguracją
	h := handlers.New(cfg)

	// 5. Rejestracja endpointów aplikacji
	router.GET("/", h.Index)
	router.GET("/search", h.Search)
	router.POST("/result", h.Result)

	// Nowe endpointy do autoryzacji Spotify OAuth2
	router.GET("/login", h.Login)
	router.GET("/callback", h.Callback)
	router.GET("/logout", h.Logout) // Dodany endpoint wylogowania

	// 6. Uruchomienie serwera HTTP
	log.Printf("Spotify Request Mechanism started on %s", cfg.ServerPort)
	if err := router.Run(cfg.ServerPort); err != nil {
		log.Fatalf("server startup failed: %v", err)
	}
}
