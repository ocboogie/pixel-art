package repositories

import (
	"github.com/ocboogie/pixel-art/models"
)

//go:generate mockgen -destination=../mocks/repo_session.go -package mocks -mock_names Session=RepositorySession github.com/ocboogie/pixel-art/repositories Session

type Session interface {
	Save(session *models.Session) error
	Find(id string) (*models.Session, error)
	Delete(id string) error
}
