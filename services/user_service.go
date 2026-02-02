package services

import (
	"errors"

	"github.com/khanhnp-2797/gin-realworld-api/dto"
	"github.com/khanhnp-2797/gin-realworld-api/models"
	"github.com/khanhnp-2797/gin-realworld-api/repositories"
	"github.com/khanhnp-2797/gin-realworld-api/utils"
	"gorm.io/gorm"
)

type UserService interface {
	Register(req *dto.RegisterRequest) (*dto.UserResponse, error)
	GetUserByID(id uint) (*dto.UserResponse, error)
	UpdateUser(id uint, req *dto.UpdateUserRequest) (*dto.UserResponse, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

// Register - Xử lý đăng ký user mới
func (s *userService) Register(req *dto.RegisterRequest) (*dto.UserResponse, error) {
	// Kiểm tra email đã tồn tại
	if _, err := s.userRepo.FindByEmail(req.User.Email); err == nil {
		return nil, errors.New("email already exists")
	}

	// Kiểm tra username đã tồn tại
	if _, err := s.userRepo.FindByUsername(req.User.Username); err == nil {
		return nil, errors.New("username already exists")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.User.Password)
	if err != nil {
		return nil, err
	}

	// Tạo user
	user := &models.User{
		Username: req.User.Username,
		Email:    req.User.Email,
		Password: hashedPassword,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		User: dto.UserData{
			Username: user.Username,
			Email:    user.Email,
			Bio:      user.Bio,
			Image:    user.Image,
		},
	}, nil
}

// GetUserByID - Lấy thông tin user theo ID
func (s *userService) GetUserByID(id uint) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &dto.UserResponse{
		User: dto.UserData{
			Username: user.Username,
			Email:    user.Email,
			Bio:      user.Bio,
			Image:    user.Image,
		},
	}, nil
}

// UpdateUser - Cập nhật thông tin user
func (s *userService) UpdateUser(id uint, req *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Update các fields nếu có
	if req.User.Username != "" {
		user.Username = req.User.Username
	}
	if req.User.Email != "" {
		user.Email = req.User.Email
	}
	if req.User.Bio != "" {
		user.Bio = req.User.Bio
	}
	if req.User.Image != "" {
		user.Image = req.User.Image
	}
	if req.User.Password != "" {
		hashedPassword, err := utils.HashPassword(req.User.Password)
		if err != nil {
			return nil, err
		}
		user.Password = hashedPassword
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		User: dto.UserData{
			Username: user.Username,
			Email:    user.Email,
			Bio:      user.Bio,
			Image:    user.Image,
		},
	}, nil
}
