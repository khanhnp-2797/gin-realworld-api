package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/khanhnp-2797/gin-realworld-api/utils"
)

// AuthMiddleware - Middleware để verify JWT token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lấy Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, utils.AuthError("Authorization header required"))
			c.Abort()
			return
		}

		// Parse header: "Token <jwt>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, utils.AuthError("Invalid authorization header format. Use: Token <jwt>"))
			c.Abort()
			return
		}

		token := parts[1]

		// Validate token
		claims, err := utils.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.AuthError("Invalid or expired token"))
			c.Abort()
			return
		}

		// Lưu user info vào context
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)

		c.Next()
	}
}

// CORSMiddleware - Xử lý CORS
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
