package repositories

import (
	"errors"

	"github.com/ocboogie/pixel-art/models"
)

//go:generate mockgen -destination=../mocks/session_repo.go -package mocks -mock_names Session=SessionRepository github.com/ocboogie/pixel-art/repositories Session

var (
	ErrSessionNotFound = errors.New("Session not found")
)

type Session interface {
	Create(session *models.Session) error
	Find(id string) (*models.Session, error)
	Delete(id string) error
}
