package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/khanhnp-2797/gin-realworld-api/controllers"
	"github.com/khanhnp-2797/gin-realworld-api/dto"
	"github.com/stretchr/testify/assert"
)

// TestUserController_Register_Success
func TestUserController_Register_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	InitTestConfig()

	mockUserService := NewMockUserService()
	mockUserService.RegisterFunc = func(req *dto.RegisterRequest) (*dto.UserResponse, error) {
		return &dto.UserResponse{
			User: dto.UserData{
				Username: req.User.Username,
				Email:    req.User.Email,
			},
		}, nil
	}

	controller := controllers.NewUserController(mockUserService)
	router := gin.New()
	router.POST("/users", controller.Register)

	payload := map[string]interface{}{
		"user": map[string]string{
			"username": "testuser",
			"email":    "test@example.com",
			"password": "password123",
		},
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

// TestUserController_Register_EmailConflict
func TestUserController_Register_EmailConflict(t *testing.T) {
	gin.SetMode(gin.TestMode)
	InitTestConfig()

	mockUserService := NewMockUserService()
	mockUserService.RegisterFunc = func(req *dto.RegisterRequest) (*dto.UserResponse, error) {
		return nil, errors.New("email already exists")
	}

	controller := controllers.NewUserController(mockUserService)
	router := gin.New()
	router.POST("/users", controller.Register)

	payload := map[string]interface{}{
		"user": map[string]string{
			"username": "testuser",
			"email":    "existing@example.com",
			"password": "password123",
		},
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)
}

// TestUserController_Register_UsernameConflict
func TestUserController_Register_UsernameConflict(t *testing.T) {
	gin.SetMode(gin.TestMode)
	InitTestConfig()

	mockUserService := NewMockUserService()
	mockUserService.RegisterFunc = func(req *dto.RegisterRequest) (*dto.UserResponse, error) {
		return nil, errors.New("username already exists")
	}

	controller := controllers.NewUserController(mockUserService)
	router := gin.New()
	router.POST("/users", controller.Register)

	payload := map[string]interface{}{
		"user": map[string]string{
			"username": "existinguser",
			"email":    "test@example.com",
			"password": "password123",
		},
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)
}

// TestUserController_Register_InvalidJSON
func TestUserController_Register_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	InitTestConfig()

	mockUserService := NewMockUserService()
	controller := controllers.NewUserController(mockUserService)
	router := gin.New()
	router.POST("/users", controller.Register)

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

// TestUserController_Register_ServiceError
func TestUserController_Register_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	InitTestConfig()

	mockUserService := NewMockUserService()
	mockUserService.RegisterFunc = func(req *dto.RegisterRequest) (*dto.UserResponse, error) {
		return nil, errors.New("service error")
	}

	controller := controllers.NewUserController(mockUserService)
	router := gin.New()
	router.POST("/users", controller.Register)

	payload := map[string]interface{}{
		"user": map[string]string{
			"username": "testuser",
			"email":    "test@example.com",
			"password": "password123",
		},
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

// TestUserController_GetCurrentUser_Success
func TestUserController_GetCurrentUser_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	InitTestConfig()

	mockUserService := NewMockUserService()
	mockUserService.GetUserByIDFunc = func(id uint) (*dto.UserResponse, error) {
		return &dto.UserResponse{
			User: dto.UserData{
				Username: "testuser",
				Email:    "test@example.com",
			},
		}, nil
	}

	controller := controllers.NewUserController(mockUserService)
	router := gin.New()
	router.GET("/user", func(c *gin.Context) {
		c.Set("userID", uint(1))
		controller.GetCurrentUser(c)
	})

	req := httptest.NewRequest(http.MethodGet, "/user", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// TestUserController_GetCurrentUser_Unauthorized
func TestUserController_GetCurrentUser_Unauthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)
	InitTestConfig()

	mockUserService := NewMockUserService()
	controller := controllers.NewUserController(mockUserService)
	router := gin.New()
	router.GET("/user", controller.GetCurrentUser)

	req := httptest.NewRequest(http.MethodGet, "/user", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// TestUserController_GetCurrentUser_NotFound
func TestUserController_GetCurrentUser_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	InitTestConfig()

	mockUserService := NewMockUserService()
	mockUserService.GetUserByIDFunc = func(id uint) (*dto.UserResponse, error) {
		return nil, errors.New("user not found")
	}

	controller := controllers.NewUserController(mockUserService)
	router := gin.New()
	router.GET("/user", func(c *gin.Context) {
		c.Set("userID", uint(999))
		controller.GetCurrentUser(c)
	})

	req := httptest.NewRequest(http.MethodGet, "/user", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

// TestUserController_UpdateUser_Success
func TestUserController_UpdateUser_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	InitTestConfig()

	mockUserService := NewMockUserService()
	mockUserService.UpdateUserFunc = func(id uint, req *dto.UpdateUserRequest) (*dto.UserResponse, error) {
		return &dto.UserResponse{
			User: dto.UserData{
				Username: req.User.Username,
				Email:    req.User.Email,
			},
		}, nil
	}

	controller := controllers.NewUserController(mockUserService)
	router := gin.New()
	router.PUT("/user", func(c *gin.Context) {
		c.Set("userID", uint(1))
		controller.UpdateUser(c)
	})

	payload := map[string]interface{}{
		"user": map[string]string{
			"username": "newusername",
			"email":    "newemail@example.com",
		},
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPut, "/user", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// TestUserController_UpdateUser_Unauthorized
func TestUserController_UpdateUser_Unauthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)
	InitTestConfig()

	mockUserService := NewMockUserService()
	controller := controllers.NewUserController(mockUserService)
	router := gin.New()
	router.PUT("/user", controller.UpdateUser)

	payload := map[string]interface{}{
		"user": map[string]string{
			"username": "newusername",
		},
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPut, "/user", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// TestUserController_UpdateUser_InvalidJSON
func TestUserController_UpdateUser_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	InitTestConfig()

	mockUserService := NewMockUserService()
	controller := controllers.NewUserController(mockUserService)
	router := gin.New()
	router.PUT("/user", func(c *gin.Context) {
		c.Set("userID", uint(1))
		controller.UpdateUser(c)
	})

	req := httptest.NewRequest(http.MethodPut, "/user", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

// TestUserController_UpdateUser_ServiceError
func TestUserController_UpdateUser_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	InitTestConfig()

	mockUserService := NewMockUserService()
	mockUserService.UpdateUserFunc = func(id uint, req *dto.UpdateUserRequest) (*dto.UserResponse, error) {
		return nil, errors.New("service error")
	}

	controller := controllers.NewUserController(mockUserService)
	router := gin.New()
	router.PUT("/user", func(c *gin.Context) {
		c.Set("userID", uint(1))
		controller.UpdateUser(c)
	})

	payload := map[string]interface{}{
		"user": map[string]string{
			"username": "newusername",
		},
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPut, "/user", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}
