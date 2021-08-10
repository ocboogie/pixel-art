package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Avatar    Avatar    `json:"avatar"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	// Virtual fields, meaning that they don't actually exist in the
	// database, but are infered from other factors
	Following *bool `json:"following"`
}

func (u User) Hide() PublicUser {
	return PublicUser{
		ID:     u.ID,
		Name:   u.Name,
		Avatar: u.Avatar,
	}
}

// PublicUser is the same as User but with private fields hidden
type PublicUser struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Avatar Avatar `json:"avatar"`
	// Virtual fields, meaning that they don't actually exist in the
	// database, but are infered from other factors
	Following *bool `json:"following"`
}

// This is the data necessary to create a user
type UserNew struct {
	Name     string `json:"name" validate:"required,min=2,max=64"`
	Avatar   Avatar `json:"avatar" validate:"required,min=2,max=64"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=256"`
}

// This is the data necessary to login
type UserCredentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=256"`
}

func (input UserNew) Validate(validate *validator.Validate, avatarSpec AvatarSpec) error {
	if err := input.Avatar.Validate(avatarSpec); err != nil {
		return err
	}

	return validate.Struct(input)
}

func (input UserCredentials) Validate(validate *validator.Validate) error {
	return validate.Struct(input)
}
