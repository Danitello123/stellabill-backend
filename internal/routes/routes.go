package routes

import (
	"github.com/gin-gonic/gin"
	"stellarbill-backend/internal/auth"
	"stellarbill-backend/internal/handlers"
)

func Register(r *gin.Engine) {
	r.Use(corsMiddleware())

	api := r.Group("/api")
	{
		api.GET("/health", handlers.Health)

		// Public read (user + admin)
		api.GET("/plans",
			auth.RequirePermission(auth.PermReadPlans),
			handlers.ListPlans,
		)

		api.GET("/subscriptions",
			auth.RequirePermission(auth.PermReadSubscriptions),
			handlers.ListSubscriptions,
		)

		api.GET("/subscriptions/:id",
			auth.RequirePermission(auth.PermReadSubscriptions),
			handlers.GetSubscription,
		)

		// Example future admin-only endpoints:
		// api.POST("/plans", auth.RequirePermission(auth.PermManagePlans), ...)
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
