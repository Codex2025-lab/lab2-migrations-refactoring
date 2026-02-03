package main

import (
	"log"
	"net/http"

	"github.com/CODEX2025/lab2-migrations-refactoring/internal/routes"
)

func main() {
	handler := routes.SetupRoutes()

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
