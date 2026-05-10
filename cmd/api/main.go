package main

import (
	"log"
	"net/http"
	"time"

	"github.com/muhammadakbarra/go-rest-api-boilerplate/internal/config"
	"github.com/muhammadakbarra/go-rest-api-boilerplate/internal/database"
	"github.com/muhammadakbarra/go-rest-api-boilerplate/internal/posts"
	_ "github.com/muhammadakbarra/go-rest-api-boilerplate/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Go REST API
// @version 1.0
// @description Simple REST API using Go, chi, and PostgreSQL.
// @host localhost:8080
// @BasePath /
func main() {
	cfg := config.Load()

	db, err := database.NewPostgresPool(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}
	defer db.Close()

	postRepo := posts.NewRepository(db)
	postHandler := posts.NewHandler(postRepo)

	r := chi.NewRouter()


	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))


	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	r.Get("/health", HealthCheck)

	postHandler.RegisterRoutes(r)

	addr := ":" + cfg.AppPort

	server := &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("server running on http://localhost%s", addr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}

// HealthCheck godoc
// @Summary Health check
// @Description Get the status of the API
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}
