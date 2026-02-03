package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs request method, path, and timestamp.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s %s",
			time.Now().Format(time.RFC3339),
			r.Method,
			r.URL.Path,
		)
		next.ServeHTTP(w, r)
	})
}

// TimingMiddleware logs how long a request takes.
func TimingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("Request to %s took %v", r.URL.Path, time.Since(start))
	})
}
