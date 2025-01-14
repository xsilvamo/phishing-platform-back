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

// CreateLandingPage maneja la solicitud para crear una nueva página de aterrizaje
func CreateLandingPage(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewLandingPageService(client, apiKey, baseURL)

	// Leer los datos del cuerpo de la solicitud
	var data map[string]interface{}
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Llamar al servicio para crear la página de aterrizaje
	page, err := service.CreateLandingPage(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"page": page})
}

// UpdateLandingPage maneja la solicitud para modificar una página de aterrizaje existente
func UpdateLandingPage(c *gin.Context) {
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

	// Leer los datos del cuerpo de la solicitud
	var data map[string]interface{}
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Llamar al servicio para modificar la página de aterrizaje
	page, err := service.UpdateLandingPage(id, data)
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

// DeleteLandingPage maneja la solicitud para eliminar una página de aterrizaje existente
func DeleteLandingPage(c *gin.Context) {
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

	// Llamar al servicio para eliminar la página de aterrizaje
	if err := service.DeleteLandingPage(id); err != nil {
		// Manejar error específico de página no encontrada
		if err.Error() == "página no encontrada" {
			c.JSON(http.StatusNotFound, gin.H{"error": "La página de aterrizaje no existe"})
			return
		}

		// Otros errores
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Página de aterrizaje eliminada exitosamente"})
}

// ImportSite maneja la solicitud para importar el HTML de una URL como base para una página de aterrizaje
func ImportSite(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewLandingPageService(client, apiKey, baseURL)

	// Leer los datos del cuerpo de la solicitud
	var body struct {
		URL              string `json:"url" binding:"required"`
		IncludeResources bool   `json:"include_resources"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Llamar al servicio para importar el sitio
	result, err := service.ImportSite(body.URL, body.IncludeResources)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"html": result["html"]})
}
