package middlewares

import (
	"example.com/m/internal/api/v1/core/application/services/auth_service"
	"example.com/m/internal/api/v1/utils"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	as auth_service.AuthService
}

func NewAuthMiddleware(as *auth_service.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		as: *as,
	}
}

func (m *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := utils.ExtractTokenFromHeaders(c)
		if err != nil {
			c.JSON(int(err.StatusCode), err)
			c.Abort()
			return
		}

		if err := utils.ValidateTokenSignature(*token); err != nil {
			c.JSON(int(err.StatusCode), err)
			c.Abort()
			return
		}

		payload, err := utils.ExtractPayloadFromJWT(*token)
		if err != nil {
			c.JSON(int(err.StatusCode), err)
			c.Abort()
			return
		}

		email := payload["email"].(string)

		exception := m.as.CheckTokenExistance(c, email, *token)
		if exception != nil {
			c.JSON(int(exception.StatusCode), exception)
			c.Abort()
			return
		}

		c.Next()
	}
}
