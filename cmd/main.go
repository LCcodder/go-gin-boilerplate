package main

import (
	"example.com/m/internal/api/v1/adapters/api/controllers"
	"example.com/m/internal/api/v1/adapters/api/middlewares"
	"example.com/m/internal/api/v1/adapters/api/router"
	"example.com/m/internal/api/v1/adapters/cache"
	database "example.com/m/internal/api/v1/adapters/database"
	"example.com/m/internal/api/v1/adapters/repositories"
	"example.com/m/internal/api/v1/core/application/services/auth_service"
	"example.com/m/internal/api/v1/core/application/services/user_service"
	"example.com/m/internal/config"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	config.InitConfig()
	database.ConnectToDatabase()
	cache.ConnectToRedis()

	defer database.Db.Close()

	userRepository := repositories.NewUserRepository(database.Db)
	tokenRepository := repositories.NewTokenRepository(cache.Redis)
	userService := user_service.NewUserService(userRepository)
	authService := auth_service.NewAuthService(userService, tokenRepository)
	authMiddleware := middlewares.NewAuthMiddleware(authService)
	userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(authService)
	r := gin.Default()

	router.BindRoutes(r, authMiddleware, userController, authController)

	r.Run("localhost:8000")
}
