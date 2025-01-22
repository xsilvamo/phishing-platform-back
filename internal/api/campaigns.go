package api

import (
	"net/http"
	"os"
	"strconv"

	"phishing-platform-backend/internal/gophish"

	"github.com/gin-gonic/gin"
)

// GetCampaigns maneja la solicitud para obtener todas las campañas
func GetCampaigns(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewCampaignService(client, apiKey, baseURL)

	campaigns, err := service.GetCampaigns()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"campaigns": campaigns})
}

// GetCampaignByID maneja la solicitud para obtener los detalles de una campaña específica
func GetCampaignByID(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewCampaignService(client, apiKey, baseURL)

	// Obtener el ID desde los parámetros de la URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	campaign, err := service.GetCampaignByID(id)
	if err != nil {
		// Manejar error específico de campaña no encontrada
		if err.Error() == "campaña no encontrada" {
			c.JSON(http.StatusNotFound, gin.H{"error": "La campaña no existe"})
			return
		}

		// Otros errores
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"campaign": campaign})
}

// CreateCampaign maneja la solicitud para crear una nueva campaña
func CreateCampaign(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewCampaignService(client, apiKey, baseURL)

	// Leer los datos del cuerpo de la solicitud
	var data map[string]interface{}
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Llamar al servicio para crear la campaña
	campaign, err := service.CreateCampaign(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"campaign": campaign})
}

// GetCampaignResults maneja la solicitud para obtener los resultados de una campaña específica
func GetCampaignResults(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewCampaignService(client, apiKey, baseURL)

	// Obtener el ID desde los parámetros de la URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	results, err := service.GetCampaignResults(id)
	if err != nil {
		// Manejar error específico de campaña no encontrada
		if err.Error() == "campaña no encontrada" {
			c.JSON(http.StatusNotFound, gin.H{"error": "La campaña no existe"})
			return
		}

		// Otros errores
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": results})
}

// GetCampaignSummary maneja la solicitud para obtener el resumen de una campaña específica
func GetCampaignSummary(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewCampaignService(client, apiKey, baseURL)

	// Obtener el ID desde los parámetros de la URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	summary, err := service.GetCampaignSummary(id)
	if err != nil {
		// Manejar error específico de campaña no encontrada
		if err.Error() == "campaña no encontrada" {
			c.JSON(http.StatusNotFound, gin.H{"error": "La campaña no existe"})
			return
		}

		// Otros errores
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"summary": summary})
}

// DeleteCampaign maneja la solicitud para eliminar una campaña específica
func DeleteCampaign(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewCampaignService(client, apiKey, baseURL)

	// Obtener el ID desde los parámetros de la URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Llamar al servicio para eliminar la campaña
	if err := service.DeleteCampaign(id); err != nil {
		// Manejar error específico de campaña no encontrada
		if err.Error() == "campaña no encontrada" {
			c.JSON(http.StatusNotFound, gin.H{"error": "La campaña no existe"})
			return
		}

		// Otros errores
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Campaña eliminada exitosamente"})
}

// CompleteCampaign maneja la solicitud para marcar una campaña como completada
func CompleteCampaign(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewCampaignService(client, apiKey, baseURL)

	// Obtener el ID desde los parámetros de la URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Llamar al servicio para marcar la campaña como completada
	if err := service.CompleteCampaign(id); err != nil {
		// Manejar errores específicos de la API de GoPhish
		if err.Error() == "la campaña no existe o no se puede completar" {
			c.JSON(http.StatusNotFound, gin.H{"error": "La campaña no existe o no puede completarse"})
			return
		}

		// Otros errores
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Campaña marcada como completada exitosamente"})
}
