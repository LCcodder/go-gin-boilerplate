package auth_service

import (
	"context"
	"time"

	"example.com/m/internal/api/v1/adapters/repositories"
	"example.com/m/internal/api/v1/core/application/dto"
	"example.com/m/internal/api/v1/core/application/errorz"
	"example.com/m/internal/api/v1/core/application/services/user_service"
	"example.com/m/internal/api/v1/utils"
	"example.com/m/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	us user_service.UserService
	tr repositories.TokenRepository
}

func NewAuthService(us *user_service.UserService, tr *repositories.TokenRepository) *AuthService {
	return &AuthService{
		us: *us,
		tr: *tr,
	}
}

func generateAndSignToken(email string, username string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    email,
		"username": username,
		"exp":      time.Now().UTC().Add(config.Config.JWTExpiration).Unix(),
		"iat":      time.Now().UTC().Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.Config.JWTSecret))
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func (s *AuthService) Authorize(ctx context.Context, email string, password string) (*string, *errorz.Error_) {
	user, exception := s.us.GetUserByEmail(ctx, email)
	if exception != nil {
		if exception.StatusCode == 404 {
			return nil, &errorz.ErrAuthWrongCredentials
		}
		return nil, exception
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, &errorz.ErrAuthWrongCredentials
	}

	token, err := generateAndSignToken(user.Email, user.Username)
	if err != nil {
		return nil, &errorz.ErrServiceUnavailable
	}

	err = s.tr.Set(&ctx, email, *token)
	if err != nil {
		return nil, &errorz.ErrServiceUnavailable
	}

	return token, nil
}

func (s *AuthService) CheckTokenExistance(ctx context.Context, email string, token string) *errorz.Error_ {
	foundToken, err := s.tr.GetByEmail(&ctx, email)
	if err != nil {
		return &errorz.ErrServiceUnavailable
	}

	if foundToken == nil || *foundToken != token {
		return &errorz.ErrAuthInvalidToken
	}
	return nil
}
func (s *AuthService) ChangePassword(ctx context.Context, email string, oldPassword string, newPassword string) *errorz.Error_ {
	if oldPassword == newPassword {
		return &errorz.ErrAuthWrongCredentials
	}

	user, exception := s.us.GetUserByEmail(ctx, email)
	if exception != nil {
		if exception.StatusCode == 404 {
			return &errorz.ErrAuthWrongCredentials
		}
		return exception
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
		return &errorz.ErrAuthWrongCredentials
	}

	newHashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return &errorz.ErrServiceUnavailable
	}

	_, exception = s.us.UpdateUserByEmail(ctx, email, dto.UpdateUserDto{
		Password: newHashedPassword,
	})
	if exception != nil {
		return exception
	}

	err = s.tr.DeleteByEmail(&ctx, email)
	if err != nil {
		return &errorz.ErrServiceUnavailable
	}

	return nil
}
