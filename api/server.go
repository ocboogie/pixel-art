package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ocboogie/pixel-art/config"
	"github.com/ocboogie/pixel-art/services/authenticating"
	"github.com/ocboogie/pixel-art/services/listing"
	"github.com/ocboogie/pixel-art/services/posting"
)

type server struct {
	e              *echo.Echo
	config         *config.Config
	authenticating authenticating.Service
	listing        listing.Service
	posting        posting.Service
}

type HTTPError interface {
	error

	Status() int
}

func New(config *config.Config,
	authenticating authenticating.Service,
	listing listing.Service,
	posting posting.Service) *server {

	s := &server{
		authenticating: authenticating,
		listing:        listing,
		posting:        posting,
		config:         config,
	}

	return s
}

type ErrorResponse struct {
	err string `json:error`
}

func (s *server) Setup() {
	s.e = echo.New()

	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())

	s.e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err != nil {
				status := http.StatusInternalServerError
				msg := "Internal server error"

				if httpErr, ok := err.(HTTPError); ok {
					status = httpErr.Status()
					msg = httpErr.Error()
				}

				c.JSON(status, ErrorResponse{err: msg})
			}

			return err
		}
	})

	s.routes()
}

func (s *server) Start() {
	if s.e == nil {
		panic("Shouldn't start before setting up the server")
	}
	s.e.Start(":4000")
}
