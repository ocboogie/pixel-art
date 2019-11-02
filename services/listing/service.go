package listing

import (
	"github.com/ocboogie/pixel-art/config"
	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/repositories"
)

//go:generate mockgen -destination=../../mocks/service_listing.go -package mocks -mock_names Service=ServiceListing github.com/ocboogie/pixel-art/services/listing Service

type Service interface {
	Latest() ([]*models.Post, error)
}

type service struct {
	postRepo repositories.Post
	config   *config.Config
}

func New(config *config.Config, postRepo repositories.Post) Service {
	return &service{
		postRepo: postRepo,
		config:   config,
	}
}

func (s *service) Latest() ([]*models.Post, error) {
	return s.postRepo.Latest(20)
}
