package repositories

import (
	"errors"

	"github.com/ocboogie/pixel-art/models"
)

//go:generate mockgen -destination=../mocks/post_repo.go -package mocks -mock_names Post=PostRepository github.com/ocboogie/pixel-art/repositories Post

var (
	ErrPostNotFound = errors.New("Post not found")
)

type Post interface {
	Find(id string) (*models.Post, error)
	Create(post *models.Post) error
	Latest(limit uint) ([]*models.Post, error)
}
