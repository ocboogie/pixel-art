package feed

import (
	"github.com/ocboogie/pixel-art/repositories"
)

var (
	ErrNotFound     = repositories.ErrPostNotFound
	ErrLikeNotFound = repositories.ErrLikeNotFound
	ErrUserNotFound = repositories.ErrUserNotFound
	ErrAlreadyLiked = repositories.ErrLikeAlreadyExists
)
