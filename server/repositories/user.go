package repositories

import (
	"errors"

	"github.com/ocboogie/pixel-art/models"
)

//go:generate mockgen -destination=../mocks/repo_user.go -package mocks -mock_names User=RepositoryUser github.com/ocboogie/pixel-art/repositories User

var (
	ErrUserNotFound = errors.New("User not found")
)

type User interface {
	Find(id string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	// Update(user *models.User) error
	Save(user *models.User) error
	ExistsEmail(email string) (bool, error)
}
