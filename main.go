package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/ocboogie/pixel-art/api"
	"github.com/ocboogie/pixel-art/config"
	"github.com/ocboogie/pixel-art/postgres"
	"github.com/ocboogie/pixel-art/services/authenticating"
	"github.com/ocboogie/pixel-art/services/post"
)

func main() {
	db, err := sqlx.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=password sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	config := config.GetConfig()

	userRepo := postgres.NewRepositoryUser(db)
	postRepo := postgres.NewPostRepository(db)
	sessionRepo := postgres.NewRepositorySession(db)

	authenticating := authenticating.New(&config, userRepo, sessionRepo)
	post := post.New(&config, userRepo, postRepo)

	server := api.New(&config, authenticating, post)

	server.Setup()
	server.Start()

	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()
	// migrations.Migrate(db)
}
