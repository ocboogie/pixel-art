package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Avatar    string    `json:"avatar"`
	Email     string    `json:"-"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

// This is the data necessary to create a user
type UserNew struct {
	Name     string `json:"name" validate:"required,min=2,max=64"`
	Avatar   string `json:"avatar" validate:"required,min=2,max=64"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=256"`
}

// This is the data necessary to login
type UserCredentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=256"`
}

// Validate will only do basic validation, meaning that the avatar will not be
// thoroughly validated. Although, this can be done using the avatar service.
func (input UserNew) Validate(validate *validator.Validate) error {
	return validate.Struct(input)
}

func (input UserCredentials) Validate(validate *validator.Validate) error {
	return validate.Struct(input)
}
