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

type Router struct {
	e  *gin.Engine
	am middlewares.AuthMiddleware
}

func NewRouter(e *gin.Engine, am *middlewares.AuthMiddleware) *Router {
	return &Router{
		e:  e,
		am: *am,
	}
}

func (r *Router) BindUserRoutes(uc *controllers.UserController) {
	r.e.POST(prefix+"/users", uc.CreateUser)
	r.e.GET(prefix+"/users/:username", r.am.Authenticate(), uc.GetUserByUsername)
	r.e.GET(prefix+"/users/me", r.am.Authenticate(), uc.GetUserProfile)
	r.e.PATCH(prefix+"/users/me", r.am.Authenticate(), uc.UpdateUserProfile)
}

func (r *Router) BindAuthRoutes(ac *controllers.AuthController) {
	r.e.POST(prefix+"/auth", ac.AuthorizeUser)
	r.e.PATCH(prefix+"/auth/changePassword", r.am.Authenticate(), ac.ChangePassword)
}

func (r *Router) BindSwaggerRoutes() {
	docs.SwaggerInfo.BasePath = prefix
	docs.SwaggerInfo.Version = "v1"

	ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:8000/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))
	r.e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func (r *Router) BindMetricsRoutes(m *controllers.MetricController) {
	r.e.GET(prefix+"/metrics", m.GetMetrics())
}
