package handlers

import (
	"fmt"
	"net/http"
)

// HealthHandler checks if the API is running.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "API is running")
}

// HelloHandler returns a greeting message.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello from Lab 2!")
}
