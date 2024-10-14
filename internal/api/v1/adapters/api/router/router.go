package router

import (
	"example.com/m/internal/api/v1/adapters/api/controllers"
	"example.com/m/internal/api/v1/adapters/api/middlewares"
	"github.com/gin-gonic/gin"
)

func BindUserRoutes(r *gin.Engine, uc *controllers.UserController, ac *controllers.AuthController) {
	authRequiredGroup := r.Group("/authRequired").Use(middlewares.Authenticate())
	r.POST("/users", uc.CreateUser)
	r.POST("/auth", ac.AuthorizeUser)
	authRequiredGroup.GET("/users/:username", middlewares.Authenticate(), uc.GetUserByUsername)
}
