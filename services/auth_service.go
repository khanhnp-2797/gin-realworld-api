package services

import (
	"errors"

	"github.com/khanhnp-2797/gin-realworld-api/dto"
	"github.com/khanhnp-2797/gin-realworld-api/repositories"
	"github.com/khanhnp-2797/gin-realworld-api/utils"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(req *dto.LoginRequest) (*dto.UserResponse, error)
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

// Login - Xử lý đăng nhập
func (s *authService) Login(req *dto.LoginRequest) (*dto.UserResponse, error) {
	// Tìm user theo email
	user, err := s.userRepo.FindByEmail(req.User.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("email or password is invalid")
		}
		return nil, err
	}

	// Kiểm tra password
	if !utils.CheckPasswordHash(req.User.Password, user.Password) {
		return nil, errors.New("email or password is invalid")
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		User: dto.UserData{
			Email:    user.Email,
			Token:    "Bearer " + token,
			Username: user.Username,
			Bio:      user.Bio,
			Image:    user.Image,
		},
	}, nil
}
