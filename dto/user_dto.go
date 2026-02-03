package dto

// RegisterRequest - Yêu cầu đăng ký
type RegisterRequest struct {
	User struct {
		Username string `json:"username" binding:"required,min=3,max=50"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	} `json:"user" binding:"required"`
}

// LoginRequest - Yêu cầu đăng nhập
type LoginRequest struct {
	User struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	} `json:"user" binding:"required"`
}

// UpdateUserRequest - Yêu cầu cập nhật user
type UpdateUserRequest struct {
	User struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Bio      string `json:"bio"`
		Image    string `json:"image"`
		Password string `json:"password"`
	} `json:"user" binding:"required"`
}

// UserDataWithoutToken - Dùng cho Register, GetCurrentUser, UpdateUser (không có token)
type UserDataWithoutToken struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}

// UserDataWithToken - Dùng cho Login (có token)
type UserDataWithToken struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
	Token    string `json:"token"`
}

// UserData - Unified user data structure (with optional token)
type UserData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
	Token    string `json:"token,omitempty"` // omitempty để không hiển thị khi empty
}

// UserResponse - Generic user response
type UserResponse struct {
	User UserData `json:"user"`
}

// RegisterResponse - Response cho register (không token)
type RegisterResponse struct {
	User UserDataWithoutToken `json:"user"`
}

// LoginResponse - Response cho login (có token)
type LoginResponse struct {
	User UserDataWithToken `json:"user"`
}

// GetUserResponse - Response cho get user (không token)
type GetUserResponse struct {
	User UserDataWithoutToken `json:"user"`
}

// UpdateUserResponse - Response cho update user (không token)
type UpdateUserResponse struct {
	User UserDataWithoutToken `json:"user"`
}
