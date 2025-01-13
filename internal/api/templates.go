package api

import (
	"net/http"
	"os"
	"strconv"

	"phishing-platform-backend/internal/gophish"

	"github.com/gin-gonic/gin"
)

// GetTemplates maneja la solicitud para obtener todos los templates
func GetTemplates(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewTemplateService(client, apiKey, baseURL)

	templates, err := service.GetTemplates()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"templates": templates})
}

// GetTemplateByID maneja la solicitud para obtener un template por su ID
func GetTemplateByID(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewTemplateService(client, apiKey, baseURL)

	// Obtener el ID desde los parámetros de la URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	template, err := service.GetTemplateByID(id)
	if err != nil {
		// Manejar error específico de template no encontrado
		if err.Error() == "template no encontrado" {
			c.JSON(http.StatusNotFound, gin.H{"error": "El template no existe"})
			return
		}

		// Otros errores
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"template": template})
}

// CreateTemplate maneja la solicitud para crear un nuevo template
func CreateTemplate(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewTemplateService(client, apiKey, baseURL)

	// Leer los datos del cuerpo de la solicitud
	var data map[string]interface{}
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Llamar al servicio para crear el template
	template, err := service.CreateTemplate(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"template": template})
}

// UpdateTemplate maneja la solicitud para modificar un template existente
func UpdateTemplate(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewTemplateService(client, apiKey, baseURL)

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

	// Llamar al servicio para modificar el template
	template, err := service.UpdateTemplate(id, data)
	if err != nil {
		// Manejar error específico de template no encontrado
		if err.Error() == "template no encontrado" {
			c.JSON(http.StatusNotFound, gin.H{"error": "El template no existe"})
			return
		}

		// Otros errores
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"template": template})
}

// DeleteTemplate maneja la solicitud para eliminar un template
func DeleteTemplate(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewTemplateService(client, apiKey, baseURL)

	// Obtener el ID desde los parámetros de la URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Llamar al servicio para eliminar el template
	if err := service.DeleteTemplate(id); err != nil {
		// Manejar error específico de template no encontrado
		if err.Error() == "template no encontrado" {
			c.JSON(http.StatusNotFound, gin.H{"error": "El template no existe"})
			return
		}

		// Otros errores
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Template eliminado exitosamente"})
}

// ImportEmail maneja la solicitud para analizar un correo electrónico
func ImportEmail(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewTemplateService(client, apiKey, baseURL)

	// Leer los datos del cuerpo de la solicitud
	var body struct {
		Content      string `json:"content" binding:"required"`
		ConvertLinks bool   `json:"convert_links"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Llamar al servicio para analizar el correo
	result, err := service.ImportEmail(body.Content, body.ConvertLinks)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}
