package authenticating

import "errors"

type ErrInvalidUser struct {
	Err error
}

func (e *ErrInvalidUser) Error() string { return "Invalid user: " + e.Err.Error() }
func (e *ErrInvalidUser) Unwrap() error { return e.Err }

var (
	ErrEmailAlreadyInUse  = errors.New("Email already in use")
	ErrInvalidCredentials = errors.New("Invalid credentials")
	ErrUserNotFound       = errors.New("User not found")
	ErrExpiredSession     = errors.New("Session expired")
)
