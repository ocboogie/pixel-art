package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/services/profile"
)

func (s *server) handleMe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := s.getUserID(w, r)
		if err != nil {
			s.error(w, r, unexpectedAPIError(err))
			return
		}

		includes, err := s.getUserIncludes(w, r)
		if err != nil {
			s.error(w, r, unexpectedAPIError(err))
			return
		}

		user, err := s.profile.Find(userID, includes)
		if err != nil {
			s.error(w, r, unexpectedAPIError(err))
			return
		}

		s.respond(w, r, http.StatusOK, user)
	}
}

// TODO: DRY: handlePostsAll and handleUserPosts
func (s *server) handleMePosts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		userID, err := s.getUserID(w, r)
		if err != nil {
			s.error(w, r, unexpectedAPIError(err))
			return
		}

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

func (s *server) handleUpdateMe() http.HandlerFunc {
	type request struct {
		Name   *string        `json:"name,omitempty" validate:"omitempty,min=2,max=64"`
		Avatar *models.Avatar `json:"avatar,omitempty" validate:"omitempty,min=2,max=64"`
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

		userID, err := s.getUserID(w, r)
		if err != nil {
			s.error(w, r, unexpectedAPIError(err))
			return
		}

		user, err := s.profile.Find(userID, profile.UserIncludes{})
		if err != nil {
			s.error(w, r, unexpectedAPIError(err))
			return
		}

		if body.Name != nil {
			user.Name = *body.Name
		}
		if body.Avatar != nil {
			if err := body.Avatar.Validate(s.avatarSpec); err != nil {
				s.error(w, r, errInvalidAvatar)
				return
			}
			user.Avatar = *body.Avatar
		}

		if err := s.profile.Update(user); err != nil {
			s.error(w, r, unexpectedAPIError(err))
			return
		}
	}
}
