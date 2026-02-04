package tests

import (
	"github.com/khanhnp-2797/gin-realworld-api/dto"
	"github.com/khanhnp-2797/gin-realworld-api/models"
	"github.com/khanhnp-2797/gin-realworld-api/services"
	"gorm.io/gorm"
)

// MockUserRepository
type MockUserRepository struct {
	FindByIDFunc       func(id uint) (*models.User, error)
	FindByEmailFunc    func(email string) (*models.User, error)
	FindByUsernameFunc func(username string) (*models.User, error)
	CreateFunc         func(user *models.User) error
	UpdateFunc         func(user *models.User) error
	DeleteFunc         func(id uint) error
}

func (m *MockUserRepository) FindByID(id uint) (*models.User, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return nil, nil
}

func (m *MockUserRepository) FindByEmail(email string) (*models.User, error) {
	if m.FindByEmailFunc != nil {
		return m.FindByEmailFunc(email)
	}
	return nil, gorm.ErrRecordNotFound
}

func (m *MockUserRepository) FindByUsername(username string) (*models.User, error) {
	if m.FindByUsernameFunc != nil {
		return m.FindByUsernameFunc(username)
	}
	return nil, gorm.ErrRecordNotFound
}

func (m *MockUserRepository) Create(user *models.User) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(user)
	}
	return nil
}

func (m *MockUserRepository) Update(user *models.User) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(user)
	}
	return nil
}

func (m *MockUserRepository) Delete(id uint) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(id)
	}
	return nil
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{}
}

// MockAuthService
type MockAuthService struct {
	LoginFunc func(req *dto.LoginRequest) (*dto.UserResponse, error)
}

func (m *MockAuthService) Login(req *dto.LoginRequest) (*dto.UserResponse, error) {
	if m.LoginFunc != nil {
		return m.LoginFunc(req)
	}
	return nil, nil
}

func NewMockAuthService() *MockAuthService {
	return &MockAuthService{}
}

// MockUserService
type MockUserService struct {
	RegisterFunc    func(req *dto.RegisterRequest) (*dto.UserResponse, error)
	GetUserByIDFunc func(id uint) (*dto.UserResponse, error)
	UpdateUserFunc  func(id uint, req *dto.UpdateUserRequest) (*dto.UserResponse, error)
}

func (m *MockUserService) Register(req *dto.RegisterRequest) (*dto.UserResponse, error) {
	if m.RegisterFunc != nil {
		return m.RegisterFunc(req)
	}
	return nil, nil
}

func (m *MockUserService) GetUserByID(id uint) (*dto.UserResponse, error) {
	if m.GetUserByIDFunc != nil {
		return m.GetUserByIDFunc(id)
	}
	return nil, nil
}

func (m *MockUserService) UpdateUser(id uint, req *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	if m.UpdateUserFunc != nil {
		return m.UpdateUserFunc(id, req)
	}
	return nil, nil
}

func NewMockUserService() *MockUserService {
	return &MockUserService{}
}

var (
	_ services.AuthService = (*MockAuthService)(nil)
	_ services.UserService = (*MockUserService)(nil)
)
