package middlewares

import (
	"example.com/m/internal/api/v1/utils"
	"github.com/gin-gonic/gin"
)

type authHeader struct {
	Token string `header:"Authorization"`
}

func Authenticate() gin.HandlerFunc {
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

		c.Next()
	}
}
