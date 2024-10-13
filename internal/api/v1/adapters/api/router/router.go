package router

import (
	"example.com/m/internal/api/v1/adapters/api/controllers"
	"github.com/gin-gonic/gin"
)

const prefix string = "/api/v1"

func BindUserRoutes(r *gin.Engine, c *controllers.UserController) {
	r.POST(prefix+"/users", c.CreateUser)
}
