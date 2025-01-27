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
	protected.GET("/landing-pages", api.GetLandingPages)
	protected.GET("/landing-pages/:id", api.GetLandingPageByID)
	protected.POST("/landing-pages", api.CreateLandingPage)
	protected.PUT("/landing-pages/:id", api.UpdateLandingPage)
	protected.DELETE("/landing-pages/:id", api.DeleteLandingPage)
	protected.POST("/landing-pages/import/site", api.ImportSite)

	// Users & Groups
	protected.GET("/groups", api.GetGroups)
	protected.GET("/groups/:id", api.GetGroupByID)
	protected.GET("/groups/summary", api.GetGroupsSummary)
	protected.GET("/groups/:id/summary", api.GetGroupSummaryByID)
	protected.POST("/groups", api.CreateGroup)
	protected.PUT("/groups/:id", api.UpdateGroup)
	protected.DELETE("/groups/:id", api.DeleteGroup)
	protected.POST("/groups/import", api.ImportGroup)

	// Rutas de campañas
	protected.GET("/campaigns", api.GetCampaigns)
	protected.GET("/campaigns/:id", api.GetCampaignByID)
	protected.POST("/campaigns", api.CreateCampaign)
	protected.GET("/campaigns/:id/results", api.GetCampaignResults)
	protected.GET("/campaigns/:id/summary", api.GetCampaignSummary)
	protected.DELETE("/campaigns/:id", api.DeleteCampaign)
	protected.GET("/campaigns/:id/complete", api.CompleteCampaign)

	// Rutas de user management
	protected.GET("/users/me", api.GetCurrentUser)
	protected.GET("/users", api.GetUsers)
	protected.GET("/users/:id", api.GetUserByID)
	protected.POST("/users", api.CreateUser)
	protected.PUT("/users/:id", api.UpdateUser)

}
