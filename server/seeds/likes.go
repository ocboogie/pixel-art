package seeds

import (
	"fmt"
	"math/rand"

	"github.com/fatih/color"
	"github.com/ocboogie/pixel-art/repositories"
)

func (s seed) C_LikesSeed() {
	users, err := s.userRepo.All(repositories.UserIncludes{})
	if err != nil {
		panic(err)
	}

	posts, err := s.postRepo.Latest(50, nil, repositories.PostIncludes{})
	if err != nil {
		panic(err)
	}

	for i := 0; i < 50; i++ {

		user := users[rand.Intn(len(users))].ID
		post := posts[rand.Intn(len(posts))].ID

		if err := s.likeRepo.Save(user, post); err != nil && err != repositories.ErrLikeAlreadyExists {
			panic(err)
		}

		if s.verbose {
			fmt.Printf("%v liked %v\n", user, post)
		} else {
			color.New(color.FgMagenta).Printf("\033[2K\rCreating likes %d/50", i+1)
		}
	}
}
