package api

import (
	"net/http"
	"os"
	"strconv"

	"phishing-platform-backend/internal/gophish"

	"github.com/gin-gonic/gin"
)

// GetProfiles maneja la solicitud para obtener todos los perfiles de envío
func GetProfiles(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewProfileService(client, apiKey, baseURL)

	profiles, err := service.GetProfiles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"profiles": profiles})
}

// GetProfileByID maneja la solicitud para obtener un perfil de envío por su ID
func GetProfileByID(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewProfileService(client, apiKey, baseURL)

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	profile, err := service.GetProfileByID(id)
	if err != nil {
		// Revisar si es un error de perfil no encontrado
		if err.Error() == "error en la respuesta: SMTP not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Perfil de envío no encontrado"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": profile})
}

// CreateProfile maneja la solicitud para crear un nuevo perfil de envío
func CreateProfile(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewProfileService(client, apiKey, baseURL)

	// Leer los datos de la solicitud
	var data map[string]interface{}
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Llamar al servicio
	profile, err := service.CreateProfile(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Responder con el perfil creado
	c.JSON(http.StatusCreated, gin.H{"profile": profile})
}

// UpdateProfile maneja la solicitud para modificar un perfil de envío
func UpdateProfile(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewProfileService(client, apiKey, baseURL)

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var data map[string]interface{}
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	profile, err := service.UpdateProfile(id, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": profile})
}

// DeleteProfile maneja la solicitud para eliminar un perfil de envío
func DeleteProfile(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewProfileService(client, apiKey, baseURL)

	// Obtener el ID desde los parámetros de la URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Llamar al servicio para eliminar el perfil
	err = service.DeleteProfile(id)
	if err != nil {
		// Manejar error específico de perfil no encontrado
		if err.Error() == "perfil no encontrado" {
			c.JSON(http.StatusNotFound, gin.H{"error": "El perfil de envío no existe"})
			return
		}

		// Otros errores
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Perfil eliminado exitosamente"})
}
