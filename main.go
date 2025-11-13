package main

import (
	"net/http"
	"log"
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/Chien0903/Go-ToDo-App/config"
	"github.com/Chien0903/Go-ToDo-App/handlers"
)

func main() {
	cfg := config.Load()

	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)

	healthHandler := handlers.NewHealthHandler()

	r.Get("/health", healthHandler.Health)

	log.Printf("Server running on %s (%s)", cfg.Port, cfg.Environment)
	if err := http.ListenAndServe(cfg.Port, r); err != nil {
		log.Fatal(err)
	}
}