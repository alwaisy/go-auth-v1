package errors

import "errors"

var (
	ErrInvalidToken           = errors.New("invalid or expired token")
	ErrRevokedToken           = errors.New("token has been revoked")
	ErrAccessTokenRequired    = errors.New("access token is required")
	ErrRefreshTokenRequired   = errors.New("refresh token is required")
	ErrUserAlreadyExists      = errors.New("user already exists")
	ErrInvalidCredentials     = errors.New("invalid email or password")
	ErrInsufficientPermission = errors.New("insufficient permissions")
	ErrUserNotFound           = errors.New("user not found")
	ErrAccountNotVerified     = errors.New("account not verified")
	ErrInternalServerError    = errors.New("internal server error")
	ErrEmailTaken             = errors.New("email already taken")
	ErrUsernameTaken          = errors.New("username already taken")
)
