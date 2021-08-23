package auth

import (
	"github.com/ocboogie/pixel-art/pkg/argon2"
)

type Config struct {
	HashConfig *argon2.Params

	SessionLifetime uint
	Secret          string
}
