package main

import (
	"log"
	"os"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/ocboogie/pixel-art/api"
	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/postgres"
	"github.com/ocboogie/pixel-art/services/auth"
	"github.com/ocboogie/pixel-art/services/feed"
	"github.com/ocboogie/pixel-art/services/profile"
)

var avatarSpec = models.AvatarSpec{
	Size: 5,
	Palette: []string{
		"#1abc9c", "#2ecc71", "#3498db", "#9b59b6", "#e74c3c",
	},
}
var artSpec = models.ArtSpec{
	Size:   3,
	Colors: 3,
	// Size:   25,
	// Colors: 8,
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	database, exists := os.LookupEnv("DATABASE")
	if !exists {
		log.Fatal("You must supply a \"DATABASE\" environment variable (see .env.example)")
	}

	db, err := sqlx.Open("postgres", database)
	if err != nil {
		panic(err)
	}
	sb := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	defer db.Close()

	validate := validator.New()
	userRepo := postgres.NewRepositoryUser(db, sb)
	postRepo := postgres.NewPostRepository(db, sb)
	likeRepo := postgres.NewLikeRepository(db)
	followRepo := postgres.NewFollowRepository(db)
	sessionRepo := postgres.NewRepositorySession(db)

	auth := auth.New(auth.DefaultConfig(), userRepo, sessionRepo)
	feed := feed.New(userRepo, postRepo, likeRepo)
	profile := profile.New(userRepo, followRepo)

	server := api.New(auth, avatarSpec, artSpec, feed, profile, validate)

	server.Setup()
	server.Start()

	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()
	// migrations.Migrate(db.DB)
}
