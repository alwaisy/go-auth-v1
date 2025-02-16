package auth

import (
	"gorm.io/gorm"
	"net/http"
)

func Init(db *gorm.DB, router *http.ServeMux, basePath string) {
	repo := NewAuthRepository(db)
	service := NewAuthService(repo)
	handler := NewAuthHandler(service)

	SetupAuthRoutes(router, basePath, handler)
}
