package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ocboogie/pixel-art/config"
	"github.com/ocboogie/pixel-art/services/auth"
	"github.com/ocboogie/pixel-art/services/post"
	"github.com/ocboogie/pixel-art/services/user"
	"github.com/rs/cors"
	"gopkg.in/go-playground/validator.v9"
)

type server struct {
	config   *config.Config
	router   *mux.Router
	validate *validator.Validate
	auth     auth.Service
	post     post.Service
	user     user.Service
}

func New(config *config.Config,
	auth auth.Service,
	post post.Service,
	user user.Service,
	validate *validator.Validate) *server {

	s := &server{
		config:   config,
		validate: validate,
		auth:     auth,
		post:     post,
		user:     user,
	}

	return s
}

func (s *server) Setup() {
	s.router = mux.NewRouter()

	s.routes()
}

func (s *server) Start() {
	if s.router == nil {
		panic("Shouldn't start before setting up the server")
	}

	handler := cors.Default().Handler(s.router)
	http.ListenAndServe(":8000", handler)
}
