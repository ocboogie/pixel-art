package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/ocboogie/pixel-art/services/feed"
	"github.com/ocboogie/pixel-art/services/profile"
)

func (s *server) getUserIncludes(w http.ResponseWriter, r *http.Request) (profile.UserIncludes, error) {
	query := r.URL.Query()

	var err error
	userID := ""
	if query.Get("following") != "" {
		userID, err = s.getUserID(w, r)
		if err != nil {
			return profile.UserIncludes{}, err
		}
	}

	return profile.UserIncludes{
		Following: userID,
	}, nil
}

func (s *server) handleUsersFind() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "id")

		includes, err := s.getUserIncludes(w, r)
		if err != nil {
			s.error(w, r, unexpectedAPIError(err))
			return
		}

		user, err := s.profile.Find(userID, includes)

		if err != nil {
			if err == feed.ErrNotFound {
				s.error(w, r, errUserNotFound)
				return
			}
			s.error(w, r, unexpectedAPIError(err))
			return
		}

		s.respond(w, r, http.StatusOK, user.HideSensitive())
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

		includes, err := s.getPostIncludes(w, r)
		if err != nil {
			s.error(w, r, unexpectedAPIError(err))
			return
		}

		posts, err := s.feed.PostsByUser(userID, limit, after, includes)
		if err != nil {
			s.error(w, r, unexpectedAPIError(err))
			return
		}

		s.respond(w, r, http.StatusOK, posts)
	}
}
