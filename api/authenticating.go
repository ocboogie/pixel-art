package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ocboogie/pixel-art/models"
)

const sessionIDCookieName = "session_id"

func (s *server) Authenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(sessionIDCookieName)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "You must login")
		}
		userID, err := s.authenticating.VerifySession(cookie.Value)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Session")
		}
		c.Set("userID", userID)

		return next(c)
	}
}

func (s *server) handlerSignUp(c echo.Context) error {
	userInput := new(models.UserInput)

	if err := c.Bind(userInput); err != nil {
		return err
	}
	if err := userInput.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	_, err := s.authenticating.SignUp(userInput)

	if err != nil {
		// TODO:
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (s *server) handlerLogin(c echo.Context) error {
	userInput := new(models.UserInput)

	if err := c.Bind(userInput); err != nil {
		return err
	}
	if err := userInput.Validate(); err != nil {
		return err
	}

	sessionID, err := s.authenticating.Login(userInput.Email, userInput.Password)

	if err != nil {
		return err
	}

	// TODO: Move this somewhere else and use secure
	sessionCookie := &http.Cookie{
		Name:    sessionIDCookieName,
		Value:   sessionID,
		Expires: time.Now().Add(time.Duration(s.config.SessionLifetime)),
	}

	c.SetCookie(sessionCookie)

	return c.NoContent(http.StatusOK)
}
