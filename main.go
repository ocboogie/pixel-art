package main

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/ocboogie/pixel-art/api"
	"github.com/ocboogie/pixel-art/config"
	"github.com/ocboogie/pixel-art/postgres"
	"github.com/ocboogie/pixel-art/services/authenticating"
	"github.com/ocboogie/pixel-art/services/listing"
	"github.com/ocboogie/pixel-art/services/posting"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=password sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	config := config.GetConfig()

	userRepo := postgres.NewUserRepository(db)
	postRepo := postgres.NewPostRepository(db)
	sessionRepo := postgres.NewSessionRepository(db)

	authenticating := &authenticating.Service{
		Config:      &config,
		UserRepo:    userRepo,
		SessionRepo: sessionRepo,
	}
	listing := &listing.Service{
		Config:   &config,
		PostRepo: postRepo,
	}
	posting := &posting.Service{
		Config:   &config,
		PostRepo: postRepo,
	}

	server := api.New(authenticating, listing, posting)

	server.Setup()
	server.Start()

	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()
	// migrations.Migrate(db)
}
