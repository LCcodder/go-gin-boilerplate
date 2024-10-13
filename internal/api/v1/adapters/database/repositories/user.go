package repositories

import (
	"database/sql"
	"errors"
	"fmt"

	"example.com/m/internal/api/v1/core/application/dto"
	"github.com/doug-martin/goqu/v9"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *dto.UserDto) error {
	query, _, _ := goqu.Insert("users").Rows(*user).ToSQL()
	_, err := r.db.Exec(query)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (r *UserRepository) GetByUsername(username *string) (*dto.UserDto, error) {
	query, _, _ := goqu.From("users").Where(goqu.Ex{
		"username": *username,
	}).ToSQL()

	var u dto.UserDto
	err := r.db.QueryRow(query).Scan(&u.Username, &u.Email, &u.Password, &u.Birthday, &u.Sex, &u.CreatedAt, &u.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &u, nil
}

func (r *UserRepository) GetByEmail(email *string) (*dto.UserDto, error) {
	query, _, _ := goqu.From("users").Where(goqu.Ex{
		"email": *email,
	}).ToSQL()

	var u dto.UserDto
	fmt.Println(query)
	err := r.db.QueryRow(query).Scan(&u.Username, &u.Email, &u.Password, &u.Birthday, &u.Sex, &u.CreatedAt, &u.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &u, nil
}
