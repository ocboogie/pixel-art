package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Post struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Art       ArtEncoded `json:"art"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	AuthorID  string     `json:"authorId" db:"author_id"`
	// Virtual fields, meaning that they don't actually exist in the
	// database, but are infered from other factors
	Author *PublicUser `json:"author" db:"author"`
	Likes  *int        `json:"likes"`
	Liked  *bool       `json:"liked"`
}

type PostNew struct {
	UserID string     `json:"userId" validate:"required,uuid"`
	Title  string     `json:"title" validate:"required,min=2,max=256"`
	Art    ArtEncoded `json:"art"`
}

func (input PostNew) Validate(validate *validator.Validate, artSpec ArtSpec) error {
	if _, err := input.Art.Decode(artSpec); err != nil {
		return err
	}

	return validate.Struct(input)
}
