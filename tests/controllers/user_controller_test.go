package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/khanhnp-2797/gin-realworld-api/controllers"
	"github.com/khanhnp-2797/gin-realworld-api/dto"
	"github.com/khanhnp-2797/gin-realworld-api/models"
	"github.com/khanhnp-2797/gin-realworld-api/services"
	"github.com/khanhnp-2797/gin-realworld-api/tests/helpers"
	"github.com/khanhnp-2797/gin-realworld-api/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestUserController_Register_Success(t *testing.T) {
	// Setup
	helpers.InitTestConfig()
	mockRepo := mocks.NewMockUserRepository()
	service := services.NewUserService(mockRepo)
	controller := controllers.NewUserController(service)

	mockRepo.On("FindByEmail", "test@example.com").Return(nil, gorm.ErrRecordNotFound)
	mockRepo.On("FindByUsername", "testuser").Return(nil, gorm.ErrRecordNotFound)
	mockRepo.On("Create", mock.AnythingOfType("*models.User")).Return(nil)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/users", controller.Register)

	registerReq := dto.RegisterRequest{}
	registerReq.User.Username = "testuser"
	registerReq.User.Email = "test@example.com"
	registerReq.User.Password = "password123"

	body, _ := json.Marshal(registerReq)
	req := httptest.NewRequest("POST", "/api/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusCreated, w.Code)

	var response dto.UserResponse
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "testuser", response.User.Username)
	assert.Equal(t, "test@example.com", response.User.Email)
	assert.Empty(t, response.User.Token)
	mockRepo.AssertExpectations(t)
}

func TestUserController_Register_EmailConflict(t *testing.T) {
	helpers.InitTestConfig()
	mockRepo := mocks.NewMockUserRepository()
	service := services.NewUserService(mockRepo)
	controller := controllers.NewUserController(service)

	mockRepo.On("FindByEmail", "test@example.com").Return(&models.User{ID: 1}, nil)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/users", controller.Register)

	registerReq := dto.RegisterRequest{}
	registerReq.User.Username = "testuser"
	registerReq.User.Email = "test@example.com"
	registerReq.User.Password = "password123"

	body, _ := json.Marshal(registerReq)
	req := httptest.NewRequest("POST", "/api/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)
}

func TestUserController_GetCurrentUser_Success(t *testing.T) {
	helpers.InitTestConfig()
	mockRepo := mocks.NewMockUserRepository()
	service := services.NewUserService(mockRepo)
	controller := controllers.NewUserController(service)

	mockUser := &models.User{
		ID:       1,
		Username: "testuser",
		Email:    "test@example.com",
		Bio:      "Test bio",
	}

	mockRepo.On("FindByID", uint(1)).Return(mockUser, nil)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/user", func(c *gin.Context) {
		c.Set("userID", uint(1))
		controller.GetCurrentUser(c)
	})

	req := httptest.NewRequest("GET", "/api/user", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response dto.UserResponse
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "testuser", response.User.Username)
}

func TestUserController_UpdateUser_Success(t *testing.T) {
	helpers.InitTestConfig()
	mockRepo := mocks.NewMockUserRepository()
	service := services.NewUserService(mockRepo)
	controller := controllers.NewUserController(service)

	existingUser := &models.User{
		ID:       1,
		Username: "testuser",
		Email:    "test@example.com",
		Bio:      "Old bio",
	}

	mockRepo.On("FindByID", uint(1)).Return(existingUser, nil)
	mockRepo.On("Update", mock.AnythingOfType("*models.User")).Return(nil)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.PUT("/api/user", func(c *gin.Context) {
		c.Set("userID", uint(1))
		controller.UpdateUser(c)
	})

	updateReq := dto.UpdateUserRequest{}
	updateReq.User.Bio = "New bio"

	body, _ := json.Marshal(updateReq)
	req := httptest.NewRequest("PUT", "/api/user", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockRepo.AssertExpectations(t)
}
