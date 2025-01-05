package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Crear router Gin
	r := gin.Default()

	// Rutas básicas
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Iniciar servidor
	fmt.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(r.Run(":8080"))
}
