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

type PostInput struct {
	UserID string `json:"userID" validate:"required,uuid"`
	Title  string `json:"title" validate:"required,min=2,max=256"`
	// TODO: validate Data
	Data []byte `json:"data"`
}

func (input PostInput) Validate() error {
	// TODO: Probably shouldn't create a new validator everytime
	validate := validator.New()

	return validate.Struct(input)
}
