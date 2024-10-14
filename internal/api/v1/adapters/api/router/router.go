package router

import (
	"example.com/m/internal/api/v1/adapters/api/controllers"
	"github.com/gin-gonic/gin"
)

const prefix string = "/api/v1"

func BindUserRoutes(r *gin.Engine, uc *controllers.UserController, ac *controllers.AuthController) {
	r.POST(prefix+"/users", uc.CreateUser)
	r.POST(prefix+"/auth", ac.AuthorizeUser)
}
