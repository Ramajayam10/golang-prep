package routes

import (
	"github.com/gin-gonic/gin"
	"golang-starter/registration/handlers"
	"golang-starter/registration/middleware"
)

func SetupRoutes(router *gin.Engine) {
	// Groups all endpoints under /api/v1 
	api := router.Group("/api/v1")
	{
		// 1. PUBLIC ROUTES
		api.GET("/welcome", handlers.Welcome)
		api.POST("/users", handlers.RegisterUser) // Registration (Signup)
		api.POST("/login", handlers.Login)        // Login (returns JWT Token)

		// 2. PROTECTED ROUTES 
		protected := api.Group("/")
		protected.Use(middleware.JWTMiddleware())
		{
			protected.GET("/users", handlers.GetAllUsers)         // Read all
			protected.GET("/users/:id", handlers.GetUserByID)     // Read by ID
			protected.PUT("/users/:id", handlers.UpdateUser)      // Update
			protected.PATCH("/users/:id", handlers.UpdateUser)    // Partial Update
			protected.DELETE("/users/:id", handlers.DeleteUser)   // Delete
		}
	}
}
