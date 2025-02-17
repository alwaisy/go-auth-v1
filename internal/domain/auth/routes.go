package auth

import "net/http"

func SetupAuthRoutes(router *http.ServeMux, basePath string, h *Handler) {
	router.HandleFunc("POST "+basePath+"/signup", h.HandleSignup)
	router.HandleFunc("POST "+basePath+"/login", h.HandleLogin)

}
