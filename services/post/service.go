package posting

import (
	"time"

	"github.com/google/uuid"
	"github.com/ocboogie/pixel-art/config"
	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/repositories"
)

//go:generate mockgen -destination=../../mocks/service_posting.go -package mocks -mock_names Service=ServicePosting github.com/ocboogie/pixel-art/services/posting Service

type Service interface {
	Create(input models.PostInput) (string, error)
}

type service struct {
	userRepo repositories.User
	postRepo repositories.Post
	config   *config.Config
}

func New(config *config.Config, userRepo repositories.User, postRepo repositories.Post) Service {
	return &service{
		userRepo: userRepo,
		postRepo: postRepo,
		config:   config,
	}
}

func (s *service) Create(input models.PostInput) (string, error) {
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

	if err := s.postRepo.Create(post); err != nil {
		return "", err
	}

	return id, nil
}
