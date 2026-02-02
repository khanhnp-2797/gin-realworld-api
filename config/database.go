package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDatabase - Khởi tạo kết nối database
func InitDatabase() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		AppConfig.Database.Host,
		AppConfig.Database.Port,
		AppConfig.Database.User,
		AppConfig.Database.Password,
		AppConfig.Database.DBName,
		AppConfig.Database.SSLMode,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connected successfully")
}

// GetDB - Lấy database instance
func GetDB() *gorm.DB {
	return DB
}
