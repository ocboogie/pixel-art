package listing

import (
	"github.com/ocboogie/pixel-art/config"
	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/repositories"
)

type Service struct {
	PostRepo repositories.Post
	Config   *config.Config
}

func (s *Service) Latest() ([]*models.Post, error) {
	return s.PostRepo.Latest(20)
}
