package router

import (
	docs "example.com/m/docs"
	"example.com/m/internal/api/v1/adapters/api/controllers"
	"example.com/m/internal/api/v1/infrastructure/middlewares"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const prefix string = "/api/v1"

func BindRoutes(r *gin.Engine, a *middlewares.AuthMiddleware, uc *controllers.UserController, ac *controllers.AuthController, m *controllers.MetricController) {
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Version = "v1"

	r.POST(prefix+"/users", uc.CreateUser)
	r.POST(prefix+"/auth", ac.AuthorizeUser)
	r.POST(prefix+"/auth/changePassword", a.Authenticate(), ac.ChangePassword)
	r.GET(prefix+"/users/:username", a.Authenticate(), uc.GetUserByUsername)
	r.GET(prefix+"/users/me", a.Authenticate(), uc.GetUserProfile)
	r.PATCH(prefix+"/users/me", a.Authenticate(), uc.UpdateUserProfile)
	r.GET(prefix+"/metrics", m.GetMetrics())

	ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:8000/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
