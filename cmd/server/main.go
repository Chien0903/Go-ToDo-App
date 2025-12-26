package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"

	"github.com/Chien0903/Go-ToDo-App/internal/config"
	"github.com/Chien0903/Go-ToDo-App/internal/database"
	"github.com/Chien0903/Go-ToDo-App/internal/handlers/rest"
	"github.com/Chien0903/Go-ToDo-App/internal/middleware"
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

	// Health check
	healthHandler := rest.NewHealthHandler()
	r.Get("/health", healthHandler.Health)

	// Public routes (không cần JWT)
	userHandler := rest.NewUserHandler(cfg)
	r.Post("/api/register", userHandler.Register)
	r.Post("/api/login", userHandler.Login)

	// Protected routes (cần JWT)
	r.Group(func(r chi.Router) {
		r.Use(middleware.JWTAuthMiddleware(cfg))
		// Thêm các route cần bảo vệ ở đây
		// Ví dụ: r.Get("/api/todos", todoHandler.GetTodos)
	})

	log.Printf("Server running on %s (%s)", cfg.Port, cfg.Environment)
	if err := http.ListenAndServe(cfg.Port, r); err != nil {
		log.Fatal(err)
	}
}
