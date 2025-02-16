package auth

import "net/http"

func SetupAuthRoutes(router *http.ServeMux, basePath string, h *Handler) {
	router.HandleFunc("POST "+basePath+"/signup", h.HandleSignup)
}
