package api

import (
	"net/http"
	"os"

	"phishing-platform-backend/internal/gophish"

	"github.com/gin-gonic/gin"
)

// ResetAPIKey maneja la solicitud para resetear la clave API de GoPhish
func ResetAPIKey(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewSettingsService(client, apiKey, baseURL)

	newAPIKey, err := service.ResetAPIKey()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Clave API reseteada exitosamente",
		"api_key": newAPIKey,
	})
}
