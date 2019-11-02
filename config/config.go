package config

import (
	"log"
	"os"
	"strconv"

	"github.com/ocboogie/pixel-art/pkg/argon2"
)

type Config struct {
	HashConfig      *argon2.Params
	SessionLifetime uint
	Secret          string
}

func getEnvNumber(name string) (int, bool) {
	value, exists := os.LookupEnv(name)

	if !exists {
		return 0, false
	}

	num, err := strconv.Atoi(value)
	if err != nil {
		return 0, false
	}

	return num, true
}

func GetConfig() Config {
	secret, exists := os.LookupEnv("SECRET")
	if !exists {
		log.Fatal("You must supply a \"SECRET\" environment variable")
	}
	sessionLifetime := uint(7 * 24 * 60 * 60 * 1000)

	hashConfig := argon2.DefaultParams()

	if val, exists := getEnvNumber("MEMORY"); exists {
		hashConfig.Memory = uint32(val)
	}
	if val, exists := getEnvNumber("ITERATIONS"); exists {
		hashConfig.Iterations = uint32(val)
	}
	if val, exists := getEnvNumber("PARALLELISM"); exists {
		hashConfig.Parallelism = uint8(val)
	}
	if val, exists := getEnvNumber("SALT_LENGTH"); exists {
		hashConfig.SaltLength = uint32(val)
	}
	if val, exists := getEnvNumber("HASH_LENGTH"); exists {
		hashConfig.HashLength = uint32(val)
	}
	if val, exists := getEnvNumber("SESSION_LIFETIME"); exists {
		sessionLifetime = uint(val)
	}

	return Config{
		HashConfig:      hashConfig,
		SessionLifetime: sessionLifetime,
		Secret:          secret,
	}
}
