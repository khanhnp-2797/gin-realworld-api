package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/khanhnp-2797/gin-realworld-api/controllers"
	"github.com/khanhnp-2797/gin-realworld-api/dto"
	"github.com/khanhnp-2797/gin-realworld-api/models"
	"github.com/khanhnp-2797/gin-realworld-api/services"
	"github.com/khanhnp-2797/gin-realworld-api/tests/helpers"
	"github.com/khanhnp-2797/gin-realworld-api/tests/mocks"
	"github.com/khanhnp-2797/gin-realworld-api/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestAuthController_Login_Success(t *testing.T) {
	// Setup
	helpers.InitTestConfig()
	mockRepo := mocks.NewMockUserRepository()
	service := services.NewAuthService(mockRepo)
	controller := controllers.NewAuthController(service)

	hashedPassword, _ := utils.HashPassword("password123")
	mockUser := &models.User{
		ID:       1,
		Username: "testuser",
		Email:    "test@example.com",
		Password: hashedPassword,
	}

	mockRepo.On("FindByEmail", "test@example.com").Return(mockUser, nil)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/users/login", controller.Login)

	loginReq := dto.LoginRequest{}
	loginReq.User.Email = "test@example.com"
	loginReq.User.Password = "password123"

	body, _ := json.Marshal(loginReq)
	req := httptest.NewRequest("POST", "/api/users/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response dto.UserResponse
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "testuser", response.User.Username)
	assert.True(t, strings.HasPrefix(response.User.Token, "Bearer "))
	mockRepo.AssertExpectations(t)
}

func TestAuthController_Login_InvalidCredentials(t *testing.T) {
	helpers.InitTestConfig()
	mockRepo := mocks.NewMockUserRepository()
	service := services.NewAuthService(mockRepo)
	controller := controllers.NewAuthController(service)

	mockRepo.On("FindByEmail", "notfound@example.com").Return(nil, gorm.ErrRecordNotFound)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/users/login", controller.Login)

	loginReq := dto.LoginRequest{}
	loginReq.User.Email = "notfound@example.com"
	loginReq.User.Password = "password123"

	body, _ := json.Marshal(loginReq)
	req := httptest.NewRequest("POST", "/api/users/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	mockRepo.AssertExpectations(t)
}
