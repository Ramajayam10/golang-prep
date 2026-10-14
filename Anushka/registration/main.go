package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"golang-starter/registration/config"
	"golang-starter/registration/routes"
)

func main() {
	// 1. Connect to PostgreSQL and run AutoMigrate (creates 'users' table)
	config.ConnectDB()

	// 2. Initialize the default Gin engine (includes Logger and Crash Recovery middleware)
	router := gin.Default()

	// 3. Attach our CRUD and Welcome routes
	routes.SetupRoutes(router)

	// 4. Read port from .env (fallback to 8080 if not set)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 5. Start the web server
	log.Printf("Server is running on http://localhost:%s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
