package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khanhnp-2797/gin-realworld-api/dto"
	"github.com/khanhnp-2797/gin-realworld-api/services"
	"github.com/khanhnp-2797/gin-realworld-api/utils"
)

type AuthController struct {
	authService services.AuthService
}

// NewAuthController - Tạo instance mới của AuthController
func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// Login - Xử lý đăng nhập
// POST /api/users/login
func (ac *AuthController) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ValidationError("Invalid request body"))
		return
	}

	response, err := ac.authService.Login(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response)
}
