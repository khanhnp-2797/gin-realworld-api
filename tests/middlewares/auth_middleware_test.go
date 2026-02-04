package middlewares_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/khanhnp-2797/gin-realworld-api/middlewares"
	"github.com/khanhnp-2797/gin-realworld-api/tests/helpers"
	"github.com/khanhnp-2797/gin-realworld-api/utils"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware_ValidToken(t *testing.T) {
	helpers.InitTestConfig()

	token, _ := utils.GenerateToken(1, "test@example.com")

	// GenerateToken returns plain JWT, middleware expects "Bearer " prefix
	authHeader := fmt.Sprintf("Bearer %s", token)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(middlewares.AuthMiddleware())
	router.GET("/api/protected", func(c *gin.Context) {
		userID, _ := c.Get("userID")
		c.JSON(http.StatusOK, gin.H{"userID": userID})
	})

	req := httptest.NewRequest("GET", "/api/protected", nil)
	req.Header.Set("Authorization", authHeader)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAuthMiddleware_MissingToken(t *testing.T) {
	helpers.InitTestConfig()

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(middlewares.AuthMiddleware())
	router.GET("/api/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	req := httptest.NewRequest("GET", "/api/protected", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	helpers.InitTestConfig()

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(middlewares.AuthMiddleware())
	router.GET("/api/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	req := httptest.NewRequest("GET", "/api/protected", nil)
	req.Header.Set("Authorization", "Bearer invalid.token.here")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthMiddleware_WrongHeaderFormat(t *testing.T) {
	helpers.InitTestConfig()

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(middlewares.AuthMiddleware())
	router.GET("/api/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	req := httptest.NewRequest("GET", "/api/protected", nil)
	req.Header.Set("Authorization", "InvalidFormat token")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
