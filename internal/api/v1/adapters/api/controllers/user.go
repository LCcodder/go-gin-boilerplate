package controllers

import (
	"example.com/m/internal/api/v1/core/application/dto"
	"example.com/m/internal/api/v1/core/application/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	s services.UserService
}

func NewUserService(s *services.UserService) *UserController {
	return &UserController{
		s: *s,
	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user dto.CreateUserDto
	ctx.BindJSON(&user)
	res, err := c.s.CreateUser(ctx, user)
	if err != nil {
		ctx.JSON(int(err.StatusCode), err)
		return
	}

	ctx.JSON(201, res)
}
