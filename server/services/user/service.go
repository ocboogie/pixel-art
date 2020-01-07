package user

import (
	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/repositories"
)

//go:generate mockgen -destination=../../mocks/service_user.go -package mocks -mock_names Service=ServiceUser github.com/ocboogie/pixel-art/services/user Service

type Service interface {
	Find(id string) (*models.User, error)
	Update(user *models.User) error
}

type service struct {
	userRepo repositories.User
}

func New(userRepo repositories.User) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (s *service) Find(id string) (*models.User, error) {
	return s.userRepo.Find(id)
}

// Update expects user to be in a valid user state
func (s *service) Update(user *models.User) error {
	return s.userRepo.Update(user)
}
