package routes

import (
	"github.com/gin-gonic/gin"
	"golang-starter/registration/handlers"
)

// SetupRoutes registers all URL endpoints to their respective handlers (like Django urls.py)
func SetupRoutes(router *gin.Engine) {
	// Group all endpoints under /api/v1 (standard REST API convention)
	api := router.Group("/api/v1")
	{
		// 1. Welcome Greeting
		api.GET("/welcome", handlers.Welcome)

		// 2. User CRUD Endpoints
		api.POST("/users", handlers.RegisterUser)       // Create
		api.GET("/users", handlers.GetAllUsers)         // Read all
		api.GET("/users/:id", handlers.GetUserByID)     // Read by ID
		api.PUT("/users/:id", handlers.UpdateUser) 
		api.PATCH("/users/:id",handlers.UpdateUser)     // Update
		api.DELETE("/users/:id", handlers.DeleteUser)   // Delete
	}
}
