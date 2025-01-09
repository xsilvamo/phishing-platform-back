package api

import (
	"net/http"
	"phishing-platform-backend/internal/gophish"

	"github.com/gin-gonic/gin"
)

func ListCampaigns(c *gin.Context) {
	service := gophish.NewGoPhishService()

	campaigns, err := service.ListCampaigns()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"campaigns": campaigns})
}
