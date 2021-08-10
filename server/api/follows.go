package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ocboogie/pixel-art/services/profile"
)

func (s *server) handleFollowsFollow() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		followedID := chi.URLParam(r, "id")
		followerID, err := s.getUserID(w, r)
		if err != nil {
			s.error(w, r, unexpectedAPIError(err))
			return
		}

		err = s.profile.Follow(followedID, followerID)

		if err != nil {
			apiErr := unexpectedAPIError(err)
			switch err {
			case profile.ErrFollowAlreadyExists:
				apiErr = errAlreadyFollowing
			case profile.ErrNotFound:
				apiErr = errUserNotFound
			}
			s.error(w, r, apiErr)
			return
		}
	}
}

func (s *server) handleFollowsUnfollow() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		followedID := chi.URLParam(r, "id")
		followerID, err := s.getUserID(w, r)
		if err != nil {
			s.error(w, r, unexpectedAPIError(err))
			return
		}

		err = s.profile.Unfollow(followedID, followerID)

		if err != nil {
			apiErr := unexpectedAPIError(err)
			if err == profile.ErrFollowNotFound {
				apiErr = errFollowNotFound
			}

			s.error(w, r, apiErr)
			return
		}
	}
}
