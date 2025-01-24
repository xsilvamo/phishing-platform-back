package api

import (
	"net/http"
	"os"
	"strconv"

	"phishing-platform-backend/internal/gophish"

	"github.com/gin-gonic/gin"
)

// GetCurrentUser maneja la solicitud para obtener información del usuario autenticado
func GetCurrentUser(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewUserService(client, apiKey, baseURL)

	users, err := service.GetCurrentUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No se encontró información del usuario"})
		return
	}

	// Retornar el primer usuario (autenticado)
	c.JSON(http.StatusOK, gin.H{"user": users[0]})
}

// CreateUser maneja la solicitud para crear un usuario en GoPhish
func CreateUser(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewUserService(client, apiKey, baseURL)

	var data map[string]interface{}
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Crear usuario en GoPhish
	user, err := service.CreateUser(data)
	if err != nil {
		// Manejar errores específicos
		if err.Error() == "error de la API: Username already taken" {
			c.JSON(http.StatusConflict, gin.H{"error": "El nombre de usuario ya está en uso"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

// UpdateUser maneja la solicitud para actualizar un usuario en GoPhish
func UpdateUser(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewUserService(client, apiKey, baseURL)

	// Obtener el ID desde los parámetros de la URL
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

	// Actualizar usuario en GoPhish
	user, err := service.UpdateUser(id, data)
	if err != nil {
		if err.Error() == "error de la API: Username already taken" {
			c.JSON(http.StatusConflict, gin.H{"error": "El nombre de usuario ya está en uso"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
