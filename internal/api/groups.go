package api

import (
	"net/http"
	"os"

	"phishing-platform-backend/internal/gophish"

	"github.com/gin-gonic/gin"
)

// GetGroups maneja la solicitud para obtener todos los grupos
func GetGroups(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewGroupService(client, apiKey, baseURL)

	groups, err := service.GetGroups()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"groups": groups})
}
