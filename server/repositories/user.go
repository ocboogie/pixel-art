package repositories

import (
	"github.com/ocboogie/pixel-art/models"
)

//go:generate mockgen -destination=../mocks/repo_user.go -package mocks -mock_names User=RepositoryUser github.com/ocboogie/pixel-art/repositories User

// PostIncludes exists to denote which virtual properties should be computed
type UserIncludes struct {
	// Whether to include if the user is being followed by the user with this ID
	// "" means don't include
	Following string
}

type User interface {
	Find(id string, includes UserIncludes) (*models.User, error)
	FindByEmail(email string, includes UserIncludes) (*models.User, error)
	Update(user *models.User) error
	Save(user *models.User) error
	ExistsEmail(email string) (bool, error)
}
