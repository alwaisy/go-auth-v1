package server

import (
	"fmt"
	"go-auth-v1/internal/config"
	"go-auth-v1/internal/database"
	"go-auth-v1/internal/domain/auth"
	"log"
	"net/http"
	"time"
)

// App types
type App struct {
	Server *http.Server
}

func MainApp() *App {
	// config
	cfg := config.LoadConfig(".")

	// init database
	database.ConnectDB()

	// port
	PORT := cfg.Server.Port
	if PORT == "" {
		PORT = "8888"
	}

	// init mux server
	mux := http.NewServeMux()

	// server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", PORT),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      setupRoutes(mux),
	}

	// init modules
	auth.Init(mux, "/api/v1/auth")

	return &App{
		Server: server,
	}

}

func (app *App) Start() {
	//fmt.Printf("listening on port %s", app.Server.Addr)
	fmt.Println("listening on port", app.Server.Addr)

	err := app.Server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
