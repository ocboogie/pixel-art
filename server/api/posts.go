package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/ocboogie/pixel-art/models"
	postService "github.com/ocboogie/pixel-art/services/post"
)

func (s *server) handlePostsFind() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postID := chi.URLParam(r, "id")
		post, err := s.post.Find(postID)

		if err != nil {
			if err == postService.ErrNotFound {
				s.error(w, r, errPostNotFound)
				return
			}
			s.error(w, r, unexpectedAPIError(err))
			return
		}

		s.respond(w, r, http.StatusOK, post)
	}
}

func (s *server) handlePostsCreate() http.HandlerFunc {
	type request struct {
		Title string `json:"title" validate:"required,min=2,max=256"`
		Data  string `json:"data"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var body request
		err := decoder.Decode(&body)

		if err != nil {
			s.error(w, r, errInvalidJSON(err))
			return
		}

		userID := s.getUserID(w, r)

		newPost := models.PostNew{
			Title:  body.Title,
			Data:   body.Data,
			UserID: userID,
		}

		if err := newPost.Validate(s.validate); err != nil {
			// FIXME:
			s.error(w, r, errInvalidBody(err))
			return
		}

		if !s.art.Validate(newPost.Data) {
			s.error(w, r, errInvalidArt)
			return
		}

		id, err := s.post.Create(newPost)

		if err != nil {
			s.error(w, r, unexpectedAPIError(err))
			return
		}

		s.respond(w, r, http.StatusOK, id)
	}
}

func (s *server) handlePostsAll() http.HandlerFunc {
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

		posts, err := s.post.Latest(limit, after)
		if err != nil {
			s.error(w, r, unexpectedAPIError(err))
			return
		}
		s.respond(w, r, http.StatusOK, posts)
	}
}
