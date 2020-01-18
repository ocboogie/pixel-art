package repositories

import "errors"

//go:generate mockgen -destination=../mocks/repo_like.go -package mocks -mock_names Like=RepositoryLike github.com/ocboogie/pixel-art/repositories Like

var (
	ErrLikeAlreadyExists = errors.New("Like already exists")
	ErrLikeNotFound      = errors.New("Like not found")
)

type Like interface {
	Save(userID string, postID string) error
	Delete(userID string, postID string) error
}
