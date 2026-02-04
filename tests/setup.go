package tests

import (
	"os"

	"github.com/khanhnp-2797/gin-realworld-api/config"
)

// InitTestConfig - Initialize config for tests
func InitTestConfig() {
	os.Setenv("JWT_SECRET", "test-secret-key-for-testing")
	os.Setenv("JWT_EXPIRATION_HOURS", "24")

	config.AppConfig.JWT.Secret = "test-secret-key-for-testing"
	config.AppConfig.JWT.ExpirationHours = 24
	config.AppConfig.Server.Port = "8080"
}
