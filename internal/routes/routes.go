package routes

import (
	"phishing-platform-backend/internal/api"
	"phishing-platform-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configura todas las rutas de la aplicación
func SetupRoutes(r *gin.Engine) {

	// Rutas públicas
	r.POST("/auth/register", api.Register)
	r.POST("/auth/login", api.Login)

	// Grupo de rutas protegidas
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())

	// Rutas protegidas
	protected.GET("/gophish/campaigns", api.ListCampaigns)
}
