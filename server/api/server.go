package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-playground/validator/v10"
	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/services/auth"
	"github.com/ocboogie/pixel-art/services/feed"
	"github.com/ocboogie/pixel-art/services/profile"
)

type server struct {
	router     *chi.Mux
	validate   *validator.Validate
	auth       auth.Service
	avatarSpec models.AvatarSpec
	artSpec    models.ArtSpec
	feed       feed.Service
	profile    profile.Service
}

func New(auth auth.Service,
	avatarSpec models.AvatarSpec,
	artSpec models.ArtSpec,
	feed feed.Service,
	profile profile.Service,
	validate *validator.Validate) *server {

	s := &server{
		validate:   validate,
		auth:       auth,
		avatarSpec: avatarSpec,
		artSpec:    artSpec,
		feed:       feed,
		profile:    profile,
	}

	return s
}

func (s *server) Setup() {
	s.router = chi.NewRouter()

	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(cors.Handler(cors.Options{
		// TODO: Make this configurable
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	s.routes()
}

func (s *server) Start() {
	if s.router == nil {
		panic("Shouldn't start before setting up the server")
	}

	http.ListenAndServe(":8000", s.router)
}
