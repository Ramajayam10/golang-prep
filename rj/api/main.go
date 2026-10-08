package main

import (
	"log"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"database/sql"
	"api/handler"
	"api/repository"
	"os"
)

func main() {
	connection_string := os.Getenv("database_url")
	log.Print(connection_string)

	if connection_string == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := sql.Open(
		"pgx",
		connection_string,
	)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	if err:=db.Ping(); err!=nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	userRepo := repository.NewUserRepository(db)
	userHandler := handler.NewUserHandler(userRepo)

	router.POST("/user", userHandler.CreateUser)
	router.GET("/users", userHandler.GetAllUsers)
	router.GET("/user/:id", userHandler.GetUserById)
	router.PUT("/user/:id", userHandler.UpdateUser)
	// router.DELETE("/user/:id", userHandler.HardDeleteUser)
	router.DELETE("user/:id", userHandler.SoftDeleteUser)
	router.Run(":8080")
}