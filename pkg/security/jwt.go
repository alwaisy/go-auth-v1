package security

import (
	"fmt"
	"github.com/google/uuid"
	"go-auth-v1/internal/config"
	"time"
)

import (
	"github.com/golang-jwt/jwt"
)

var (
	JWTAlgorithm      = jwt.SigningMethodHS256
	AccessTokenExpiry = 15 * time.Minute
)

// Claims struct for JWT payload, embedded with StandardClaims
type Claims struct {
	UserData map[string]interface{} `json:"user"`
	Refresh  bool                   `json:"refresh"`
	Type     string                 `json:"type,omitempty"`
	jwt.StandardClaims
}

// GenerateJWT generates a JWT for user authentication
func GenerateJWT(userData map[string]interface{}, expiry time.Duration, refresh bool) (string, error) {
	cfg := config.LoadConfig(".")

	if expiry == 0 {
		expiry = AccessTokenExpiry // Default to 15 minutes if no expiry is provided
	}

	claims := Claims{
		UserData: userData,
		Refresh:  refresh,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiry).Unix(),
			Id:        uuid.New().String(),
		},
	}

	token := jwt.NewWithClaims(JWTAlgorithm, claims)

	signedToken, err := token.SignedString([]byte(cfg.Auth.JwtSecret))
	if err != nil {
		return "", fmt.Errorf("could not sign token: %w", err)
	}

	return signedToken, nil
}
