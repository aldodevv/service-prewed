package http

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(
	r *gin.Engine,
	authH *AuthHandler,
	themeH *ThemeHandler,
	contextH *ContextHandler,
	guestH *GuestHandler,
	assetH *AssetHandler,
	publicH *PublicHandler,
	contactH *ContactHandler,
	rsvpH *RSVPHandler,
) {
	// Apply global CORS middleware
	r.Use(CORSMiddleware())

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})

	v1 := r.Group("/v1")
	{
		// Public Auth routes
		v1.POST("/auth/login", authH.Login)
		v1.POST("/auth/register", authH.Register) // Helper to bootstrap admin user
		v1.POST("/auth/refresh", authH.Refresh)

		// Public Wedding Page route
		v1.GET("/public/wedding/:theme/:guest", publicH.GetPublicWedding)

		// Public RSVP route
		v1.POST("/public/rsvp", rsvpH.SubmitRSVP)

		// Public Contact message route
		v1.POST("/contacts", contactH.Create)

		// Private Admin routes (Auth Protected)
		admin := v1.Group("")
		admin.Use(AuthMiddleware())
		{
			// Auth me/logout
			admin.GET("/auth/me", authH.GetMe)
			admin.POST("/auth/logout", authH.Logout)

			// Theme management
			admin.GET("/themes", themeH.GetAll)
			admin.GET("/themes/:id", themeH.GetByID)
			admin.POST("/themes", themeH.Create)
			admin.PUT("/themes/:id", themeH.Update)
			admin.DELETE("/themes/:id", themeH.Delete)

			// Context management
			admin.GET("/contexts", contextH.GetAll)
			admin.GET("/contexts/:id", contextH.GetByID)
			admin.POST("/contexts", contextH.Create)
			admin.PUT("/contexts/:id", contextH.Update)
			admin.DELETE("/contexts/:id", contextH.Delete)

			// Guest management under a client context
			admin.GET("/contexts/:id/guests", guestH.GetAll)
			admin.POST("/contexts/:id/guests", guestH.Create)
			admin.PUT("/contexts/:id/guests/:guestId", guestH.Update)
			admin.DELETE("/contexts/:id/guests/:guestId", guestH.Delete)

			// RSVP management for context
			admin.GET("/contexts/:id/rsvps", rsvpH.GetAllByContextID)

			// Asset management
			admin.GET("/assets", assetH.GetAll)
			admin.POST("/assets", assetH.Create)
			admin.POST("/assets/upload", assetH.Upload)
			admin.PUT("/assets/:id", assetH.Update)
			admin.DELETE("/assets/:id", assetH.Delete)

			// Contact messages management
			admin.GET("/contacts", contactH.GetAll)
		}
	}
}
