package repositories

//go:generate mockgen -destination=../mocks/repo_like.go -package mocks -mock_names Like=RepositoryLike github.com/ocboogie/pixel-art/repositories Like

type Like interface {
	Save(userID string, postID string) error
	Delete(userID string, postID string) error
}
