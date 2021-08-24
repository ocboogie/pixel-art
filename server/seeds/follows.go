package seeds

import (
	"fmt"
	"math/rand"

	"github.com/fatih/color"
	"github.com/ocboogie/pixel-art/repositories"
)

func (s seed) B_FollowsSeed() {
	users, err := s.userRepo.All(repositories.UserIncludes{})
	if err != nil {
		panic(err)
	}

	for i := 0; i < 50; i++ {
		u1 := users[rand.Intn(len(users))].ID
		u2 := users[rand.Intn(len(users))].ID

		if err := s.followRepo.Save(u1, u2); err != nil && err != repositories.ErrFollowAlreadyExists && err != repositories.ErrFollowSelf {
			panic(err)
		}

		if s.verbose {
			fmt.Printf("%v followed %v\n", u2, u1)
		} else {
			color.New(color.FgMagenta).Printf("\033[2K\rCreating follows %d/50", i+1)
		}
	}
}
