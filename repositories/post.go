package repositories

import (
	"errors"

	"github.com/ocboogie/pixel-art/models"
)

//go:generate mockgen -destination=../mocks/repo_post.go -package mocks -mock_names Post=RepositoryPost github.com/ocboogie/pixel-art/repositories Post

var (
	ErrPostNotFound = errors.New("Post not found")
)

type Post interface {
	Find(id string) (*models.Post, error)
	Save(post *models.Post) error
	Latest(limit int) ([]*models.Post, error)
}
