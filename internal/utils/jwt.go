package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// GenerateJWT genera un token JWT con el userID
func GenerateJWT(userID uint) (string, error) {
	// Definir los claims del token
	claims := jwt.MapClaims{
		"id":  float64(userID),                       // Convertir a float64 para compatibilidad JSON
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Expira en 24 horas
	}

	// Crear el token con los claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firmar el token con la clave secreta
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWT valida un token JWT y retorna sus claims
func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	// Obtener la clave secreta
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	// Parsear el token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validar el método de firma
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de firma inválido")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Extraer y validar los claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token inválido")
}

// ExtractUserID extrae el ID de usuario de los claims
func ExtractUserID(claims jwt.MapClaims) (uint, error) {
	// Obtener el ID del usuario de los claims
	userIDFloat, ok := claims["id"].(float64)
	if !ok {
		return 0, errors.New("ID de usuario no encontrado en el token")
	}

	return uint(userIDFloat), nil
}
