package utils

import (
	"time"

	"example.com/m/internal/api/v1/core/application/dto"
	"golang.org/x/crypto/bcrypt"
)

func UpdateUserTimestamps(u *dto.UpdateUserDto) {
	u.UpdatedAt = time.Now().UTC().Format("2006-01-02T15:04:05Z")
}

func ExcludeUserCredentials(u *dto.UserDto) dto.GetUserDto {
	// copy all fields except password
	return dto.GetUserDto{
		Username:  u.Username,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 7)
	return string(bytes), err
}
