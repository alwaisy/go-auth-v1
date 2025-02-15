package auth

import (
	"go-auth-v1/pkg/random"
)

func CreateUser(input UserStoreSchema) (User, error) {
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
