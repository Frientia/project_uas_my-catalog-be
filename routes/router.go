package routes

import (
	"github.com/Frientia/my-firebase-backend/handlers"
	"github.com/Frientia/my-firebase-backend/middleware"
	"github.com/Frientia/my-firebase-backend/repositories"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// CORS Middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// 1. Init Repositories
	userRepo := repositories.NewUserRepository()

	// 2. Init Handlers
	authHandler := handlers.NewAuthHandler()
	productHandler := handlers.NewProductHandler()
	profileHandler := handlers.NewProfileHandler(userRepo) // Init Profile Handler

	// API v1 group
	v1 := r.Group("/v1")
	v1.GET("/cart", handlers.GetCart)
	v1.POST("/cart", handlers.AddToCart)
	v1.PUT("/cart/:id", handlers.UpdateCart)
	v1.DELETE("/cart/:id", handlers.RemoveCart)
	v1.DELETE("/cart", handlers.ClearCart)
	v1.POST("/orders/checkout", handlers.Checkout)
	v1.GET("/orders", handlers.GetMyOrders)
	v1.GET("/orders/:id", handlers.GetOrderDetail)
	v1.POST("/orders/:id/pay", handlers.PaymentCallback)
	
	{
		// Health check
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok", "service": "my-firebase-backend"})
		})

		// Auth routes (public)
		auth := v1.Group("/auth")
		{
			auth.POST("/verify-token", authHandler.VerifyToken)
		}

		// Protected routes (require Backend JWT/Firebase Token)
		protected := v1.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			// 3. DAFTARKAN ROUTE PROFIL DI SINI
			protected.GET("/profile", profileHandler.GetProfile)

			products := protected.Group("/products")
			{
				products.GET("", productHandler.GetAll)
				products.GET("/:id", productHandler.GetByID)

				// Admin only
				adminProducts := products.Group("")
				adminProducts.Use(middleware.AdminOnly())
				{
					adminProducts.POST("", productHandler.Create)
					adminProducts.PUT("/:id", productHandler.Update)
					adminProducts.DELETE("/:id", productHandler.Delete)
				}
			}
		}
	}

	return r
}