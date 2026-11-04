package config

import (
	"fmt"
	"log"
	"os"

	"banking-service/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// 1. Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: No .env file found, reading from system env")
	}

	// 2. Read database environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	// 3. Build connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode)

	// 4. Connect to PostgreSQL
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to PostgreSQL:", err)
	}

	// 5. AutoMigrate creates all 3 tables with their Foreign Keys & Indexes!
	err = database.AutoMigrate(
		&models.User{},
		&models.BankAccount{},
		&models.Transaction{},
		&models.UserCredential{},
	)
	if err != nil {
		log.Fatal("❌ AutoMigrate failed:", err)
	}

	DB = database
	fmt.Println("✅ Successfully connected to PostgreSQL & migrated all 3 tables!")
}
