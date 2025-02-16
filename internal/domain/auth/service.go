package auth

import (
	"go-auth-v1/pkg/random"
	"go-auth-v1/pkg/security"
)

type Service struct {
	repo *Repository
}

func NewAuthService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(input UserStoreSchema) (*User, error) {
	id := random.NewRandomID()

	// Hash password
	hashedPassword, err := security.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:           id,
		Username:     input.Username,
		Email:        input.Email,
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		PasswordHash: hashedPassword,
		Role:         "user",
		IsVerified:   false,
		CreatedAt:    random.NewRandomTime(),
		UpdatedAt:    random.NewRandomTime(),
	}

	// Create user record
	err = s.repo.Index(user)

	return user, nil
}
