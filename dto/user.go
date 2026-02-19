package dto

type UserDTO struct {
	Username string
	Email    string
	Password string
}

type LoginUserDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
