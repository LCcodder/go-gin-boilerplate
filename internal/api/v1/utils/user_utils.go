package utils

import (
	"time"

	"example.com/m/internal/api/v1/core/application/dto"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserTimestamps(user *dto.CreateUserDto) dto.UserDto {
	return dto.UserDto{
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		Birthday:  user.Birthday,
		Sex:       user.Sex,
		Bio:       user.Bio,
		CreatedAt: time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		UpdatedAt: time.Now().UTC().Format("2006-01-02T15:04:05Z"),
	}
}

func ExcludeUserCredentials(user *dto.UserDto) dto.GetUserDto {
	return dto.GetUserDto{
		Username:  user.Username,
		Email:     user.Email,
		Birthday:  user.Birthday,
		Sex:       user.Sex,
		Bio:       user.Bio,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
