package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"vaqua/models"
)

var Db *gorm.DB

func InitDb() *gorm.DB {

	// Load .env if it exists (non-fatal)
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found, relying on system environment variables")
	} else {
		log.Println("✅ Loaded environment variables from .env")
	}

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Connect to Postgres
	var err error
	Db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("✅ Connected to database successfully!")

	// Auto-migrate your models
	if err := Db.AutoMigrate(&models.User{}, &models.Account{}, &models.Transaction{}); err != nil {
		log.Fatalf("Unable to migrate schema: %v", err)
	}

	log.Println("✅ Migration completed successfully")
	return Db

}
