package auth

import "net/http"

func SetupAuthRoutes(router *http.ServeMux, basePath string) {
	router.HandleFunc("POST "+basePath+"/signup", HandleSignup)
}
