package router

import (
	"example.com/m/internal/api/v1/adapters/api/controllers"
	"example.com/m/internal/api/v1/adapters/api/middlewares"
	"github.com/gin-gonic/gin"
)

const prefix string = "/api/v1"

func BindRoutes(r *gin.Engine, uc *controllers.UserController, ac *controllers.AuthController) {
	r.POST(prefix+"/users", uc.CreateUser)
	r.POST(prefix+"/auth", ac.AuthorizeUser)
	r.GET(prefix+"/users/:username", middlewares.Authenticate(), uc.GetUserByUsername)
	r.GET(prefix+"/users/me", middlewares.Authenticate(), uc.GetUserProfile)
}
