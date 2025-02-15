package auth

import "net/http"

func Init(router *http.ServeMux, basePath string) {
	SetupAuthRoutes(router, basePath)
}
