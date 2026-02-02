package utils

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/khanhnp-2797/gin-realworld-api/config"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword - Hash password với bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash - So sánh password với hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// JWTClaims - Định nghĩa data trong JWT token
type JWTClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateToken - Tạo JWT token cho user
func GenerateToken(userID uint, email string) (string, error) {
	expirationTime := time.Now().Add(
		time.Duration(config.AppConfig.JWT.ExpirationHours) * time.Hour,
	)

	claims := &JWTClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.AppConfig.JWT.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken - Validate JWT token và trả về claims
func ValidateToken(tokenString string) (*JWTClaims, error) {
	// Strip "Bearer " prefix if present
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	claims := &JWTClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
