package services_test

import (
	"testing"

	"github.com/khanhnp-2797/gin-realworld-api/dto"
	"github.com/khanhnp-2797/gin-realworld-api/models"
	"github.com/khanhnp-2797/gin-realworld-api/services"
	"github.com/khanhnp-2797/gin-realworld-api/tests/helpers"
	"github.com/khanhnp-2797/gin-realworld-api/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func setupUserServiceTest(t *testing.T) (*mocks.MockUserRepository, services.UserService) {
	helpers.InitTestConfig()
	mockRepo := mocks.NewMockUserRepository()
	service := services.NewUserService(mockRepo)
	return mockRepo, service
}

func TestUserService_Register_Success(t *testing.T) {
	// Setup
	mockRepo, service := setupUserServiceTest(t)

	req := &dto.RegisterRequest{}
	req.User.Username = "testuser"
	req.User.Email = "test@example.com"
	req.User.Password = "password123"

	mockRepo.On("FindByEmail", "test@example.com").Return(nil, gorm.ErrRecordNotFound)
	mockRepo.On("FindByUsername", "testuser").Return(nil, gorm.ErrRecordNotFound)
	mockRepo.On("Create", mock.AnythingOfType("*models.User")).Return(nil)

	// Act
	response, err := service.Register(req)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "testuser", response.User.Username)
	assert.Equal(t, "test@example.com", response.User.Email)
	assert.Empty(t, response.User.Token)
	mockRepo.AssertExpectations(t)
}

func TestUserService_Register_Failures(t *testing.T) {
	tests := []struct {
		name          string
		setupMock     func(*mocks.MockUserRepository)
		expectedError string
	}{
		{
			name: "Email already exists",
			setupMock: func(m *mocks.MockUserRepository) {
				m.On("FindByEmail", "test@example.com").Return(&models.User{ID: 1}, nil)
			},
			expectedError: "email already exists",
		},
		{
			name: "Username already exists",
			setupMock: func(m *mocks.MockUserRepository) {
				m.On("FindByEmail", "test@example.com").Return(nil, gorm.ErrRecordNotFound)
				m.On("FindByUsername", "testuser").Return(&models.User{ID: 1}, nil)
			},
			expectedError: "username already exists",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			mockRepo, service := setupUserServiceTest(t)
			tt.setupMock(mockRepo)

			req := &dto.RegisterRequest{}
			req.User.Username = "testuser"
			req.User.Email = "test@example.com"
			req.User.Password = "password123"

			// Act
			_, err := service.Register(req)

			// Assert
			assert.Error(t, err)
			assert.Equal(t, tt.expectedError, err.Error())
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestUserService_GetUserByID_Success(t *testing.T) {
	// Setup
	mockRepo, service := setupUserServiceTest(t)

	mockUser := &models.User{
		ID:       1,
		Username: "testuser",
		Email:    "test@example.com",
		Bio:      "Test bio",
		Image:    "image.jpg",
	}

	mockRepo.On("FindByID", uint(1)).Return(mockUser, nil)

	// Act
	response, err := service.GetUserByID(1)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "testuser", response.User.Username)
	assert.Equal(t, "test@example.com", response.User.Email)
	assert.Equal(t, "Test bio", response.User.Bio)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUserByID_NotFound(t *testing.T) {
	// Setup
	mockRepo, service := setupUserServiceTest(t)

	mockRepo.On("FindByID", uint(999)).Return(nil, gorm.ErrRecordNotFound)

	// Act
	_, err := service.GetUserByID(999)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestUserService_UpdateUser_Success(t *testing.T) {
	// Setup
	mockRepo, service := setupUserServiceTest(t)

	existingUser := &models.User{
		ID:       1,
		Username: "testuser",
		Email:    "test@example.com",
		Bio:      "Old bio",
		Image:    "old-image.jpg",
	}

	mockRepo.On("FindByID", uint(1)).Return(existingUser, nil)
	mockRepo.On("Update", mock.MatchedBy(func(u *models.User) bool {
		return u.Bio == "New bio" && u.Image == "new-image.jpg"
	})).Return(nil)

	updateReq := &dto.UpdateUserRequest{}
	updateReq.User.Bio = "New bio"
	updateReq.User.Image = "new-image.jpg"

	// Act
	response, err := service.UpdateUser(1, updateReq)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "New bio", response.User.Bio)
	assert.Equal(t, "new-image.jpg", response.User.Image)
	mockRepo.AssertExpectations(t)
}
