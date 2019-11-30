package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ocboogie/pixel-art/models"
	postService "github.com/ocboogie/pixel-art/services/post"
)

var (
	errPostNotFound = newSimpleAPIError(http.StatusNotFound, false, "Post not found")
)

func (s *server) handlePostsFind() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		postID := vars["id"]
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
		// TODO: validate Data
		Data string `json:"data"`
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

		id, err := s.post.Create(newPost)

		if err != nil {
			s.error(w, r, unexpectedAPIError(err))
			return
		}

		s.respond(w, r, http.StatusOK, id)
	}
}
