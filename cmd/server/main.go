package main

import (
	"log"
	"net/http"
	"os"

	"phishing-platform-backend/internal/api"
	"phishing-platform-backend/internal/middleware"
	"phishing-platform-backend/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT no configurado en .env")
	}

	// Inicializamos la base de datos
	repository.InitDB()

	// Inicia servidor (usa el puerto de .env)
	r := gin.Default()

	// Rutas públicas
	r.POST("/auth/register", api.Register)
	r.POST("/auth/login", api.Login)

	// Grupo de rutas protegidas
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/protected", func(c *gin.Context) {
		userID := c.MustGet("userID").(uint)
		c.JSON(http.StatusOK, gin.H{
			"message": "Acceso autorizado",
			"userID":  userID,
		})
	})

	r.Run(":" + port)
}
