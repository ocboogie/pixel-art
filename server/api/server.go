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
	"github.com/sirupsen/logrus"
)

type server struct {
	router      *chi.Mux
	validate    *validator.Validate
	auth        auth.Service
	avatarSpec  models.AvatarSpec
	artSpec     models.ArtSpec
	feed        feed.Service
	profile     profile.Service
	addr        string
	corsOptions cors.Options
}

func New(addr string,
	corsOptios cors.Options,
	auth auth.Service,
	avatarSpec models.AvatarSpec,
	artSpec models.ArtSpec,
	feed feed.Service,
	profile profile.Service,
	validate *validator.Validate,
) *server {
	s := &server{
		validate:    validate,
		auth:        auth,
		avatarSpec:  avatarSpec,
		artSpec:     artSpec,
		feed:        feed,
		profile:     profile,
		addr:        addr,
		corsOptions: corsOptios,
	}

	return s
}

func (s *server) Setup() {
	s.router = chi.NewRouter()

	s.router.Use(middleware.RealIP)
	middleware.DefaultLogger = middleware.RequestLogger(&middleware.DefaultLogFormatter{
		Logger: logrus.New(),
	})
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(cors.Handler(s.corsOptions))

	s.routes()
}

func (s *server) Start() {
	if s.router == nil {
		panic("Shouldn't start before setting up the server")
	}

	logrus.Infof("🚀 Server started on %v", s.addr)
	logrus.Fatal(http.ListenAndServe(s.addr, s.router))
}
