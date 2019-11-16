package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ocboogie/pixel-art/config"
	"github.com/ocboogie/pixel-art/services/authenticating"
	"github.com/ocboogie/pixel-art/services/listing"
	"github.com/ocboogie/pixel-art/services/post"
)

type server struct {
	e              *echo.Echo
	config         *config.Config
	authenticating authenticating.Service
	listing        listing.Service
	post           post.Service
}

func New(config *config.Config,
	authenticating authenticating.Service,
	listing listing.Service,
	post post.Service) *server {

	s := &server{
		authenticating: authenticating,
		listing:        listing,
		post:           post,
		config:         config,
	}

	return s
}

type ResponsePayload struct {
	data interface{} `json:data,emitempty`
	err  string      `json:error,emitempty`
}

func (s *server) Setup() {
	s.e = echo.New()

	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())

	s.e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		msg := "Internal server error"

		if he, ok := err.(*echo.HTTPError); ok {
			msg = he.Error()
			code = he.Code
		}

		c.JSON(code, ResponsePayload{err: msg})

	}

	s.routes()
}

func (s *server) Start() {
	if s.e == nil {
		panic("Shouldn't start before setting up the server")
	}

	s.e.Start(":4000")
}
