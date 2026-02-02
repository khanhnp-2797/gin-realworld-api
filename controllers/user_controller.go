package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khanhnp-2797/gin-realworld-api/dto"
	"github.com/khanhnp-2797/gin-realworld-api/services"
	"github.com/khanhnp-2797/gin-realworld-api/utils"
)

type UserController struct {
	userService services.UserService
}

// NewUserController - Tạo instance mới của UserController
func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// Register - Xử lý đăng ký user mới
// POST /api/users
func (uc *UserController) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ValidationError("Invalid request body"))
		return
	}

	response, err := uc.userService.Register(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetCurrentUser - Lấy thông tin user hiện tại
// GET /api/user
func (uc *UserController) GetCurrentUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.AuthError(""))
		return
	}

	response, err := uc.userService.GetUserByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.NotFoundError("User"))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateUser - Cập nhật thông tin user
// PUT /api/user
func (uc *UserController) UpdateUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.AuthError(""))
		return
	}

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ValidationError("Invalid request body"))
		return
	}

	response, err := uc.userService.UpdateUser(userID.(uint), &req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response)
}
