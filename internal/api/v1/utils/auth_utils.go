package utils

import (
	"strings"

	"example.com/m/internal/api/v1/core/application/exceptions"
	"example.com/m/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateTokenSignature(token string) *exceptions.Error_ {
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.JWTSecret), nil
	})
	if err != nil {
		return &exceptions.ErrAuthInvalidToken
	}
	return nil
}

func ExtractPayloadFromJWT(token string) (jwt.MapClaims, *exceptions.Error_) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Config.JWTSecret), nil
	})
	if err != nil {
		return nil, &exceptions.ErrAuthInvalidToken
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, nil
	} else {
		return nil, &exceptions.ErrAuthInvalidToken
	}
}

type authHeader struct {
	Token string `header:"Authorization"`
}

func ExtractTokenFromHeaders(c *gin.Context) (*string, *exceptions.Error_) {
	h := authHeader{}
	if err := c.ShouldBindHeader(&h); err != nil {
		return nil, &exceptions.ErrAuthInvalidToken
	}
	token := strings.Split(h.Token, "Bearer ")

	if len(token) < 2 {
		return nil, &exceptions.ErrAuthInvalidToken
	}

	return &token[1], nil
}
