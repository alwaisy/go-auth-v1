package auth

import (
	"errors"
	"fmt"
	customErr "go-auth-v1/pkg/errors"
	"go-auth-v1/pkg/random"
	"go-auth-v1/pkg/security"
	"time"
)

type Service struct {
	repo *Repository
}

func NewAuthService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(input UserStoreSchema) (*User, error) {
	// Check for existing user
	for field, value := range map[string]string{"email": input.Email, "username": input.Username} {
		if exists, err := s.repo.CheckUser(field, value); err != nil {
			return nil, err
		} else if exists {
			// Ensure this gives precise error details
			return nil, errors.New(fmt.Sprintf("%s already used", field))
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
	err = s.repo.CreateUser(user)

	return user, nil
}

func (s *Service) Login(input UserLoginSchema) (string, error) {
	// check email existence
	user, err := s.repo.ShowUserByEmail(input.Email)
	if err != nil {
		return "", customErr.ErrInvalidCredentials
	}

	// compare password
	isMatched := security.VerifyPassword(user.PasswordHash, input.Password)
	if !isMatched {
		return "", customErr.ErrInvalidCredentials
	}

	userData := map[string]interface{}{
		"id":    user.ID,
		"email": input.Email,
	}

	// generate token
	token, err := security.GenerateJWT(userData, 30*time.Minute, false)

	return token, nil
}
