package main

import (
	"flag"

	_ "github.com/joho/godotenv/autoload"
	"github.com/ocboogie/pixel-art/config"
	"github.com/ocboogie/pixel-art/postgres"
	"github.com/ocboogie/pixel-art/seeds"
	"github.com/ocboogie/pixel-art/services/auth"
	log "github.com/sirupsen/logrus"
)

var verbose bool

func init() {
	flag.BoolVar(&verbose, "v", false, "")
	flag.BoolVar(&verbose, "verbose", false, "Whether to show the details of the seeds")
}

func main() {
	flag.Parse()

	db, sb := config.NewPGDB()
	defer db.Close()

	log.SetLevel(log.WarnLevel)

	artSpec := config.NewArtSpec()
	avatarSpec := config.NewAvatarSpec()

	userRepo := postgres.NewRepositoryUser(db, sb)
	postRepo := postgres.NewPostRepository(db, sb)
	likeRepo := postgres.NewLikeRepository(db)
	followRepo := postgres.NewFollowRepository(db)
	sessionRepo := postgres.NewRepositorySession(db)

	authConfig := config.NewAuthConfig()
	auth := auth.New(authConfig, userRepo, sessionRepo)

	seeds.Seed(verbose, db, userRepo, postRepo, likeRepo, followRepo, auth, avatarSpec, artSpec)
}
