package api

import (
	"net/http"
	"phishing-platform-backend/internal/models"
	"phishing-platform-backend/internal/repository"
	"phishing-platform-backend/internal/utils"

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

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	// Validar el body de la solicitud
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Buscar el usuario por email
	repo := repository.UserRepository{DB: repository.DB}
	user, err := repo.FindByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario o contraseña incorrectos"})
		return
	}

	// Comparar contraseñas
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario o contraseña incorrectos"})
		return
	}

	// Generar el token JWT
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar el token"})
		return
	}

	// Respuesta con el token
	c.JSON(http.StatusOK, gin.H{
		"message": "Inicio de sesión exitoso",
		"token":   token,
	})
}
