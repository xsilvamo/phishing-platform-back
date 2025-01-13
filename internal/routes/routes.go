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

	// Settings
	protected.POST("/settings/reset_api_key", api.ResetAPIKey)

	// Perfiles de envío
	protected.GET("/profiles", api.GetProfiles)
	protected.GET("/profiles/:id", api.GetProfileByID)
	protected.POST("/profiles", api.CreateProfile)
	protected.PUT("/profiles/:id", api.UpdateProfile)
	protected.DELETE("/profiles/:id", api.DeleteProfile)

	// Templates
	protected.GET("/templates", api.GetTemplates)
	protected.GET("/templates/:id", api.GetTemplateByID)
	protected.POST("/templates", api.CreateTemplate)
	protected.PUT("/templates/:id", api.UpdateTemplate)
	protected.DELETE("/templates/:id", api.DeleteTemplate)
	protected.POST("/templates/import/email", api.ImportEmail)

	// Landing Pages

	// Rutas de campañas
	protected.GET("/gophish/campaigns", api.ListCampaigns)
}
