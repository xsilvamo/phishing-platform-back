package api

import (
	"io"
	"net/http"
	"os"
	"strconv"

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

// GetGroupByID maneja la solicitud para obtener un grupo específico por su ID
func GetGroupByID(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewGroupService(client, apiKey, baseURL)

	// Obtener el ID desde los parámetros de la URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	group, err := service.GetGroupByID(id)
	if err != nil {
		// Manejar error específico de grupo no encontrado
		if err.Error() == "grupo no encontrado" {
			c.JSON(http.StatusNotFound, gin.H{"error": "El grupo no existe"})
			return
		}

		// Otros errores
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"group": group})
}

// GetGroupsSummary maneja la solicitud para obtener un resumen de todos los grupos
func GetGroupsSummary(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewGroupService(client, apiKey, baseURL)

	summaries, err := service.GetGroupsSummary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, summaries)
}

// GetGroupSummaryByID maneja la solicitud para obtener un resumen de un grupo específico por su ID
func GetGroupSummaryByID(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewGroupService(client, apiKey, baseURL)

	// Obtener el ID desde los parámetros de la URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	summary, err := service.GetGroupSummaryByID(id)
	if err != nil {
		// Manejar error específico de grupo no encontrado
		if err.Error() == "grupo no encontrado" {
			c.JSON(http.StatusNotFound, gin.H{"error": "El grupo no existe"})
			return
		}

		// Otros errores
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"group_summary": summary})
}

// CreateGroup maneja la solicitud para crear un nuevo grupo
func CreateGroup(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewGroupService(client, apiKey, baseURL)

	// Leer los datos del cuerpo de la solicitud
	var data map[string]interface{}
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Llamar al servicio para crear el grupo
	group, err := service.CreateGroup(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"group": group})
}

// UpdateGroup maneja la solicitud para modificar un grupo existente
func UpdateGroup(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewGroupService(client, apiKey, baseURL)

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

	// Llamar al servicio para modificar el grupo
	updatedGroup, err := service.UpdateGroup(id, data)
	if err != nil {
		// Manejar error específico de grupo no encontrado
		if err.Error() == "grupo no encontrado" {
			c.JSON(http.StatusNotFound, gin.H{"error": "El grupo no existe"})
			return
		}

		// Otros errores
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"group": updatedGroup})
}

// DeleteGroup maneja la solicitud para eliminar un grupo existente
func DeleteGroup(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewGroupService(client, apiKey, baseURL)

	// Obtener el ID desde los parámetros de la URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Llamar al servicio para eliminar el grupo
	if err := service.DeleteGroup(id); err != nil {
		// Manejar error específico de grupo no encontrado
		if err.Error() == "grupo no encontrado" {
			c.JSON(http.StatusNotFound, gin.H{"error": "El grupo no existe"})
			return
		}

		// Otros errores
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Grupo eliminado exitosamente"})
}

// ImportGroup maneja la solicitud para procesar un archivo CSV y devolver los objetivos del grupo
func ImportGroup(c *gin.Context) {
	client := &http.Client{}
	apiKey := os.Getenv("GOPHISH_API_KEY")
	baseURL := os.Getenv("GOPHISH_API_URL")

	service := gophish.NewGroupService(client, apiKey, baseURL)

	// Obtener el archivo de la solicitud
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Archivo no proporcionado"})
		return
	}

	// Guardar temporalmente el archivo para su procesamiento
	tempFile, err := os.CreateTemp("", "upload-*.csv")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creando archivo temporal"})
		return
	}
	defer os.Remove(tempFile.Name())

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error abriendo el archivo"})
		return
	}
	defer src.Close()

	_, err = io.Copy(tempFile, src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error copiando el archivo"})
		return
	}

	// Llamar al servicio para procesar el archivo CSV
	targets, err := service.ImportGroup(tempFile.Name())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"targets": targets})
}
