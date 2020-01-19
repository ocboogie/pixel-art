package api

import (
	"net/http"

	"github.com/go-chi/chi"
	postService "github.com/ocboogie/pixel-art/services/post"
)

func (s *server) handleLikesLike() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postID := chi.URLParam(r, "id")
		userID := s.getUserID(w, r)

		err := s.post.Like(userID, postID)

		if err != nil {
			apiErr := unexpectedAPIError(err)
			switch err {
			case postService.ErrAlreadyLiked:
				apiErr = errAlreadyLiked
			case postService.ErrNotFound:
				apiErr = errPostNotFound
			case postService.ErrUserNotFound:
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
		userID := s.getUserID(w, r)

		err := s.post.Unlike(userID, postID)

		if err != nil {
			apiErr := unexpectedAPIError(err)
			if err == postService.ErrLikeNotFound {
				apiErr = errLikeNotFound
			}

			s.error(w, r, apiErr)
			return
		}
	}
}
