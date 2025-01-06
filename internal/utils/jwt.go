package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Expira en 24 horas
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := []byte(os.Getenv("JWT_SECRET"))
	return token.SignedString(secretKey)
}
