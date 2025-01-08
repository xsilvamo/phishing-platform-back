package middleware

import (
	"net/http"
	"strings"

	"phishing-platform-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Falta el encabezado Authorization"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de encabezado Authorization inválido"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido: " + err.Error()})
			c.Abort()
			return
		}

		// Extraer el userID como float64 y convertirlo a uint
		userIDFloat, ok := claims["id"].(float64)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "userID no válido en el token"})
			c.Abort()
			return
		}

		// Convertir de float64 a uint
		userID := uint(userIDFloat)

		// Almacenar en el contexto
		c.Set("userID", userID)
		c.Next()
	}
}
