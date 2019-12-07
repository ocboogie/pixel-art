package post

import (
	"time"

	"github.com/google/uuid"
	"github.com/ocboogie/pixel-art/config"
	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/repositories"
)

//go:generate mockgen -destination=../../mocks/service_post.go -package mocks -mock_names Service=ServicePost github.com/ocboogie/pixel-art/services/post Service

type Service interface {
	Create(input models.PostNew) (string, error)
	Latest(limit int, after *time.Time) ([]*models.Post, error)
	Find(id string) (*models.Post, error)
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

func (s *service) Create(input models.PostNew) (string, error) {
	id := uuid.New().String()

	post := &models.Post{
		ID: id,
		Author: models.User{
			ID: input.UserID,
		},
		Title: input.Title,
		// FIXME: Decode from the input
		Data:      []byte{},
		CreatedAt: time.Now(),
	}

	if err := s.postRepo.Save(post); err != nil {
		return "", err
	}

	return id, nil
}

func (s *service) Find(id string) (*models.Post, error) {
	return s.postRepo.Find(id)
}

func (s *service) Latest(limit int, after *time.Time) ([]*models.Post, error) {
	return s.postRepo.Latest(limit, after)
}
