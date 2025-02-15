package server

import (
	m "go-auth-v1/internal/middlewares"
	"net/http"
)

// setupRoutes applies API prefix and middlewares
func setupRoutes(mux *http.ServeMux) http.Handler {
	// Use a sub-router for API v1
	apiPrefix := "/api/v1"

	// Register prefixed routes
	mux.HandleFunc(apiPrefix+"/", HandleRoot)
	mux.HandleFunc(apiPrefix+"/health", HandleHealthCheck)

	// Apply middlewares
	handler := m.RecoveryMiddleware(m.LoggerMiddleware(mux))

	return handler
}
