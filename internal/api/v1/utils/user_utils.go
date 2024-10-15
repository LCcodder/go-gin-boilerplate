package utils

import (
	"time"

	"example.com/m/internal/api/v1/core/application/dto"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserTimestamps(u *dto.CreateUserDto) dto.UserDto {
	return dto.UserDto{
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		Birthday:  u.Birthday,
		Sex:       u.Sex,
		Bio:       u.Bio,
		CreatedAt: time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		UpdatedAt: time.Now().UTC().Format("2006-01-02T15:04:05Z"),
	}
}

func UpdateUserTimestamps(u *dto.UpdateUserDto) {
	u.UpdatedAt = time.Now().UTC().Format("2006-01-02T15:04:05Z")
}

func ExcludeUserCredentials(u *dto.UserDto) dto.GetUserDto {
	return dto.GetUserDto{
		Username:  u.Username,
		Email:     u.Email,
		Birthday:  u.Birthday,
		Sex:       u.Sex,
		Bio:       u.Bio,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 7)
	return string(bytes), err
}
