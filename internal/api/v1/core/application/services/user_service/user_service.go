package user_service

import (
	"context"

	"example.com/m/internal/api/v1/adapters/repositories"
	"example.com/m/internal/api/v1/core/application/dto"
	"example.com/m/internal/api/v1/core/application/errorz"
	"example.com/m/internal/api/v1/utils"
)

type UserService struct {
	r repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{r: *repo}
}

func (s *UserService) isUserUnique(email string, username string) (*bool, *errorz.Error_) {
	foundUserByEmail, err := s.r.GetByEmail(&email)
	if err != nil {
		return nil, &errorz.ErrDatabaseError
	}

	foundUserByUsername, err := s.r.GetByUsername(&username)
	if err != nil {
		return nil, &errorz.ErrDatabaseError
	}

	state := true
	if foundUserByEmail != nil || foundUserByUsername != nil {
		state = false
	}

	return &state, nil
}

func (s *UserService) CreateUser(ctx context.Context, u dto.CreateUserDto) (*dto.UserDto, *errorz.Error_) {
	userIsUnique, exception := s.isUserUnique(u.Email, u.Username)
	if exception != nil {
		return nil, exception
	}
	if !(*userIsUnique) {
		return nil, &errorz.ErrUserAlreadyExists
	}

	userToCreate := utils.CreateUserTimestamps(&u)
	userToCreate.Password, _ = utils.HashPassword(userToCreate.Password)

	err := s.r.Create(&userToCreate)
	if err != nil {
		return nil, &errorz.ErrDatabaseError
	}

	return &userToCreate, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*dto.UserDto, *errorz.Error_) {
	u, err := s.r.GetByEmail(&email)
	if err != nil {
		return nil, &errorz.ErrDatabaseError
	}

	if u == nil {
		return nil, &errorz.ErrUserNotFound
	}

	return u, nil
}

func (s *UserService) GetUserByUsername(ctx context.Context, username string) (*dto.UserDto, *errorz.Error_) {
	u, err := s.r.GetByUsername(&username)
	if err != nil {
		return nil, &errorz.ErrDatabaseError
	}

	if u == nil {
		return nil, &errorz.ErrUserNotFound
	}

	return u, nil
}

func (s *UserService) UpdateUserByEmail(ctx context.Context, email string, u dto.UpdateUserDto) (*dto.UserDto, *errorz.Error_) {
	userIsUnique, exception := s.isUserUnique(u.Email, u.Username)
	if exception != nil {
		return nil, exception
	}
	if !(*userIsUnique) {
		return nil, &errorz.ErrUserAlreadyExists
	}

	_, exception = s.GetUserByEmail(ctx, email)
	if exception != nil {
		return nil, exception
	}

	utils.UpdateUserTimestamps(&u)

	if err := s.r.UpdateByEmail(&email, &u); err != nil {
		return nil, &errorz.ErrDatabaseError
	}

	updatedUser, exception := s.GetUserByEmail(ctx, email)
	if exception != nil {
		return nil, exception
	}

	return updatedUser, nil
}
