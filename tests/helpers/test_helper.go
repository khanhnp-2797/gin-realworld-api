package helpers

import (
	"sync"

	"github.com/khanhnp-2797/gin-realworld-api/config"
)

var once sync.Once

// InitTestConfig khởi tạo config cho testing
func InitTestConfig() {
	once.Do(func() {
		if config.AppConfig == nil {
			config.AppConfig = &config.Config{}
		}
		config.AppConfig.JWT.Secret = "test-secret-key-for-testing-only"
		config.AppConfig.JWT.ExpirationHours = 24
		config.AppConfig.Server.Port = "8080"
	})
}
