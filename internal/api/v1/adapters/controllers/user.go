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

// @BasePath /api/v1

// Create user
// @Schemes
// @Description Creates new user and returns it
// @Tags user
// @Accept json
// @Produce json
// @Param user body dto.CreateUserDto true "User data"
// @Success 201 {object} dto.GetUserDto
// @Failure 500 {object} exceptions.Error_
// @Failure 503 {object} exceptions.Error_
// @Failure 400 {object} exceptions.Error_
// @Router /users [post]
func (c *UserController) CreateUser(ctx *gin.Context) {
	var user dto.CreateUserDto
	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := c.us.CreateUser(ctx, user)
	if err != nil {
		ctx.JSON(int(err.StatusCode), err)
		return
	}
	userToReturn := utils.ExcludeUserCredentials(createdUser)
	ctx.JSON(201, &userToReturn)
}

// Get user profile (by username)
// @Schemes
// @Description Returns user profile by username (requires JWT in "Bearer" header)
// @Tags user
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} dto.GetUserDto
// @Failure 500 {object} exceptions.Error_
// @Failure 503 {object} exceptions.Error_
// @Failure 401 {object} exceptions.Error_
// @Param Authorization header string true "Bearer JWT token"
// @Router /users/{username} [get]
func (c *UserController) GetUserByUsername(ctx *gin.Context) {
	username := ctx.Param("username")

	user, err := c.us.GetUserByUsername(ctx, username)
	if err != nil {
		ctx.JSON(int(err.StatusCode), err)
		return
	}

	userToReturn := utils.ExcludeUserCredentials(user)

	ctx.JSON(200, &userToReturn)
}

// Get user profile (via jwt)
// @Schemes
// @Description Returns user profile (requires JWT in "Bearer" header)
// @Tags user
// @Produce json
// @Success 200 {object} dto.GetUserDto
// @Failure 500 {object} exceptions.Error_
// @Failure 503 {object} exceptions.Error_
// @Failure 401 {object} exceptions.Error_
// @Param Authorization header string true "Bearer JWT token"
// @Router /users/me [get]
func (c *UserController) GetUserProfile(ctx *gin.Context) {
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

	username := payload["username"].(string)

	user, err := c.us.GetUserByUsername(ctx, username)
	if err != nil {
		ctx.JSON(int(err.StatusCode), err)
		return
	}

	userToReturn := utils.ExcludeUserCredentials(user)

	ctx.JSON(200, &userToReturn)
}

// Update user profile
// @Schemes
// @Description Updates user profile and returns it (requires JWT in "Bearer" header)
// @Tags user
// @Produce json
// @Param user body dto.UpdateUserDto true "User data"
// @Success 200 {object} dto.GetUserDto
// @Failure 500 {object} exceptions.Error_
// @Failure 503 {object} exceptions.Error_
// @Failure 401 {object} exceptions.Error_
// @Param Authorization header string true "Bearer JWT token"
// @Router /users/me [patch]
func (c *UserController) UpdateUserProfile(ctx *gin.Context) {
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
