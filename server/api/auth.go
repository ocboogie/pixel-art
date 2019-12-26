package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/services/auth"
)

var (
	errInvalidCredentials = newSimpleAPIError(http.StatusBadRequest, true, "Username or password is invalid")
	errEmailAlreadyInUse  = newSimpleAPIError(http.StatusBadRequest, true, "Email already in use")
	errUnauthenticated    = newSimpleAPIError(http.StatusUnauthorized, true, "You must be logged in")
	errInvalidAvatar      = newSimpleAPIError(http.StatusBadRequest, false, "Invalid avatar")
)

const sessionCookie = "sessionId"

type userIDContextKey struct{}

func (s *server) saveSession(w http.ResponseWriter, r *http.Request, session *models.Session) {
	http.SetCookie(w, &http.Cookie{
		Name:  sessionCookie,
		Value: session.ID,

		Expires: session.ExpiresAt,
		Path:    "/",

		// TODO: Enable this in producation
		// Secure: true,
		// HttpOnly: true,
	})
}

func (s *server) deleteSession(w http.ResponseWriter, r *http.Request) {
	// TODO: Make all shared cookie data in one place
	http.SetCookie(w, &http.Cookie{
		Name:  sessionCookie,
		Value: "",

		Expires: time.Unix(0, 0),
		Path:    "/",

		// TODO: Enable this in producation
		// Secure: true,
		// HttpOnly: true,
	})
}

func (s *server) getSessionID(w http.ResponseWriter, r *http.Request) string {
	cookie, err := r.Cookie(sessionCookie)

	if err != nil {
		return ""
	}

	return cookie.Value
}

func (s *server) authenticated(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		sessionID := s.getSessionID(w, r)

		if sessionID == "" {
			s.error(w, r, errUnauthenticated)
			return
		}
		userID, err := s.auth.VerifySession(sessionID)
		if err != nil {
			if err == auth.ErrSessionNotFound || err == auth.ErrExpiredSession {
				s.error(w, r, errUnauthenticated)
				return
			}
			s.error(w, r, unexpectedAPIError(err))
			return
		}

		ctx := context.WithValue(r.Context(), userIDContextKey{}, userID)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func (s *server) getUserID(w http.ResponseWriter, r *http.Request) string {
	// TODO: Test if this works
	id, ok := r.Context().Value(userIDContextKey{}).(string)
	if !ok {
		panic("getUserID was call without the authenticated middleware")
	}
	return id
}

func (s *server) handleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var body models.UserCredentials
		err := decoder.Decode(&body)

		if err != nil {
			s.error(w, r, errInvalidJSON(err))
			return
		}

		if err := body.Validate(s.validate); err != nil {
			// FIXME:
			s.error(w, r, errInvalidBody(err))
			return
		}

		session, err := s.auth.Login(&body)

		if err != nil {
			if err == auth.ErrInvalidCredentials {
				s.error(w, r, errInvalidCredentials)
				return
			}

			s.error(w, r, unexpectedAPIError(err))
			return
		}

		s.saveSession(w, r, session)
	}
}

func (s *server) handleSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var body models.UserNew
		err := decoder.Decode(&body)

		if err != nil {
			s.error(w, r, errInvalidJSON(err))
			return
		}

		body.Avatar = s.avatar.GenerateRandom()

		if err := body.Validate(s.validate); err != nil {
			// FIXME:
			s.error(w, r, errInvalidBody(err))
			return
		}

		session, err := s.auth.SignUp(&body)

		if err != nil {
			if err == auth.ErrEmailAlreadyInUse {
				s.error(w, r, errEmailAlreadyInUse)
				return
			}
			if err == auth.ErrInvalidAvatar {
				s.error(w, r, errInvalidAvatar)
				return
			}

			s.error(w, r, unexpectedAPIError(err))
			return
		}

		s.saveSession(w, r, session)
	}
}

func (s *server) handleLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionID := s.getSessionID(w, r)
		if sessionID == "" {
			s.error(w, r, errUnauthenticated)
			return
		}
		err := s.auth.Logout(sessionID)
		if err != nil {
			s.error(w, r, unexpectedAPIError(err))
			return
		}
		s.deleteSession(w, r)
	}
}
