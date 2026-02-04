package services_test

import (
	"strings"
	"testing"

	"github.com/khanhnp-2797/gin-realworld-api/dto"
	"github.com/khanhnp-2797/gin-realworld-api/models"
	"github.com/khanhnp-2797/gin-realworld-api/services"
	"github.com/khanhnp-2797/gin-realworld-api/tests/helpers"
	"github.com/khanhnp-2797/gin-realworld-api/tests/mocks"
	"github.com/khanhnp-2797/gin-realworld-api/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestAuthService_Login_Success(t *testing.T) {
	// Setup
	helpers.InitTestConfig()
	mockRepo := mocks.NewMockUserRepository()
	service := services.NewAuthService(mockRepo)

	hashedPassword, _ := utils.HashPassword("password123")
	mockUser := &models.User{
		ID:       1,
		Username: "testuser",
		Email:    "test@example.com",
		Password: hashedPassword,
		Bio:      "Test bio",
		Image:    "image.jpg",
	}

	mockRepo.On("FindByEmail", "test@example.com").Return(mockUser, nil)

	req := &dto.LoginRequest{}
	req.User.Email = "test@example.com"
	req.User.Password = "password123"

	// Act
	response, err := service.Login(req)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "testuser", response.User.Username)
	assert.Equal(t, "test@example.com", response.User.Email)
	assert.True(t, strings.HasPrefix(response.User.Token, "Bearer "))
	assert.NotEmpty(t, response.User.Token)
	mockRepo.AssertExpectations(t)
}

func TestAuthService_Login_UserNotFound(t *testing.T) {
	// Setup
	helpers.InitTestConfig()
	mockRepo := mocks.NewMockUserRepository()
	service := services.NewAuthService(mockRepo)

	mockRepo.On("FindByEmail", "notfound@example.com").Return(nil, gorm.ErrRecordNotFound)

	req := &dto.LoginRequest{}
	req.User.Email = "notfound@example.com"
	req.User.Password = "password123"

	// Act
	_, err := service.Login(req)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "email or password is invalid", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestAuthService_Login_WrongPassword(t *testing.T) {
	// Setup
	helpers.InitTestConfig()
	mockRepo := mocks.NewMockUserRepository()
	service := services.NewAuthService(mockRepo)

	hashedPassword, _ := utils.HashPassword("correctpassword")
	mockUser := &models.User{
		ID:       1,
		Username: "testuser",
		Email:    "test@example.com",
		Password: hashedPassword,
	}

	mockRepo.On("FindByEmail", "test@example.com").Return(mockUser, nil)

	req := &dto.LoginRequest{}
	req.User.Email = "test@example.com"
	req.User.Password = "wrongpassword"

	// Act
	_, err := service.Login(req)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "email or password is invalid", err.Error())
	mockRepo.AssertExpectations(t)
}
