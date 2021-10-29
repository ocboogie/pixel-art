package api

import (
	"github.com/go-chi/chi"
)

func (s *server) routes() {
	s.router.Route("/auth", func(r chi.Router) {
		r.Post("/login", s.handleLogin())
		r.Post("/signUp", s.handleSignUp())
		r.Post("/logout", s.handleLogout())
	})
	s.router.Route("/me", func(r chi.Router) {
		r.Use(s.authenticated)
		r.Get("/", s.handleMe())
		r.Patch("/", s.handleUpdateMe())
		r.Get("/posts", s.handleMePosts())
	})

	s.router.Get("/avatar/spec", s.handleAvatarSpec())
	s.router.Get("/art/spec", s.handleArtSpec())

	s.router.Route("/users/{id}", func(r chi.Router) {
		r.Get("/", s.handleUsersFind())
		r.Get("/posts", s.handleUsersPosts())
	})

	s.router.Route("/feed", func(r chi.Router) {
		r.Use(s.authenticated)
		r.Get("/", s.handleFeedGet())
	})

	s.router.Route("/posts", func(r chi.Router) {
		r.Get("/{id}", s.handlePostsFind())
		r.Get("/", s.handlePostsAll())
		r.Group(func(r chi.Router) {
			r.Use(s.authenticated)
			r.Delete("/{id}", s.handlePostsDelete())
			r.Post("/", s.handlePostsCreate())
		})
	})

	// TODO: Make this stateless (see https://stackoverflow.com/a/5668406/4910911
	// and https://www.mscharhag.com/api-design/rest-many-to-many-relations)
	// PUT `/posts/{postID}/likes/{userID}
	// DELETE `/posts/{postID}/likes/{userID}
	s.router.Route("/likes/{id}", func(r chi.Router) {
		r.Use(s.authenticated)
		r.Put("/", s.handleLikesLike())
		r.Delete("/", s.handleLikesUnlike())
	})

	// TODO: Make this stateless:
	// PUT `/users/{followedID}/followers/{followerID}
	// DELETE `/users/{followedID}/followers/{followerID}
	s.router.Route("/follows/{id}", func(r chi.Router) {
		r.Use(s.authenticated)
		r.Put("/", s.handleFollowsFollow())
		r.Delete("/", s.handleFollowsUnfollow())
	})
}
