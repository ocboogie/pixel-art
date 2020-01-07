package api

import (
	"encoding/json"
	"net/http"
)

func (s *server) handleMe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := s.getUserID(w, r)
		user, err := s.user.Find(userID)
		if err != nil {
			s.error(w, r, unexpectedAPIError(err))
		}

		s.respond(w, r, http.StatusOK, user)
	}
}

func (s *server) handleUpdateMe() http.HandlerFunc {
	type request struct {
		Name   *string `json:"name,omitempty" validate:"omitempty,min=2,max=64"`
		Avatar *string `json:"avatar,omitempty" validate:"omitempty,min=2,max=64"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var body request
		err := decoder.Decode(&body)

		if err != nil {
			s.error(w, r, errInvalidJSON(err))
			return
		}

		if err := s.validate.Struct(body); err != nil {
			// FIXME:
			s.error(w, r, errInvalidBody(err))
			return
		}

		user, err := s.user.Find(s.getUserID(w, r))
		if err != nil {
			s.error(w, r, unexpectedAPIError(err))
			return
		}

		if body.Name != nil {
			user.Name = *body.Name
		}
		if body.Avatar != nil {
			if !s.avatar.Validate(*body.Avatar) {
				s.error(w, r, errInvalidAvatar)
				return
			}
			user.Avatar = *body.Avatar
		}

		if err := s.user.Update(user); err != nil {
			s.error(w, r, unexpectedAPIError(err))
			return
		}
	}
}
