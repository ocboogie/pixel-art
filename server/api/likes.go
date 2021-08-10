package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ocboogie/pixel-art/services/feed"
)

func (s *server) handleLikesLike() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postID := chi.URLParam(r, "id")
		userID, err := s.getUserID(w, r)
		if err != nil {
			s.error(w, r, unexpectedAPIError(err))
			return
		}

		err = s.feed.Like(userID, postID)

		if err != nil {
			apiErr := unexpectedAPIError(err)
			switch err {
			case feed.ErrAlreadyLiked:
				apiErr = errAlreadyLiked
			case feed.ErrNotFound:
				apiErr = errPostNotFound
			case feed.ErrUserNotFound:
				apiErr = errInvalidUserState
			}
			s.error(w, r, apiErr)
			return
		}
	}
}

func (s *server) handleLikesUnlike() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postID := chi.URLParam(r, "id")
		userID, err := s.getUserID(w, r)
		if err != nil {
			s.error(w, r, unexpectedAPIError(err))
			return
		}

		err = s.feed.Unlike(userID, postID)

		if err != nil {
			apiErr := unexpectedAPIError(err)
			if err == feed.ErrLikeNotFound {
				apiErr = errLikeNotFound
			}

			s.error(w, r, apiErr)
			return
		}
	}
}
