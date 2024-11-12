package controllers

import (
	"example.com/m/internal/api/v1/core/application/dto"
	"example.com/m/internal/api/v1/core/application/services/auth_service"
	"example.com/m/internal/api/v1/utils"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	as auth_service.AuthService
}

func NewAuthController(as *auth_service.AuthService) *AuthController {
	return &AuthController{
		as: *as,
	}
}

// @BasePath /api/v1

// Authorize user
// @Schemes
// @Description Authorizes user and returns JWT
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body dto.AuthorizeUserDto true "User credentials"
// @Success 200 {object} dto.GetUserDto
// @Failure 500 {object} exceptions.Error_
// @Failure 503 {object} exceptions.Error_
// @Failure 400 {object} exceptions.Error_
// @Router /auth [post]
func (c *AuthController) AuthorizeUser(ctx *gin.Context) {
	var credentials dto.AuthorizeUserDto
	if err := ctx.ShouldBindBodyWithJSON(&credentials); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := c.as.Authorize(ctx, credentials.Email, credentials.Password)
	if err != nil {
		ctx.JSON(int(err.StatusCode), err)
		return
	}

	ctx.JSON(200, gin.H{
		"token": token,
	})
}

// Change password
// @Schemes
// @Description Changes user password and makes current token invalid
// @Tags auth
// @Accept json
// @Produce json
// @Param passwords body dto.ChangeUserPasswordDto true "User passwords"
// @Success 200
// @Failure 500 {object} exceptions.Error_
// @Failure 503 {object} exceptions.Error_
// @Failure 400 {object} exceptions.Error_
// @Failure 401 {object} exceptions.Error_
// @Param Authorization header string true "Bearer JWT token"
// @Router /auth/changePassword [post]
func (c *AuthController) ChangePassword(ctx *gin.Context) {
	var passwords dto.ChangeUserPasswordDto
	if err := ctx.ShouldBindBodyWithJSON(&passwords); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.ExtractTokenFromHeaders(ctx)
	if err != nil {
		ctx.JSON(int(err.StatusCode), err)
		return
	}
	payload, err := utils.ExtractPayloadFromJWT(*token)
	if err != nil {
		ctx.JSON(int(err.StatusCode), err)
		return
	}
	email := payload["email"].(string)

	if err := c.as.ChangePassword(ctx, email, passwords.OldPassword, passwords.NewPassword); err != nil {
		ctx.JSON(int(err.StatusCode), err)
		return
	}

	ctx.JSON(200, gin.H{
		"success": true,
	})
}
