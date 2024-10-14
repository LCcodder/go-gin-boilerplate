package main

import (
	"example.com/m/internal/api/v1/adapters/api/controllers"
	"example.com/m/internal/api/v1/adapters/api/router"
	database "example.com/m/internal/api/v1/adapters/database/postgres"
	"example.com/m/internal/api/v1/adapters/database/repositories"
	"example.com/m/internal/api/v1/core/application/services/auth_service"
	"example.com/m/internal/api/v1/core/application/services/user_service"
	"example.com/m/internal/config"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	config.InitConfig()
	database.ConnectToDatabase()

	defer database.Db.Close()

	userRepository := repositories.NewUserRepository(database.Db)
	userService := user_service.NewUserService(userRepository)
	authService := auth_service.NewAuthService(userService)
	userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(authService)
	r := gin.Default()

	router.BindUserRoutes(r, userController, authController)

	r.Run("localhost:8080")
}

// package main

// import (
// 	"log"

// 	"golang.org/x/crypto/bcrypt"
// )

// const (
// 	hash     = "$2a$07$9ox/5ISZ4mOtlDCHEUNKiOBnLCIJUrNoW.UTmSR2nqtGl7D..lHRG"
// 	password = "12345678"
// )

// func main() {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println("Password matches")
// }
