package main

import (
	"go-auth-v1/internal/server"
)

func main() {
	app := server.NewApp()

	app.Start()
}
