package seeds

import (
	"fmt"

	faker "github.com/bxcodec/faker/v3"
	"github.com/fatih/color"
	"github.com/ocboogie/pixel-art/models"
)

func (s seed) A_UserSeed() {
	for i := 0; i < 10; i++ {
		newUser := &models.UserNew{
			Avatar:   models.GenerateRandomAvatar(s.avatarSpec),
			Name:     faker.Username(),
			Email:    faker.Email(),
			Password: faker.Password(),
		}

		_, err := s.auth.SignUp(newUser)
		if err != nil {
			panic(err)
		}

		if s.verbose {
			fmt.Printf("Created user: %+v\n", newUser)
		} else {
			color.New(color.FgMagenta).Printf("\033[2K\rCreating users %d/10", i+1)
		}
	}
}
