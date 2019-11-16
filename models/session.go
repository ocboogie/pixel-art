package models

import (
	"time"
)

type Session struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId" db:"user_id"`
	ExpiresAt time.Time `json:"expiresAt" db:"expires_at"`
}
