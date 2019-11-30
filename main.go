package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/ocboogie/pixel-art/api"
	"github.com/ocboogie/pixel-art/config"
	"github.com/ocboogie/pixel-art/postgres"
	"github.com/ocboogie/pixel-art/services/auth"
	"github.com/ocboogie/pixel-art/services/post"
	"github.com/ocboogie/pixel-art/services/user"
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	db, err := sqlx.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=password sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	config := config.GetConfig()

	validate := validator.New()
	userRepo := postgres.NewRepositoryUser(db)
	postRepo := postgres.NewPostRepository(db)
	sessionRepo := postgres.NewRepositorySession(db)

	auth := auth.New(&config, userRepo, sessionRepo)
	post := post.New(&config, userRepo, postRepo)
	user := user.New(&config, userRepo)

	server := api.New(&config, auth, post, user, validate)

	server.Setup()
	server.Start()

	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()
	// migrations.Migrate(db)
}
