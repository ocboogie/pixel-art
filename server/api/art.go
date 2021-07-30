package api

import (
	"net/http"
)

func (s *server) handleArtSpec() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, s.artSpec)
	}
}
