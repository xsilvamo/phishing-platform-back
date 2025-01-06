package api

import (
	"net/http"
	"phishing-platform-backend/internal/models"
	"phishing-platform-backend/internal/repository"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Encriptar contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encriptando la contraseña"})
		return
	}

	// Crear usuario
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	repo := repository.UserRepository{DB: repository.DB}
	if err := repo.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creando el usuario"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuario registrado existosamente"})
}
