package models

import (
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type UserInput struct {
	ID       string `json:"id" validate:"omitempty,uuid"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=256"`
}

func (input UserInput) Validate() error {
	// TODO: Probably shouldn't create a new validator everytime
	validate := validator.New()

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
