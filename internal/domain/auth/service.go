package auth

import (
	"go-auth-v1/pkg/random"
)

type Service struct {
	repo *Repository
}

func NewAuthService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(input UserStoreSchema) (User, error) {
	id := random.NewRandomID()

	user := User{
		ID:           id,
		Username:     input.Username,
		Email:        input.Email,
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		PasswordHash: "hash",
		IsVerified:   false,
		CreatedAt:    random.NewRandomTime(),
		UpdatedAt:    random.NewRandomTime(),
	}

	// Hash password

	// Create user record

	return user, nil
}
