package posting

import (
	"time"

	"github.com/google/uuid"
	"github.com/ocboogie/pixel-art/config"
	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/repositories"
)

type Service struct {
	UserRepo repositories.User
	PostRepo repositories.Post
	Config   *config.Config
}

func (s *Service) Post(input models.PostInput) (string, error) {
	if err := input.Validate(); err != nil {
		return "", &ErrInvalidPost{Err: err}
	}

	idBytes, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	id := idBytes.String()

	post := &models.Post{
		ID:        id,
		UserID:    input.UserID,
		Title:     input.Title,
		Data:      input.Data,
		CreatedAt: time.Now(),
	}

	if err := s.PostRepo.Create(post); err != nil {
		return "", err
	}

	return id, nil
}
