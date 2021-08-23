package repositories

import "errors"

var (
	ErrFollowAlreadyExists = errors.New("Follow already exists")
	ErrFollowNotFound      = errors.New("Follow not found")
	ErrFollowSelf          = errors.New("Mustn't follow self")
	ErrLikeAlreadyExists   = errors.New("Like already exists")
	ErrLikeNotFound        = errors.New("Like not found")
	ErrPostNotFound        = errors.New("Post not found")
	ErrSessionNotFound     = errors.New("Session not found")
	ErrUserNotFound        = errors.New("User not found")
)
