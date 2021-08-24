package seeds

import (
	"fmt"
	"math/rand"
	"time"

	faker "github.com/bxcodec/faker/v3"
	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/repositories"
)

func (s seed) B_PostSeed() {
	users, err := s.userRepo.All(repositories.UserIncludes{})
	if err != nil {
		panic(err)
	}

	for i := 0; i < 50; i++ {
		id := uuid.New().String()

		art, err := models.GenerateRandomArt(s.artSpec)
		if err != nil {
			panic(err)
		}

		post := &models.Post{
			ID:        id,
			AuthorID:  users[rand.Intn(len(users))].ID,
			Title:     faker.Word() + " " + faker.Word(),
			Art:       art,
			CreatedAt: time.Unix(faker.UnixTime(), 0),
		}

		if err := s.postRepo.Save(post); err != nil {
			panic(err)
		}

		if s.verbose {
			fmt.Printf("Created post: %+v\n", post)
		} else {
			color.New(color.FgMagenta).Printf("\033[2K\rCreating posts %d/50", i+1)
		}
	}
}
