package repositories

import (
	"errors"
	"time"

	"github.com/ocboogie/pixel-art/models"
)

//go:generate mockgen -destination=../mocks/repo_post.go -package mocks -mock_names Post=RepositoryPost github.com/ocboogie/pixel-art/repositories Post

var (
	ErrPostNotFound = errors.New("Post not found")
)

type Post interface {
	Find(id string) (*models.Post, error)
	Save(post *models.Post) error
	Latest(limit int, after *time.Time) ([]*models.Post, error)
}
