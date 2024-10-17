package controllers

import (
	"example.com/m/internal/api/v1/core/application/dto"
	"example.com/m/internal/api/v1/core/application/services/user_service"
	"example.com/m/internal/api/v1/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	us user_service.UserService
}

func NewUserController(s *user_service.UserService) *UserController {
	return &UserController{
		us: *s,
	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user dto.CreateUserDto
	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	u, err := c.us.CreateUser(ctx, user)
	if err != nil {
		ctx.JSON(int(err.StatusCode), err)
		return
	}
	userToReturn := utils.ExcludeUserCredentials(u)
	ctx.JSON(201, &userToReturn)
}

func (c *UserController) GetUserByUsername(ctx *gin.Context) {
	username := ctx.Param("username")

	u, err := c.us.GetUserByUsername(ctx, username)
	if err != nil {
		ctx.JSON(int(err.StatusCode), err)
		return
	}

	userToReturn := utils.ExcludeUserCredentials(u)

	ctx.JSON(200, &userToReturn)
}

func (c *UserController) GetUserProfile(ctx *gin.Context) {
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

	username := p["username"].(string)

	u, err := c.us.GetUserByUsername(ctx, username)
	if err != nil {
		ctx.JSON(int(err.StatusCode), err)
		return
	}

	userToReturn := utils.ExcludeUserCredentials(u)

	ctx.JSON(200, &userToReturn)
}

func (c *UserController) UpdateUserProfile(ctx *gin.Context) {
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

	var updateData dto.UpdateUserDto
	if err := ctx.ShouldBindBodyWithJSON(&updateData); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := c.us.UpdateUserByEmail(ctx, email, updateData)
	if err != nil {
		ctx.JSON(int(err.StatusCode), err)
		return
	}

	userToReturn := utils.ExcludeUserCredentials(updatedUser)

	ctx.JSON(200, &userToReturn)
}
