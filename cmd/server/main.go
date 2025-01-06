package main

import (
	"log"
	"os"

	"phishing-platform-backend/internal/api"
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

	r.GET("/users", func(c *gin.Context) {
		repo := repository.UserRepository{DB: repository.DB}
		users, err := repo.GetAll()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, users)
	})

	r.POST("/auth/register", api.Register)

	r.Run(":" + port)
}
