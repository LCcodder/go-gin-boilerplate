package user_service

import (
	"context"
	"fmt"
	"time"

	"example.com/m/internal/api/v1/adapters/repositories"
	"example.com/m/internal/api/v1/core/application/dto"
	"example.com/m/internal/api/v1/core/application/exceptions"
	"example.com/m/internal/api/v1/infrastructure/prom"
	"example.com/m/internal/api/v1/utils"
)

type UserService struct {
	ur repositories.UserRepository
}

func NewUserService(ur *repositories.UserRepository) *UserService {
	return &UserService{ur: *ur}
}

func (s *UserService) isUserExist(email string, username string) (*bool, *exceptions.Error_) {
	foundUserByEmail, err := s.ur.GetByEmail(&email)
	if err != nil {
		return nil, &exceptions.ErrDatabaseError
	}

	foundUserByUsername, err := s.ur.GetByUsername(&username)
	if err != nil {
		return nil, &exceptions.ErrDatabaseError
	}

	state := foundUserByEmail != nil || foundUserByUsername != nil
	fmt.Println(state)
	return &state, nil
}

func (s *UserService) CreateUser(ctx context.Context, u dto.CreateUserDto) (*dto.UserDto, *exceptions.Error_) {
	userExists, exception := s.isUserExist(u.Email, u.Username)
	if exception != nil {
		return nil, exception
	}
	if *userExists {
		return nil, &exceptions.ErrUserAlreadyExists
	}

	hashedPassword, _ := utils.HashPassword(u.Password)
	userToCreate := dto.UserDto{
		Email:     u.Email,
		Password:  hashedPassword,
		Username:  u.Username,
		CreatedAt: time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		UpdatedAt: time.Now().UTC().Format("2006-01-02T15:04:05Z"),
	}

	err := s.ur.Create(&userToCreate)
	if err != nil {
		return nil, &exceptions.ErrDatabaseError
	}

	prom.UserCreatedCounter.WithLabelValues("method").Inc()
	return &userToCreate, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*dto.UserDto, *exceptions.Error_) {
	user, err := s.ur.GetByEmail(&email)
	if err != nil {
		return nil, &exceptions.ErrDatabaseError
	}

	if user == nil {
		return nil, &exceptions.ErrUserNotFound
	}

	return user, nil
}

func (s *UserService) GetUserByUsername(ctx context.Context, username string) (*dto.UserDto, *exceptions.Error_) {
	user, err := s.ur.GetByUsername(&username)
	if err != nil {
		return nil, &exceptions.ErrDatabaseError
	}

	if user == nil {
		return nil, &exceptions.ErrUserNotFound
	}

	return user, nil
}

func (s *UserService) UpdateUserByEmail(ctx context.Context, email string, u dto.UpdateUserDto) (*dto.UserDto, *exceptions.Error_) {
	userExists, exception := s.isUserExist("", u.Username)
	if exception != nil {
		return nil, exception
	}
	if *userExists {
		return nil, &exceptions.ErrUserAlreadyExists
	}

	_, exception = s.GetUserByEmail(ctx, email)
	if exception != nil {
		return nil, exception
	}

	utils.UpdateUserTimestamps(&u)

	if err := s.ur.UpdateByEmail(&email, &u); err != nil {
		return nil, &exceptions.ErrDatabaseError
	}

	updatedUser, exception := s.GetUserByEmail(ctx, email)
	if exception != nil {
		return nil, exception
	}

	return updatedUser, nil
}
