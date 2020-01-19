package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
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

// TODO: DRY: handlePostsAll and handleUserPosts
func (s *server) handleMePosts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		userID := s.getUserID(w, r)
		posts, err := s.post.PostsByUser(userID, limit, after)
		if err != nil {
			s.error(w, r, unexpectedAPIError(err))
			return
		}

		s.respond(w, r, http.StatusOK, posts)
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
