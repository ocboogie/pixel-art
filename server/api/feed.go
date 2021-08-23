package api

import (
	"net/http"
	"strconv"
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
