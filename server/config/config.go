package config

import (
	"os"
	"strconv"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/pkg/argon2"
	"github.com/ocboogie/pixel-art/services/auth"
	log "github.com/sirupsen/logrus"
)

func getEnvVar(name string) string {
	envVar, exists := os.LookupEnv(name)
	if !exists {
		log.Fatalf("You must supply a \"%v\" environment variable (see .env.example)", name)
	}

	return envVar
}

func getCommaSepEnvVar(name string) []string {
	string := getEnvVar(name)
	array := strings.Split(string, ",")
	for i := range array {
		array[i] = strings.TrimSpace(array[i])
	}

	return array
}

// NewAddr creates an address in the form ":{PORT}" where {PORT} is the
// environment variable "PORT"
func NewAddr() string {
	port := getEnvVar("PORT")

	return ":" + port
}

// NewCorsOptions creates cors options from "CORS_ALLOWED_ORIGINS", "CORS_ALLOWED_METHODS", "CORS_ALLOW_CREDENTIALS", and "CORS_MAX_AGE"
func NewCorsOptions() cors.Options {
	maxAge := getEnvVar("CORS_MAX_AGE")
	maxAgeInt, err := strconv.ParseInt(maxAge, 10, 32)
	if err != nil {
		log.Fatalf("\"CORS_MAX_AGE\" must be a number: %v", err)
	}

	return cors.Options{
		AllowedOrigins:   getCommaSepEnvVar("CORS_ALLOWED_ORIGINS"),
		AllowedMethods:   getCommaSepEnvVar("CORS_ALLOWED_METHODS"),
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: getEnvVar("CORS_ALLOW_CREDENTIALS") == "true",
		MaxAge:           int(maxAgeInt),
	}
}

// NewPGDB creates the database pool and squirrel statement builder. The database
// is found through the PG_DATABASE environment variable. Don't forget to
// `defer db.Close()`
func NewPGDB() (*sqlx.DB, sq.StatementBuilderType) {
	database, exists := os.LookupEnv("PG_DATABASE")
	if !exists {
		log.Fatal("You must supply a \"PG_DATABASE\" environment variable (see .env.example)")
	}

	db, err := sqlx.Open("postgres", database)
	if err != nil {
		panic(err)
	}

	sb := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	return db, sb
}

// NewAuthConfig creates a new auth config from the environment variables
// "SESSION_LIFETIME" and "SECRET".
func NewAuthConfig() auth.Config {
	sessionLifetime := getEnvVar("SESSION_LIFETIME")
	sessionLifetimeInt, err := strconv.ParseInt(sessionLifetime, 10, 32)
	if err != nil {
		log.Fatalf("\"SESSION_LIFETIME\" must be a number: %v", err)
	}

	return auth.Config{
		// TODO: Get this through environment variables
		HashConfig: argon2.DefaultParams(),

		SessionLifetime: uint(sessionLifetimeInt),
		Secret:          getEnvVar("SECRET"),
	}
}

// NewAvatarSpec creates a new avatar spec from the "AVATAR_SIZE" and
// "AVATAR_PALETTE" environment variables. This will panic if they are not
// found.
func NewAvatarSpec() models.AvatarSpec {
	size := getEnvVar("AVATAR_SIZE")
	sizeInt, err := strconv.ParseInt(size, 10, 8)
	if err != nil {
		log.Fatalf("\"AVATAR_SIZE\" must be a number: %v", err)
	}

	palette := getCommaSepEnvVar("AVATAR_PALETTE")

	return models.AvatarSpec{
		Size:    uint8(sizeInt),
		Palette: palette,
	}
}

// NewArtSpec creates a new avatar spec from the "ART_SIZE" and
// "ART_COLORS" environment variables. This will panic if they are not found.
func NewArtSpec() models.ArtSpec {
	size := getEnvVar("ART_SIZE")
	sizeInt, err := strconv.ParseInt(size, 10, 16)
	if err != nil {
		log.Fatalf("\"ART_SIZE\" must be a number: %v", err)
	}

	colors := getEnvVar("ART_COLORS")
	colorsInt, err := strconv.ParseInt(colors, 10, 8)
	if err != nil {
		log.Fatalf("\"ART_COLORS\" must be a number: %v", err)
	}

	return models.ArtSpec{
		Size:   uint16(sizeInt),
		Colors: uint8(colorsInt),
	}
}
