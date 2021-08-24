package seeds

import (
	"fmt"
	"reflect"

	"github.com/fatih/color"
	"github.com/jmoiron/sqlx"
	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/repositories"
	"github.com/ocboogie/pixel-art/services/auth"
)

type seed struct {
	verbose    bool
	db         *sqlx.DB
	userRepo   repositories.User
	postRepo   repositories.Post
	likeRepo   repositories.Like
	followRepo repositories.Follow
	auth       auth.Service
	avatarSpec models.AvatarSpec
	artSpec    models.ArtSpec
}

func Seed(
	verbose bool,
	db *sqlx.DB,
	userRepo repositories.User,
	postRepo repositories.Post,
	likeRepo repositories.Like,
	followRepo repositories.Follow,
	auth auth.Service,
	avatarSpec models.AvatarSpec,
	artSpec models.ArtSpec,
) {
	s := seed{
		verbose:    verbose,
		db:         db,
		userRepo:   userRepo,
		postRepo:   postRepo,
		likeRepo:   likeRepo,
		followRepo: followRepo,
		auth:       auth,
		avatarSpec: avatarSpec,
		artSpec:    artSpec,
	}

	seedType := reflect.TypeOf(s)

	color.New(color.Bold, color.Underline, color.FgYellow).Println("Running all seeds...\n")
	// We are looping over the method on a Seed struct
	for i := 0; i < seedType.NumMethod(); i++ {
		// Get the method in the current iteration
		method := seedType.Method(i)
		// Execute seeder
		seedMethod(s, method.Name)
	}
}

// Execute will executes the given seeder method
// func Execute(db *sql.DB, seedMethodNames ...string) {
// 	s := Seed{db}
//
// 	seedType := reflect.TypeOf(s)
//
// 	// Execute all seeders if no method name is given
// 	if len(seedMethodNames) == 0 {
// 		log.Println("Running all seeder...")
// 		// We are looping over the method on a Seed struct
// 		for i := 0; i < seedType.NumMethod(); i++ {
// 			// Get the method in the current iteration
// 			method := seedType.Method(i)
// 			// Execute seeder
// 			seed(s, method.Name)
// 		}
// 	}
//
// 	// Execute only the given method names
// 	// for _, item := range seedMethodNames {
// 	// 	seed(s, item)
// 	// }
// }

func seedMethod(s seed, seedMethodName string) {
	// Get the reflect value of the method
	m := reflect.ValueOf(s).MethodByName(seedMethodName)
	// Exit if the method doesn't exist
	if !m.IsValid() {
		panic(fmt.Sprintln("No method called", seedMethodName))
	}
	// Execute the method
	color.New(color.Bold, color.FgGreen).Println("Seeding", seedMethodName)
	m.Call(nil)
	if !s.verbose {
		fmt.Print("\n")
	}
	color.New(color.Bold, color.FgGreen).Println(seedMethodName, "Finished\n")
}
