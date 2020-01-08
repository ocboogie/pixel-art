package post

import (
	"encoding/base64"
	"time"

	"github.com/google/uuid"
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
}

func New(userRepo repositories.User, postRepo repositories.Post) Service {
	return &service{
		userRepo: userRepo,
		postRepo: postRepo,
	}
}

func (s *service) Create(input models.PostNew) (string, error) {
	id := uuid.New().String()

	data, err := base64.StdEncoding.DecodeString(input.Data)

	if err != nil {
		return "", err
	}

	post := &models.Post{
		ID: id,
		Author: models.User{
			ID: input.UserID,
		},
		Title:     input.Title,
		Data:      data,
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
