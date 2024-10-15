package dto

type CreateUserDto struct {
	Username string `json:"username" db:"username" binding:"required,max=32,min=6"`
	Email    string `json:"email" db:"email" binding:"required,email,max=64,min=6"`
	Password string `json:"password" db:"password" binding:"required,max=64,min=6"`
	Birthday string `json:"birthday" db:"birthday" binding:"required,min=8,max=10"`
	Sex      string `json:"sex" db:"sex" binding:"required,min=4,max=6"`
	Bio      string `json:"bio" db:"bio" binding:"required,max=100"`
}

type UserDto struct {
	Username  string `json:"username" db:"username"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	Birthday  string `json:"birthday" db:"birthday"`
	Sex       string `json:"sex" db:"sex"`
	Bio       string `json:"bio" db:"bio"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

type GetUserDto struct {
	Username  string `json:"username" db:"username"`
	Email     string `json:"email" db:"email"`
	Birthday  string `json:"birthday" db:"birthday"`
	Sex       string `json:"sex" db:"sex"`
	Bio       string `json:"bio" db:"bio"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

type UpdateUserDto struct {
	Username string `json:"username,omitempty" db:"username" binding:"omitempty,max=32,min=6"`
	Email    string `json:"email,omitempty" db:"email" binding:"omitempty,email,max=64,min=6"`
	Birthday string `json:"birthday,omitempty" db:"birthday" binding:"omitempty,min=8,max=10"`
	Sex      string `json:"sex,omitempty" db:"sex" binding:"omitempty,min=4,max=6"`
	Bio      string `json:"bio,omitempty" db:"bio" binding:"omitempty,max=100"`
	Password string `json:"password" db:"password"`
	// sets in service automatically
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}
