package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"

	"github.com/Chien0903/Go-ToDo-App/internal/config"
	"github.com/Chien0903/Go-ToDo-App/internal/database"
	"github.com/Chien0903/Go-ToDo-App/internal/handlers/rest"
)

func main() {
	cfg := config.Load()

	// Khởi tạo kết nối DB MySQL (GORM) – dùng database.Connect với AppConfig
	if err := database.Connect(cfg); err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}
	// Sau này có thể dùng database.DB để truyền vào repository/services

	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)

	healthHandler := rest.NewHealthHandler()
	r.Get("/health", healthHandler.Health)

	log.Printf("Server running on %s (%s)", cfg.Port, cfg.Environment)
	if err := http.ListenAndServe(cfg.Port, r); err != nil {
		log.Fatal(err)
	}
}
