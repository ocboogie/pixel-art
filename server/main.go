package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/ocboogie/pixel-art/api"
	"github.com/ocboogie/pixel-art/config"
	"github.com/ocboogie/pixel-art/postgres"
	"github.com/ocboogie/pixel-art/services/auth"
	"github.com/ocboogie/pixel-art/services/feed"
	"github.com/ocboogie/pixel-art/services/profile"
)

func main() {
	db, sb := config.NewPGDB()
	defer db.Close()

	artSpec := config.NewArtSpec()
	avatarSpec := config.NewAvatarSpec()

	validate := validator.New()
	userRepo := postgres.NewRepositoryUser(db, sb)
	postRepo := postgres.NewPostRepository(db, sb)
	likeRepo := postgres.NewLikeRepository(db)
	followRepo := postgres.NewFollowRepository(db)
	sessionRepo := postgres.NewRepositorySession(db)

	authConfig := config.NewAuthConfig()
	auth := auth.New(authConfig, userRepo, sessionRepo)
	feed := feed.New(userRepo, postRepo, likeRepo)
	profile := profile.New(userRepo, followRepo)

	addr := config.NewAddr()
	corsOptions := config.NewCorsOptions()
	server := api.New(addr, corsOptions, auth, avatarSpec, artSpec, feed, profile, validate)

	server.Setup()
	server.Start()
}
