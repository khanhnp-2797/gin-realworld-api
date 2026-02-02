package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/khanhnp-2797/gin-realworld-api/controllers"
	"github.com/khanhnp-2797/gin-realworld-api/middlewares"
)

// SetupRoutes - Cấu hình tất cả các routes
func SetupRoutes(
	router *gin.Engine,
	authController *controllers.AuthController,
	userController *controllers.UserController,
	articleController *controllers.ArticleController,
) {
	// CORS middleware
	router.Use(middlewares.CORSMiddleware())

	// API routes
	api := router.Group("/api")
	{
		// Public routes
		users := api.Group("/users")
		{
			users.POST("", userController.Register)    // POST /api/users
			users.POST("/login", authController.Login) // POST /api/users/login
		}

		// Articles routes
		articles := api.Group("/articles")
		{
			// Public article routes
			articles.GET("/:slug", articleController.GetArticle) // GET /api/articles/:slug

			// Protected article routes
			articles.Use(middlewares.AuthMiddleware())
			articles.GET("/feed", articleController.GetFeed)           // GET /api/articles/feed
			articles.POST("", articleController.CreateArticle)         // POST /api/articles
			articles.PUT("/:slug", articleController.UpdateArticle)    // PUT /api/articles/:slug
			articles.DELETE("/:slug", articleController.DeleteArticle) // DELETE /api/articles/:slug
		}

		// Protected user routes
		user := api.Group("/user")
		user.Use(middlewares.AuthMiddleware())
		{
			user.GET("", userController.GetCurrentUser) // GET /api/user
			user.PUT("", userController.UpdateUser)     // PUT /api/user
		}
	}
}
