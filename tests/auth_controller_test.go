package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/khanhnp-2797/gin-realworld-api/controllers"
	"github.com/khanhnp-2797/gin-realworld-api/dto"
	"github.com/stretchr/testify/assert"
)

// TestAuthController_Login_Success
func TestAuthController_Login_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	InitTestConfig()

	mockAuthService := NewMockAuthService()
	mockAuthService.LoginFunc = func(req *dto.LoginRequest) (*dto.UserResponse, error) {
		return &dto.UserResponse{
			User: dto.UserData{
				Username: "testuser",
				Email:    req.User.Email,
				Token:    "test-token",
			},
		}, nil
	}

	controller := controllers.NewAuthController(mockAuthService)
	router := gin.New()
	router.POST("/users/login", controller.Login)

	payload := map[string]interface{}{
		"user": map[string]string{
			"email":    "test@example.com",
			"password": "password123",
		},
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/users/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response dto.UserResponse
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "testuser", response.User.Username)
	assert.Equal(t, "test-token", response.User.Token)
}

// TestAuthController_Login_InvalidCredentials
func TestAuthController_Login_InvalidCredentials(t *testing.T) {
	gin.SetMode(gin.TestMode)
	InitTestConfig()

	mockAuthService := NewMockAuthService()
	mockAuthService.LoginFunc = func(req *dto.LoginRequest) (*dto.UserResponse, error) {
		return nil, errors.New("invalid email or password")
	}

	controller := controllers.NewAuthController(mockAuthService)
	router := gin.New()
	router.POST("/users/login", controller.Login)

	payload := map[string]interface{}{
		"user": map[string]string{
			"email":    "wrong@example.com",
			"password": "wrongpassword",
		},
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/users/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// TestAuthController_Login_InvalidJSON
func TestAuthController_Login_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	InitTestConfig()

	mockAuthService := NewMockAuthService()
	controller := controllers.NewAuthController(mockAuthService)
	router := gin.New()
	router.POST("/users/login", controller.Login)

	req := httptest.NewRequest(http.MethodPost, "/users/login", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}
