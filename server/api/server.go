package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/ocboogie/pixel-art/services/art"
	"github.com/ocboogie/pixel-art/services/auth"
	"github.com/ocboogie/pixel-art/services/avatar"
	"github.com/ocboogie/pixel-art/services/post"
	"github.com/ocboogie/pixel-art/services/user"
	"gopkg.in/go-playground/validator.v9"
)

type server struct {
	router   *chi.Mux
	validate *validator.Validate
	auth     auth.Service
	avatar   avatar.Service
	art      art.Service
	post     post.Service
	user     user.Service
}

func New(auth auth.Service,
	avatar avatar.Service,
	art art.Service,
	post post.Service,
	user user.Service,
	validate *validator.Validate) *server {

	s := &server{
		validate: validate,
		auth:     auth,
		avatar:   avatar,
		art:      art,
		post:     post,
		user:     user,
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
		AllowedOrigins:   []string{"http://localhost:8080"},
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
