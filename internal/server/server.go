package server

import (
	"context"
	"errors"
	"fmt"
	"go-auth-v1/internal/config"
	"go-auth-v1/internal/database"
	"go-auth-v1/internal/domain/auth"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// App struct
type App struct {
	Server *http.Server
}

// MainApp initializes the app
func MainApp() *App {
	// Load config
	cfg := config.LoadConfig(".")

	// Initialize database
	database.ConnectDB()

	// Get port from config (default: 8888)
	PORT := cfg.Server.Port
	if PORT == "" {
		PORT = "8888"
	}

	// Create a new HTTP multiplexer
	mux := http.NewServeMux()

	// Create HTTP server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", PORT),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      setupRoutes(mux),
	}

	// Initialize authentication routes
	auth.Init(mux, "/api/v1/auth")

	return &App{
		Server: server,
	}
}

// Start runs the server and manages graceful shutdown
func (app *App) Start() {
	fmt.Println("Listening on port", app.Server.Addr)

	// Run server in a separate goroutine
	go func() {
		if err := app.Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server error: %v", err)
		}
	}()

	handleGracefulShutdown(app.Server)
}

// handleGracefulShutdown listens for termination signals and shuts down the server
func handleGracefulShutdown(server *http.Server) {
	cfg := config.LoadConfig(".")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Wait for termination signal
	<-stop
	fmt.Println("\nReceived shutdown signal...")

	// Check if running in development mode
	if cfg.App.Env == "development" {
		fmt.Println("Skipping full shutdown in development mode.")
		return
	}

	fmt.Println("Shutting down gracefully...")

	// Create a timeout context for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to shut down the server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	// Close database connections
	database.CloseDB()

	fmt.Println("Server shutdown complete.")
}
