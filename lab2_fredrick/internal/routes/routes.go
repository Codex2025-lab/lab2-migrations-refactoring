package routes

import (
	"net/http"

	"github.com/CODEX2025/lab2-migrations-refactoring/internal/handlers"
	"github.com/CODEX2025/lab2-migrations-refactoring/internal/middleware"
)

// SetupRoutes registers routes and applies middleware.
func SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/hello", handlers.HelloHandler)

	// Apply middleware chain
	handler := middleware.LoggingMiddleware(
		middleware.TimingMiddleware(mux),
	)

	return handler
}
