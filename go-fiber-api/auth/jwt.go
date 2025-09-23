package auth

import (
	"time"
	"go-fiber-api/config"

	"github.com/golang-jwt/jwt/v5"
)

var cfg *config.Config

func InitAuth(c *config.Config) {
	cfg = c
}

// GenerateToken creates a JWT token for a user
func GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(), // expires in 72h
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecret))
}

// SecretKey returns the secret for Fiber middleware
func SecretKey() []byte {
	return []byte(cfg.JWTSecret)
}
