package repositories

import (
	"time"

	"github.com/ocboogie/pixel-art/models"
)

//go:generate mockgen -destination=../mocks/repo_post.go -package mocks -mock_names Post=RepositoryPost github.com/ocboogie/pixel-art/repositories Post

type Post interface {
	Find(id string) (*models.Post, error)
	Save(post *models.Post) error
	Latest(limit int, after *time.Time) ([]*models.Post, error)
	PostsByUser(userID string, limit int, after *time.Time) ([]*models.Post, error)
}
