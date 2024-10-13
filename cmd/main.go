package main

import (
	"example.com/m/internal/api/v1/adapters/api/controllers"
	"example.com/m/internal/api/v1/adapters/api/router"
	database "example.com/m/internal/api/v1/adapters/database/postgres"
	"example.com/m/internal/api/v1/adapters/database/repositories"
	"example.com/m/internal/api/v1/core/application/services"
	"example.com/m/internal/config"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	config.InitConfig()
	database.ConnectToDatabase()

	defer database.Db.Close()

	userRepository := repositories.NewUserRepository(database.Db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserService(userService)

	r := gin.Default()

	router.BindUserRoutes(r, userController)

	r.Run("localhost:8080")
}
