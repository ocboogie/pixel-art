package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/ocboogie/pixel-art/services/feed"
	"github.com/ocboogie/pixel-art/services/profile"
)

func (s *server) getUserIncludes(w http.ResponseWriter, r *http.Request) (profile.UserIncludes, error) {
	var err error
	userID := ""
	if paramExists(r, "isFollowing") {
		userID, err = s.getUserID(w, r)
		if err != nil {
			return profile.UserIncludes{}, err
		}
	}

	return profile.UserIncludes{
		IsFollowing:    userID,
		Followers:      paramExists(r, "followers"),
		FollowingCount: paramExists(r, "FollowingCount"),
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
		var err error
		limit := 50
		if paramExists(r, "limit") {
			limit, err = paramNumber(r, "limit")
			if err != nil {
				s.error(w, r, errInvalidLimit)
				return
			}
		}

		var after *time.Time = nil
		if paramExists(r, "after") {
			after, err = paramTime(r, "after")
			if err != nil {
				s.error(w, r, errInvalidAfter)
				return
			}
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
