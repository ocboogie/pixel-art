package profile

import (
	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/repositories"
	log "github.com/sirupsen/logrus"
)

//go:generate mockgen -destination=../../mocks/service_profile.go -package mocks -mock_names Service=ServiceProfile github.com/ocboogie/pixel-art/services/profile Service

type UserIncludes = repositories.UserIncludes

type Service interface {
	Find(id string, includes UserIncludes) (*models.User, error)
	Update(user *models.User) error
	Follow(followedID string, followerID string) error
	Unfollow(followedID string, followerID string) error
}

type service struct {
	userRepo    repositories.User
	followsRepo repositories.Follow
}

func New(userRepo repositories.User, followsRepo repositories.Follow) Service {
	return &service{
		userRepo:    userRepo,
		followsRepo: followsRepo,
	}
}

func (s *service) Find(id string, includes UserIncludes) (*models.User, error) {
	return s.userRepo.Find(id, includes)
}

// Update expects user to be in a valid user state
func (s *service) Update(user *models.User) error {
	if err := s.userRepo.Update(user); err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"username": user.Name,
		"userID":   user.ID,
	}).Info("User updated")

	return nil
}

func (s *service) Follow(followedID string, followerID string) error {
	if err := s.followsRepo.Save(followedID, followerID); err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"followedID": followedID,
		"followerID": followerID,
	}).Info("User followed")

	return nil
}

func (s *service) Unfollow(followedID string, followerID string) error {
	if err := s.followsRepo.Delete(followedID, followerID); err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"followedID": followedID,
		"followerID": followerID,
	}).Info("User unfollowed")

	return nil
}
