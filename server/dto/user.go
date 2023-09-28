package dto

type CreateUser struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
	FullName string `validate:"required"`
}

type LoginUser struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}
