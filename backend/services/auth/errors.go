package auth

import (
	"errors"

	"github.com/ocboogie/pixel-art/repositories"
)

var (
	ErrEmailAlreadyInUse  = errors.New("Email already in use")
	ErrInvalidCredentials = errors.New("Invalid credentials")
	ErrUserNotFound       = errors.New("User not found")
	ErrExpiredSession     = errors.New("Session expired")
	ErrSessionNotFound    = repositories.ErrSessionNotFound
)
