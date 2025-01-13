package api

import (
	"net/http"
	"os"
	"strconv"

	"phishing-platform-backend/internal/gophish"

	"github.com/gin-gonic/gin"
)

// GetLandingPages maneja la solicitud para obtener todas las páginas de aterrizaje
func GetLandingPages(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewLandingPageService(client, apiKey, baseURL)

	pages, err := service.GetLandingPages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"pages": pages})
}

// GetLandingPageByID maneja la solicitud para obtener una página de aterrizaje por su ID
func GetLandingPageByID(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewLandingPageService(client, apiKey, baseURL)

	// Obtener el ID desde los parámetros de la URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	page, err := service.GetLandingPageByID(id)
	if err != nil {
		// Manejar error específico de página no encontrada
		if err.Error() == "página no encontrada" {
			c.JSON(http.StatusNotFound, gin.H{"error": "La página de aterrizaje no existe"})
			return
		}

		// Otros errores
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"page": page})
}
