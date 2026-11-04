package routes

import (
	"banking-service/handlers"
	"banking-service/middleware"

	"github.com/gin-gonic/gin"
)

// SetupBankRoutes registers all endpoints for the banking service
func SetupBankRoutes(router *gin.Engine, bankHandler *handlers.BankHandler) {
	api := router.Group("/api/v1")
	{
		// 1. PUBLIC ROUTES (No Token Needed)
		api.POST("/users", bankHandler.CreateUser)
		api.POST("/login", bankHandler.Login)

		// Optional Auth Route Group
		auth := api.Group("/auth")
		{
			auth.POST("/register", bankHandler.CreateUser)
			auth.POST("/login", bankHandler.Login)
		}

		// 2. PROTECTED BANKING ROUTES (Protected with JWTMiddleware)
		protected := api.Group("")
		protected.Use(middleware.JWTMiddleware())
		{
			// User Account Management
			protected.POST("/accounts", bankHandler.CreateAccount)

			// Money Movement & Transactions
			protected.POST("/transfer", bankHandler.Transfer)

			// Statements & Passbook (Query Param: ?account_number=...)
			protected.GET("/accounts/balance", bankHandler.GetBalance)
			protected.GET("/accounts/statement", bankHandler.GetStatement)
		}
	}
}
