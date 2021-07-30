package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Post struct {
	ID        string    `json:"id"`
	Author    User      `json:"author" db:"author"`
	Title     string    `json:"title"`
	Art       Art       `json:"art"`
	Likes     int       `json:"likes"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type PostNew struct {
	UserID string `json:"userId" validate:"required,uuid"`
	Title  string `json:"title" validate:"required,min=2,max=256"`
	Art    Art    `json:"art"`
}

func (input PostNew) Validate(validate *validator.Validate, artSpec ArtSpec) error {
	if err := input.Art.Validate(artSpec); err != nil {
		return err
	}

	return validate.Struct(input)
}
