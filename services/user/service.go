package user

import (
	"github.com/ocboogie/pixel-art/config"
	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/repositories"
)

//go:generate mockgen -destination=../../mocks/service_user.go -package mocks -mock_names Service=ServiceUser github.com/ocboogie/pixel-art/services/user Service

type Service interface {
	Find(id string) (*models.User, error)
}

type service struct {
	userRepo repositories.User
	config   *config.Config
}

func New(config *config.Config, userRepo repositories.User) Service {
	return &service{
		userRepo: userRepo,
		config:   config,
	}
}

func (s *service) Find(id string) (*models.User, error) {
	return s.userRepo.Find(id)
}
