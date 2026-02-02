package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/khanhnp-2797/gin-realworld-api/config"
	"github.com/khanhnp-2797/gin-realworld-api/controllers"
	"github.com/khanhnp-2797/gin-realworld-api/repositories"
	"github.com/khanhnp-2797/gin-realworld-api/routes"
	"github.com/khanhnp-2797/gin-realworld-api/services"
)

func main() {
	// Load config
	config.LoadConfig()

	// Khởi tạo database
	config.InitDatabase()

	// Chạy migrations
	config.RunMigrations()

	// Khởi tạo repositories
	userRepo := repositories.NewUserRepository(config.GetDB())
	articleRepo := repositories.NewArticleRepository(config.GetDB())

	// Khởi tạo services
	authService := services.NewAuthService(userRepo)
	userService := services.NewUserService(userRepo)
	articleService := services.NewArticleService(articleRepo, userRepo)

	// Khởi tạo controllers
	authController := controllers.NewAuthController(authService)
	userController := controllers.NewUserController(userService)
	articleController := controllers.NewArticleController(articleService)

	// Setup Gin router
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router, authController, userController, articleController)

	// Start server
	port := config.AppConfig.Server.Port
	log.Printf("Server starting on port %s...", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
