package feed

import (
	"time"

	"github.com/google/uuid"
	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/repositories"
	log "github.com/sirupsen/logrus"
)

//go:generate mockgen -destination=../../mocks/service_feed.go -package mocks -mock_names Service=ServiceFeed github.com/ocboogie/pixel-art/services/feed Service

type PostIncludes = repositories.PostIncludes

type Service interface {
	Create(input models.PostNew) (string, error)
	Delete(id string) error
	Latest(limit int, after *time.Time, includes PostIncludes) ([]*models.Post, error)
	Find(id string, includes PostIncludes) (*models.Post, error)
	PostsByUser(id string, limit int, after *time.Time, includes PostIncludes) ([]*models.Post, error)
	Like(userID string, postID string) error
	Unlike(userID string, postID string) error
}

type service struct {
	userRepo repositories.User
	postRepo repositories.Post
	likeRepo repositories.Like
}

func New(userRepo repositories.User, postRepo repositories.Post, likeRepo repositories.Like) Service {
	return &service{
		userRepo: userRepo,
		postRepo: postRepo,
		likeRepo: likeRepo,
	}
}

func (s *service) Create(input models.PostNew) (string, error) {
	id := uuid.New().String()

	post := &models.Post{
		ID:        id,
		AuthorID:  input.UserID,
		Title:     input.Title,
		Art:       input.Art,
		CreatedAt: time.Now(),
	}

	if err := s.postRepo.Save(post); err != nil {
		return "", err
	}

	log.WithFields(log.Fields{
		"postID": id,
		"title":  input.Title,
		"userID": input.UserID,
	}).Info("Post created")

	return id, nil
}

func (s *service) Delete(id string) error {
	return s.postRepo.Delete(id)
}

func (s *service) Find(id string, includes PostIncludes) (*models.Post, error) {
	return s.postRepo.Find(id, includes)
}

func (s *service) Latest(limit int, after *time.Time, includes PostIncludes) ([]*models.Post, error) {
	return s.postRepo.Latest(limit, after, includes)
}

func (s *service) PostsByUser(userID string, limit int, after *time.Time, includes PostIncludes) ([]*models.Post, error) {
	return s.postRepo.PostsByUser(userID, limit, after, includes)
}

func (s *service) Like(userID string, postID string) error {
	if err := s.likeRepo.Save(userID, postID); err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"postID": postID,
		"userID": userID,
	}).Info("Post liked")

	return nil
}

func (s *service) Unlike(userID string, postID string) error {
	if err := s.likeRepo.Delete(userID, postID); err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"postID": postID,
		"userID": userID,
	}).Info("Post unliked")

	return nil
}
