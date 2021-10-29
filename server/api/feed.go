package api

import (
	"net/http"
	"time"
)

func (s *server) handleFeedGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := s.getUserID(w, r)
		if err != nil {
			s.error(w, r, unexpectedAPIError(err))
			return
		}

		// TODO: Abstract away default limits
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

		posts, err := s.feed.Feed(userID, limit, after, includes)
		if err != nil {
			s.error(w, r, unexpectedAPIError(err))
			return
		}
		s.respond(w, r, http.StatusOK, posts)
	}
}
