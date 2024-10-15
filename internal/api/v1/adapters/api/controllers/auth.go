package controllers

import (
	"example.com/m/internal/api/v1/core/application/dto"
	"example.com/m/internal/api/v1/core/application/services/auth_service"
	"example.com/m/internal/api/v1/utils"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	s auth_service.AuthService
}

func NewAuthController(s *auth_service.AuthService) *AuthController {
	return &AuthController{
		s: *s,
	}
}

func (c *AuthController) AuthorizeUser(ctx *gin.Context) {
	var credentials dto.AuthorizeUserDto
	if err := ctx.ShouldBindBodyWithJSON(&credentials); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	t, err := c.s.Authorize(ctx, credentials.Email, credentials.Password)
	if err != nil {
		ctx.JSON(int(err.StatusCode), err)
		return
	}

	ctx.JSON(200, gin.H{
		"token": t,
	})
}

func (c *AuthController) ChangePassword(ctx *gin.Context) {
	var passwords dto.ChangeUserPassword
	if err := ctx.ShouldBindBodyWithJSON(&passwords); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	t, err := utils.ExtractTokenFromHeaders(ctx)
	if err != nil {
		ctx.JSON(int(err.StatusCode), err)
		return
	}
	p, err := utils.ExtractPayloadFromJWT(*t)
	if err != nil {
		ctx.JSON(int(err.StatusCode), err)
		return
	}
	email := p["email"].(string)

	if err := c.s.ChangePassword(ctx, email, passwords.OldPassword, passwords.NewPassword); err != nil {
		ctx.JSON(int(err.StatusCode), err)
		return
	}

	ctx.JSON(200, gin.H{
		"success": true,
	})
}
