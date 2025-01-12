package api

import (
	"net/http"
	"os"

	"phishing-platform-backend/internal/gophish"

	"github.com/gin-gonic/gin"
)

// ListCampaigns maneja la solicitud para obtener las campañas
func ListCampaigns(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewCampaignService(client, apiKey, baseURL)

	campaigns, err := service.ListCampaigns()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"campaigns": campaigns})
}
