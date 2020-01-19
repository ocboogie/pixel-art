package repositories

import "errors"

var (
	ErrLikeAlreadyExists = errors.New("Like already exists")
	ErrLikeNotFound      = errors.New("Like not found")
	ErrPostNotFound      = errors.New("Post not found")
	ErrSessionNotFound   = errors.New("Session not found")
	ErrUserNotFound      = errors.New("User not found")
)
