package config

import (
	"log"

	"github.com/khanhnp-2797/gin-realworld-api/models"
)

// RunMigrations - Chạy auto migration cho các models
func RunMigrations() {
	log.Println("Running database migrations...")

	err := DB.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Migrations completed successfully")
}
