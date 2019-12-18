package auth

import (
	"log"
	"os"

	"github.com/ocboogie/pixel-art/pkg/argon2"
)

type Config struct {
	HashConfig *argon2.Params

	SessionLifetime uint
	Secret          string
}

func DefaultConfig() Config {
	secret, exists := os.LookupEnv("SECRET")
	if !exists {
		log.Fatal("You must supply a \"SECRET\" environment variable")
	}

	return Config{
		HashConfig: argon2.DefaultParams(),

		SessionLifetime: 7 * 24 * 60 * 60 * 1000,
		Secret:          secret,
	}
}
