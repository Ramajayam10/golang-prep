package main

import (
	"fmt"
	"os"

	"banking-service/config"
	"banking-service/handlers"
	"banking-service/repository"
	"banking-service/routes"
	"banking-service/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Database Connection & Migration
	config.ConnectDB()

	// 2. Dependency Injection: Repo -> Service -> Handler
	bankRepo := repository.NewPostgresBankRepository(config.DB)
	bankService := services.NewBankService(bankRepo)
	bankHandler := handlers.NewBankHandler(bankService)

	// 3. Setup Gin Router & Register Routes
	router := gin.Default()
	routes.SetupBankRoutes(router, bankHandler) 

	// 4. Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	fmt.Printf("🚀 Banking Microservice is running on http://localhost:%s\n", port)
	router.Run(":" + port)
}
