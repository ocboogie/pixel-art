package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/services/feed"
)

func (s *server) getPostIncludes(w http.ResponseWriter, r *http.Request) (feed.PostIncludes, error) {
	var err error
	userID := ""
	if paramExists(r, "liked") {
		userID, err = s.getUserID(w, r)
		if err != nil {
			return feed.PostIncludes{}, err
		}
	}

	return feed.PostIncludes{
		Author: paramExists(r, "author"),
		Likes:  paramExists(r, "likes"),
		Liked:  userID,
	}, nil
}

func (s *server) handlePostsFind() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postID := chi.URLParam(r, "id")

		includes, err := s.getPostIncludes(w, r)
		if err != nil {
			s.error(w, r, unexpectedAPIError(err))
			return
		}

		post, err := s.feed.Find(postID, includes)

		if err != nil {
			if err == feed.ErrNotFound {
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
		Title string     `json:"title" validate:"required,min=2,max=256"`
		Art   models.Art `json:"art"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var body request
		err := decoder.Decode(&body)

		if err != nil {
			s.error(w, r, errInvalidJSON(err))
			return
		}

		userID, err := s.getUserID(w, r)
		if err != nil {
			s.error(w, r, unexpectedAPIError(err))
			return
		}

		newPost := models.PostNew{
			Title:  body.Title,
			Art:    body.Art,
			UserID: userID,
		}

		if err := newPost.Validate(s.validate, s.artSpec); err != nil {
			// FIXME:
			s.error(w, r, errInvalidBody(err))
			return
		}

		id, err := s.feed.Create(newPost)

		if err != nil {
			s.error(w, r, unexpectedAPIError(err))
			return
		}

		s.respond(w, r, http.StatusOK, id)
	}
}

func (s *server) handlePostsDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postID := chi.URLParam(r, "id")

		post, err := s.feed.Find(postID, feed.PostIncludes{})

		if err != nil {
			if err == feed.ErrNotFound {
				s.error(w, r, errPostNotFound)
				return
			}
			s.error(w, r, unexpectedAPIError(err))
			return
		}

		userID, err := s.getUserID(w, r)
		if err != nil {
			s.error(w, r, unexpectedAPIError(err))
			return
		}

		// TODO: Abstract this away somewhere.
		if userID != post.AuthorID {
			s.error(w, r, errInvalidPermissions)
			return
		}

		err = s.feed.Delete(post.ID)

		if err != nil {
			if err == feed.ErrNotFound {
				s.error(w, r, errPostNotFound)
				return
			}
			s.error(w, r, unexpectedAPIError(err))
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) handlePostsAll() http.HandlerFunc {
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

		includes, err := s.getPostIncludes(w, r)
		if err != nil {
			s.error(w, r, unexpectedAPIError(err))
			return
		}

		posts, err := s.feed.Latest(limit, after, includes)
		if err != nil {
			s.error(w, r, unexpectedAPIError(err))
			return
		}
		s.respond(w, r, http.StatusOK, posts)
	}
}
