package repositories

import (
	"time"

	"github.com/ocboogie/pixel-art/models"
)

//go:generate mockgen -destination=../mocks/repo_post.go -package mocks -mock_names Post=RepositoryPost github.com/ocboogie/pixel-art/repositories Post

// PostIncludes exists to denote which virtual properties should be computed
type PostIncludes struct {
	Author bool
	Likes  bool
	// Whether to include if the post is liked by the user with this ID.
	// "" means don't include
	Liked string
}

type Post interface {
	Find(id string, includes PostIncludes) (*models.Post, error)
	Delete(id string) error
	Save(post *models.Post) error
	Latest(limit int, after *time.Time, includes PostIncludes) ([]*models.Post, error)
	PostsByUser(userID string, limit int, after *time.Time, includes PostIncludes) ([]*models.Post, error)
	// Posts by users who the user with the id of userID follows
	Feed(userID string, limit int, after *time.Time, includes PostIncludes) ([]*models.Post, error)
}
