package utils

import (
	"github.com/khanhnp-2797/gin-realworld-api/dto"
)

// NewErrorResponse - Tạo error response theo RealWorld API spec
func NewErrorResponse(messages ...string) dto.ErrorResponse {
	return dto.ErrorResponse{
		Errors: dto.ErrorBody{
			Body: messages,
		},
	}
}

// ValidationError - Tạo validation error response
func ValidationError(message string) dto.ErrorResponse {
	return NewErrorResponse(message)
}

// AuthError - Tạo authentication error response
func AuthError(message string) dto.ErrorResponse {
	if message == "" {
		message = "Authentication required"
	}
	return NewErrorResponse(message)
}

// NotFoundError - Tạo not found error response
func NotFoundError(resource string) dto.ErrorResponse {
	return NewErrorResponse(resource + " not found")
}
