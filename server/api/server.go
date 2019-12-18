package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ocboogie/pixel-art/services/auth"
	"github.com/ocboogie/pixel-art/services/avatar"
	"github.com/ocboogie/pixel-art/services/post"
	"github.com/ocboogie/pixel-art/services/user"
	"github.com/rs/cors"
	"gopkg.in/go-playground/validator.v9"
)

type server struct {
	router   *chi.Mux
	validate *validator.Validate
	auth     auth.Service
	avatar   avatar.Service
	post     post.Service
	user     user.Service
}

func New(auth auth.Service,
	avatar avatar.Service,
	post post.Service,
	user user.Service,
	validate *validator.Validate) *server {

	s := &server{
		validate: validate,
		auth:     auth,
		avatar:   avatar,
		post:     post,
		user:     user,
	}

	return s
}

func (s *server) Setup() {
	s.router = chi.NewRouter()

	s.routes()
}

func (s *server) Start() {
	if s.router == nil {
		panic("Shouldn't start before setting up the server")
	}

	handler := cors.Default().Handler(s.router)
	http.ListenAndServe(":8000", handler)
}
