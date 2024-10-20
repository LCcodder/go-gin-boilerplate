package main

import (
	"log"

	"example.com/m/internal/api/v1/adapters/controllers"
	"example.com/m/internal/api/v1/adapters/repositories"
	"example.com/m/internal/api/v1/core/application/services/auth_service"
	"example.com/m/internal/api/v1/core/application/services/user_service"
	"example.com/m/internal/api/v1/infrastructure/cache"
	database "example.com/m/internal/api/v1/infrastructure/database"
	"example.com/m/internal/api/v1/infrastructure/middlewares"
	"example.com/m/internal/api/v1/infrastructure/router"
	"example.com/m/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	loadEnv()
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
	metricController := controllers.NewMetricController()
	r := gin.Default()

	router.BindRoutes(r, authMiddleware, userController, authController, metricController)

	r.Run(":8000")
}
