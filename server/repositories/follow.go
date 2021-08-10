package repositories

//go:generate mockgen -destination=../mocks/repo_follow.go -package mocks -mock_names Follow=RepositoryFollow github.com/ocboogie/pixel-art/repositories Follow

type Follow interface {
	Save(followedID string, followerID string) error
	Delete(followedID string, followerID string) error
}
