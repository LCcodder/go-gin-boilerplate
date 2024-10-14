package dto

type AuthorizeUserDto struct {
	Email    string `json:"email" db:"email" binding:"required,email,max=64,min=6"`
	Password string `json:"password" db:"password" binding:"required,max=64,min=6"`
}
