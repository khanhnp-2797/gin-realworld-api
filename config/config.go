package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config - Cấu trúc config của app
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

// ServerConfig - Cấu hình server
type ServerConfig struct {
	Port string
	Env  string
}

// DatabaseConfig - Cấu hình database
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// JWTConfig - Cấu hình JWT
type JWTConfig struct {
	Secret          string
	ExpirationHours int
}

var AppConfig *Config

// LoadConfig - Load config từ environment variables
func LoadConfig() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	expirationHours, _ := strconv.Atoi(getEnv("JWT_EXPIRATION_HOURS", "72"))

	AppConfig = &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
			Env:  getEnv("ENV", "development"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DBName:   getEnv("DB_NAME", "gin_realworld"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			Secret:          getEnv("JWT_SECRET", "your-secret-key"),
			ExpirationHours: expirationHours,
		},
	}
}

// getEnv - Lấy env variable với default value
func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
