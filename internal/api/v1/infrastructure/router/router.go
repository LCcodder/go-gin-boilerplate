package router

import (
	docs "example.com/m/docs"
	"example.com/m/internal/api/v1/adapters/controllers"
	"example.com/m/internal/api/v1/infrastructure/middlewares"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const prefix string = "/api/v1"

func BindRoutes(e *gin.Engine, a *middlewares.AuthMiddleware, uc *controllers.UserController, ac *controllers.AuthController, m *controllers.MetricController) {
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Version = "v1"

	e.POST(prefix+"/users", uc.CreateUser)
	e.POST(prefix+"/auth", ac.AuthorizeUser)
	e.POST(prefix+"/auth/changePassword", a.Authenticate(), ac.ChangePassword)
	e.GET(prefix+"/users/:username", a.Authenticate(), uc.GetUserByUsername)
	e.GET(prefix+"/users/me", a.Authenticate(), uc.GetUserProfile)
	e.PATCH(prefix+"/users/me", a.Authenticate(), uc.UpdateUserProfile)
	e.GET(prefix+"/metrics", m.GetMetrics())

	ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:8000/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
