package models

import (
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"-"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

// This is the data necessary to create a user
type UserNew struct {
	Name     string `json:"name" validate:"required,min=2,max=64"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=256"`
}

// This is the data necessary to login
type UserCredentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=256"`
}

func (input UserNew) Validate(validate *validator.Validate) error {
	return validate.Struct(input)
}

func (input UserCredentials) Validate(validate *validator.Validate) error {
	return validate.Struct(input)
}

// func (user *UserInput) () error {
// 	if user.ID == "" {
// 		uuid, err := uuid.NewRandom()
// 		if err != nil {
// 			return err
// 		}

// 		user.ID = uuid.String()
// 	}

// 	return nil
// }
