package dto

type UserSex string

const (
	Male   UserSex = "male"
	Female UserSex = "female"
	Other  UserSex = "other"
)

type CreateUserDto struct {
	Username string  `json:"username" db:"username"`
	Email    string  `json:"email" db:"email"`
	Password string  `json:"password" db:"password"`
	Birthday string  `json:"birthday" db:"birthday"`
	Sex      UserSex `json:"sex" db:"sex"`
	Bio      string  `json:"bio" db:"bio"`
}

type UserDto struct {
	Username  string  `json:"username" db:"username"`
	Email     string  `json:"email" db:"email"`
	Password  string  `json:"password" db:"password"`
	Birthday  string  `json:"birthday" db:"birthday"`
	Sex       UserSex `json:"sex" db:"sex"`
	Bio       string  `json:"bio" db:"bio"`
	CreatedAt string  `json:"created_at" db:"created_at"`
	UpdatedAt string  `json:"updated_at" db:"updated_at"`
}

type GetUserDto struct {
	Username  string  `json:"username" db:"username"`
	Email     string  `json:"email" db:"email"`
	Birthday  string  `json:"birthday" db:"birthday"`
	Sex       UserSex `json:"sex" db:"sex"`
	Bio       string  `json:"bio" db:"bio"`
	CreatedAt string  `json:"created_at" db:"created_at"`
	UpdatedAt string  `json:"updated_at" db:"updated_at"`
}
