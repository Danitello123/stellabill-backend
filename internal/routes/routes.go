package routes

import (
	"github.com/gin-gonic/gin"
	"stellarbill-backend/internal/handlers"
	"stellarbill-backend/internal/middleware"
)

func Register(r *gin.Engine) {
	r.Use(corsMiddleware())

	api := r.Group("/api")
	{
		api.GET("/health", handlers.Health)

		subscriptions := api.Group("/subscriptions")
		{
			subscriptions.GET("", middleware.ValidateQuery[handlers.SubscriptionQuery](), handlers.ListSubscriptions)
			subscriptions.GET("/:id", middleware.ValidatePath[handlers.SubscriptionPath](), handlers.GetSubscription)
		}

		plans := api.Group("/plans")
		{
			plans.GET("", middleware.ValidateQuery[handlers.PlanQuery](), handlers.ListPlans)
		}
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
