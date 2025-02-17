package auth

import (
	"errors"
	"fmt"
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
	// Check for existing user
	for field, value := range map[string]string{"email": input.Email, "username": input.Username} {
		if exists, err := s.repo.CheckUser(field, value); err != nil {
			return nil, err
		} else if exists {
			return nil, errors.New(fmt.Sprintf("%s already taken", field)) // Ensure this gives precise error details
		}
	}

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
	err = s.repo.Store(user)

	return user, nil
}
