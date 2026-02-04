package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig
	JWT      JWTConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type JWTConfig struct {
	Secret          string
	ExpirationHours int
}

type ServerConfig struct {
	Port string
}

var (
	AppConfig = &Config{
		JWT: JWTConfig{
			Secret:          "default-secret",
			ExpirationHours: 72,
		},
		Server: ServerConfig{
			Port: "8080",
		},
	}
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	AppConfig.Database = DatabaseConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", ""),
		DBName:   getEnv("DB_NAME", "realworld_db"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	AppConfig.JWT = JWTConfig{
		Secret:          getEnv("JWT_SECRET", "your-secret-key"),
		ExpirationHours: getEnvAsInt("JWT_EXPIRATION_HOURS", 72),
	}

	AppConfig.Server = ServerConfig{
		Port: getEnv("SERVER_PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	value, err := strconv.Atoi(getEnv(key, strconv.Itoa(defaultValue)))
	if err != nil {
		return defaultValue
	}
	return value
}
