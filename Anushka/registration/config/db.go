package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"golang-starter/registration/models"
)

// Global database connection pointer (accessible across our project)
var DB *gorm.DB

func ConnectDB() {
	// 1. Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}

	// 2. Read variables from .env
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	// 3. Build connection string (DSN)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode)

	// 4. Open connection with GORM
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	// 5. AutoMigrate creates/updates tables automatically (like Django 'python manage.py migrate')
	err = database.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	DB = database
	fmt.Println(" Connected to PostgreSQL database successfully!")
}
