package tests

import (
	"errors"
	"testing"

	"github.com/khanhnp-2797/gin-realworld-api/dto"
	"github.com/khanhnp-2797/gin-realworld-api/models"
	"github.com/khanhnp-2797/gin-realworld-api/services"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// TestUserService_Register_Success
func TestUserService_Register_Success(t *testing.T) {
	InitTestConfig()
	userRepo := NewMockUserRepository()
	userService := services.NewUserService(userRepo)

	req := &dto.RegisterRequest{}
	req.User.Username = "testuser"
	req.User.Email = "test@example.com"
	req.User.Password = "password123"

	response, err := userService.Register(req)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "testuser", response.User.Username)
	assert.Equal(t, "test@example.com", response.User.Email)
}

// TestUserService_Register_EmailExists
func TestUserService_Register_EmailExists(t *testing.T) {
	InitTestConfig()
	userRepo := NewMockUserRepository()
	userRepo.FindByEmailFunc = func(email string) (*models.User, error) {
		return &models.User{Email: email}, nil
	}
	userService := services.NewUserService(userRepo)

	req := &dto.RegisterRequest{}
	req.User.Username = "testuser"
	req.User.Email = "existing@example.com"
	req.User.Password = "password123"

	response, err := userService.Register(req)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "email already exists", err.Error())
}

// TestUserService_Register_UsernameExists
func TestUserService_Register_UsernameExists(t *testing.T) {
	InitTestConfig()
	userRepo := NewMockUserRepository()
	userRepo.FindByUsernameFunc = func(username string) (*models.User, error) {
		return &models.User{Username: username}, nil
	}
	userService := services.NewUserService(userRepo)

	req := &dto.RegisterRequest{}
	req.User.Username = "existinguser"
	req.User.Email = "test@example.com"
	req.User.Password = "password123"

	response, err := userService.Register(req)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "username already exists", err.Error())
}

// TestUserService_Register_CreateError
func TestUserService_Register_CreateError(t *testing.T) {
	InitTestConfig()
	userRepo := NewMockUserRepository()
	userRepo.CreateFunc = func(user *models.User) error {
		return errors.New("database error")
	}
	userService := services.NewUserService(userRepo)

	req := &dto.RegisterRequest{}
	req.User.Username = "testuser"
	req.User.Email = "test@example.com"
	req.User.Password = "password123"

	response, err := userService.Register(req)

	assert.Error(t, err)
	assert.Nil(t, response)
}

// TestUserService_GetUserByID_Success
func TestUserService_GetUserByID_Success(t *testing.T) {
	InitTestConfig()
	userRepo := NewMockUserRepository()
	userRepo.FindByIDFunc = func(id uint) (*models.User, error) {
		return &models.User{
			Username: "testuser",
			Email:    "test@example.com",
			Bio:      "Test bio",
			Image:    "image.jpg",
		}, nil
	}
	userService := services.NewUserService(userRepo)

	response, err := userService.GetUserByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "testuser", response.User.Username)
	assert.Equal(t, "Test bio", response.User.Bio)
}

// TestUserService_GetUserByID_NotFound
func TestUserService_GetUserByID_NotFound(t *testing.T) {
	InitTestConfig()
	userRepo := NewMockUserRepository()
	userRepo.FindByIDFunc = func(id uint) (*models.User, error) {
		return nil, gorm.ErrRecordNotFound
	}
	userService := services.NewUserService(userRepo)

	response, err := userService.GetUserByID(999)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "user not found", err.Error())
}

// TestUserService_GetUserByID_DBError
func TestUserService_GetUserByID_DBError(t *testing.T) {
	InitTestConfig()
	userRepo := NewMockUserRepository()
	userRepo.FindByIDFunc = func(id uint) (*models.User, error) {
		return nil, errors.New("database connection error")
	}
	userService := services.NewUserService(userRepo)

	response, err := userService.GetUserByID(1)

	assert.Error(t, err)
	assert.Nil(t, response)
}

// TestUserService_UpdateUser_Success
func TestUserService_UpdateUser_Success(t *testing.T) {
	InitTestConfig()
	userRepo := NewMockUserRepository()
	userRepo.FindByIDFunc = func(id uint) (*models.User, error) {
		return &models.User{
			Username: "oldusername",
			Email:    "old@example.com",
		}, nil
	}
	userService := services.NewUserService(userRepo)

	req := &dto.UpdateUserRequest{}
	req.User.Username = "newusername"
	req.User.Bio = "New bio"

	response, err := userService.UpdateUser(1, req)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "newusername", response.User.Username)
	assert.Equal(t, "New bio", response.User.Bio)
}

// TestUserService_UpdateUser_UpdatePassword
func TestUserService_UpdateUser_UpdatePassword(t *testing.T) {
	InitTestConfig()
	userRepo := NewMockUserRepository()
	userRepo.FindByIDFunc = func(id uint) (*models.User, error) {
		return &models.User{
			Username: "testuser",
			Email:    "test@example.com",
			Password: "oldhash",
		}, nil
	}
	userService := services.NewUserService(userRepo)

	req := &dto.UpdateUserRequest{}
	req.User.Password = "newpassword123"

	response, err := userService.UpdateUser(1, req)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

// TestUserService_UpdateUser_NotFound
func TestUserService_UpdateUser_NotFound(t *testing.T) {
	InitTestConfig()
	userRepo := NewMockUserRepository()
	userRepo.FindByIDFunc = func(id uint) (*models.User, error) {
		return nil, gorm.ErrRecordNotFound
	}
	userService := services.NewUserService(userRepo)

	req := &dto.UpdateUserRequest{}
	req.User.Username = "newusername"

	response, err := userService.UpdateUser(999, req)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "user not found", err.Error())
}

// TestUserService_UpdateUser_FindError
func TestUserService_UpdateUser_FindError(t *testing.T) {
	InitTestConfig()
	userRepo := NewMockUserRepository()
	userRepo.FindByIDFunc = func(id uint) (*models.User, error) {
		return nil, errors.New("database error")
	}
	userService := services.NewUserService(userRepo)

	req := &dto.UpdateUserRequest{}
	req.User.Username = "newusername"

	response, err := userService.UpdateUser(1, req)

	assert.Error(t, err)
	assert.Nil(t, response)
}

// TestUserService_UpdateUser_SaveError
func TestUserService_UpdateUser_SaveError(t *testing.T) {
	InitTestConfig()
	userRepo := NewMockUserRepository()
	userRepo.FindByIDFunc = func(id uint) (*models.User, error) {
		return &models.User{
			Username: "testuser",
			Email:    "test@example.com",
		}, nil
	}
	userRepo.UpdateFunc = func(user *models.User) error {
		return errors.New("save error")
	}
	userService := services.NewUserService(userRepo)

	req := &dto.UpdateUserRequest{}
	req.User.Username = "newusername"

	response, err := userService.UpdateUser(1, req)

	assert.Error(t, err)
	assert.Nil(t, response)
}
