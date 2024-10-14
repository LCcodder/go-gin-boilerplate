package utils

import (
	"strings"

	"example.com/m/internal/api/v1/core/application/errorz"
	"example.com/m/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateTokenSignature(token string) *errorz.Error_ {
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.JWTSecret), nil
	})
	if err != nil {
		return &errorz.ErrAuthInvalidToken
	}
	return nil
}

func ExtractPayloadFromJWT(token string) (jwt.MapClaims, *errorz.Error_) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Config.JWTSecret), nil
	})
	if err != nil {
		return nil, &errorz.ErrAuthInvalidToken
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		return claims, nil
	} else {
		return nil, &errorz.ErrAuthInvalidToken
	}
}

type authHeader struct {
	Token string `header:"Authorization"`
}

func ExtractTokenFromHeaders(c *gin.Context) (*string, *errorz.Error_) {
	h := authHeader{}
	if err := c.ShouldBindHeader(&h); err != nil {
		return nil, &errorz.ErrAuthInvalidToken
	}
	token := strings.Split(h.Token, "Bearer ")

	if len(token) < 2 {
		return nil, &errorz.ErrAuthInvalidToken
	}

	return &token[1], nil
}
