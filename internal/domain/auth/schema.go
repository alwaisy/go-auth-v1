package auth

import (
	"github.com/google/uuid"
	"time"
)

type UserStoreSchema struct {
	FirstName string `json:"FirstName" validate:"max=25,required"`
	LastName  string `json:"LastName" validate:"max=25,required"`
	Username  string `json:"Username" validate:"max=8,required"`
	Email     string `json:"Email" validate:"max=40,email,required"`
	Password  string `json:"Password" validate:"min=6,required"`
}

type UserLoginSchema struct {
	Email    string `json:"email" validate:"max=40,email,required"`
	Password string `json:"password" validate:"min=6,required"`
}

type UserSchema struct {
	ID         uuid.UUID `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	FirstName  string    `json:"f_name"`
	LastName   string    `json:"l_name"`
	IsVerified bool      `json:"is_verified"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type EmailSchema struct {
	Addresses []string `json:"addresses" validate:"required,dive,email"`
}

type PasswordResetRequestSchema struct {
	Email string `json:"email" validate:"email,required"`
}

type PasswordResetConfirmSchema struct {
	NewPassword        string `json:"new_password" validate:"min=6,required"`
	ConfirmNewPassword string `json:"confirm_new_password" validate:"min=6,required,eqfield=NewPassword"`
}
