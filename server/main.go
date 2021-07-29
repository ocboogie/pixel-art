package main

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/ocboogie/pixel-art/api"
	"github.com/ocboogie/pixel-art/postgres"
	"github.com/ocboogie/pixel-art/services/art"
	"github.com/ocboogie/pixel-art/services/auth"
	"github.com/ocboogie/pixel-art/services/avatar"
	"github.com/ocboogie/pixel-art/services/post"
	"github.com/ocboogie/pixel-art/services/user"
	"github.com/sirupsen/logrus"
)

func main() {
	db, err := sqlx.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=password sslmode=disable")
	if err != nil {
		panic(err)
	}
	sb := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	defer db.Close()

	validate := validator.New()
	log := logrus.New()
	userRepo := postgres.NewRepositoryUser(db)
	postRepo := postgres.NewPostRepository(db, sb)
	likeRepo := postgres.NewLikeRepository(db)
	sessionRepo := postgres.NewRepositorySession(db)

	avatar := avatar.New(avatar.DefaultConfig())
	auth := auth.New(log, auth.DefaultConfig(), userRepo, sessionRepo, avatar)
	art := art.New(art.DefaultConfig())
	post := post.New(log, userRepo, postRepo, likeRepo)
	user := user.New(log, userRepo)

	server := api.New(auth, avatar, art, post, user, validate)

	server.Setup()
	server.Start()

	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()
	// migrations.Migrate(db)
}
