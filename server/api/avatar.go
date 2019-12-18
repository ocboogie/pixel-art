package api

import (
	"net/http"
)

func (s *server) handleAvatarFormat() http.HandlerFunc {
	// Since this shouldn't change, we can just run it once and save it.
	format := s.avatar.Format()
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, format)
	}
}
