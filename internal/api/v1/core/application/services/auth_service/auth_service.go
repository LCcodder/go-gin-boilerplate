package auth_service

import (
	"context"
	"fmt"
	"time"

	"example.com/m/internal/api/v1/adapters/repositories"
	"example.com/m/internal/api/v1/core/application/errorz"
	"example.com/m/internal/api/v1/core/application/services/user_service"
	"example.com/m/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	s user_service.UserService
	r *repositories.TokenRepository
}

func NewAuthService(s *user_service.UserService, r *repositories.TokenRepository) *AuthService {
	return &AuthService{
		s: *s,
		r: r,
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
	u, exception := s.s.GetUserByEmail(ctx, email)
	if exception != nil {
		if exception.StatusCode == 404 {
			return nil, &errorz.ErrAuthWrongCredentials
		}
		return nil, exception
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return nil, &errorz.ErrAuthWrongCredentials
	}

	token, err := generateAndSignToken(u.Email, u.Username)
	if err != nil {
		return nil, &errorz.ErrServiceUnavailable
	}

	err = s.r.Set(&ctx, email, *token)
	if err != nil {
		return nil, &errorz.ErrServiceUnavailable
	}

	return token, nil
}

func (s *AuthService) CheckTokenExistance(ctx context.Context, email string, token string) *errorz.Error_ {
	t, err := s.r.GetByEmail(&ctx, email)
	fmt.Println(*t)
	if err != nil {
		return &errorz.ErrServiceUnavailable
	}
	if *t != token {
		return &errorz.ErrAuthInvalidToken
	}
	return nil
}
func (s *AuthService) ChangePassword(ctx context.Context) {}
