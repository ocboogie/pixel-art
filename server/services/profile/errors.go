package profile

import (
	"github.com/ocboogie/pixel-art/repositories"
)

var (
	ErrNotFound            = repositories.ErrUserNotFound
	ErrFollowAlreadyExists = repositories.ErrFollowAlreadyExists
	ErrFollowNotFound      = repositories.ErrFollowNotFound
	ErrFollowSelf          = repositories.ErrFollowSelf
)
