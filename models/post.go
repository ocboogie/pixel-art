package models

import (
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type Post struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userID" db:"user_id"`
	Title     string    `json:"title"`
	Data      []byte    `json:"data"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type PostNew struct {
	UserID string `json:"userId" validate:"required,uuid"`
	Title  string `json:"title" validate:"required,min=2,max=256"`
	// TODO: validate Data
	Data string `json:"data"`
}

func (input PostNew) Validate(validate *validator.Validate) error {
	return validate.Struct(input)
}
