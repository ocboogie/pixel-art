package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	userService "github.com/ocboogie/pixel-art/services/user"
)

func (s *server) handleUsersFind() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "id")
		user, err := s.user.Find(userID)

		if err != nil {
			if err == userService.ErrNotFound {
				s.error(w, r, errUserNotFound)
				return
			}
			s.error(w, r, unexpectedAPIError(err))
			return
		}

		s.respond(w, r, http.StatusOK, user)
	}
}

func (s *server) handleUsersPosts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: DRY: handlePostsAll and handleMePosts
		limit := 50
		limitQuery := r.URL.Query().Get("limit")
		if limitQuery != "" {
			i, err := strconv.Atoi(limitQuery)

			if err != nil {
				s.error(w, r, errInvalidLimit)
				return
			}

			limit = i
		}

		var after *time.Time = nil
		afterQuery := r.URL.Query().Get("after")
		if afterQuery != "" {
			afterDate, err := time.Parse(time.RFC3339, afterQuery)

			if err != nil {
				s.error(w, r, errInvalidAfter)
				return
			}

			after = &afterDate
		}

		userID := chi.URLParam(r, "id")
		posts, err := s.feed.PostsByUser(userID, limit, after)
		if err != nil {
			s.error(w, r, unexpectedAPIError(err))
			return
		}

		s.respond(w, r, http.StatusOK, posts)
	}
}
