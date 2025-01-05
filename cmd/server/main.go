package main

import (
	"log"
	"os"

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

	// Inicia servidor (usa el puerto de .env)
	r := gin.Default()
	r.Run(":" + port)
}
