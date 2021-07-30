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

	s.router.Route("/posts", func(r chi.Router) {
		r.Get("/{id}", s.handlePostsFind())
		r.Get("/", s.handlePostsAll())
		r.Group(func(r chi.Router) {
			r.Use(s.authenticated)
			r.Delete("/{id}", s.handlePostsDelete())
			r.Post("/", s.handlePostsCreate())
		})
	})

	s.router.Route("/likes/{id}", func(r chi.Router) {
		r.Use(s.authenticated)
		r.Post("/", s.handleLikesLike())
		r.Delete("/", s.handleLikesUnlike())
	})
}
