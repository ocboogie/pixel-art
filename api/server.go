package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ocboogie/pixel-art/config"
	"github.com/ocboogie/pixel-art/services/authenticating"
	"github.com/ocboogie/pixel-art/services/post"
)

type server struct {
	e              *echo.Echo
	config         *config.Config
	authenticating authenticating.Service
	post           post.Service
}

func New(config *config.Config,
	authenticating authenticating.Service,
	post post.Service) *server {

	s := &server{
		authenticating: authenticating,
		post:           post,
		config:         config,
	}

	return s
}

type ResponsePayload struct {
	Data interface{} `json:"data,omitempty"`
	Err  string      `json:"error,omitempty"`
}

func (s *server) Setup() {
	s.e = echo.New()

	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())

	s.e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		msg := "Internal server error"

		if he, ok := err.(*echo.HTTPError); ok {
			msg = he.Message.(string)
			code = he.Code
		}

		c.JSON(code, ResponsePayload{
			Err: msg,
		})

	}

	s.routes()
}

func (s *server) Start() {
	if s.e == nil {
		panic("Shouldn't start before setting up the server")
	}

	s.e.Start(":4000")
}
