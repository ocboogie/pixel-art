package api

import (
	"bytes"
	"image"
	"image/gif"
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
		for _, post := range posts {
			art, err := post.Art.Decode(s.artSpec)
			if err != nil {
				panic("")
			}

			paletted := art.ToPaletted()
			var artGif gif.GIF
			artGif.BackgroundIndex = 0

			artGif.Delay = make([]int, 1)
			artGif.Delay[0] = 0

			artGif.Image = make([]*image.Paletted, 1)
			artGif.Image[0] = &paletted

			bytes := bytes.NewBuffer(make([]byte, 0))
			if err := gif.EncodeAll(bytes, &artGif); err != nil {
				println(err.Error())
			}

			post.Art = bytes.Bytes()
			println(post.Art)
		}
		s.respond(w, r, http.StatusOK, posts)
	}
}
