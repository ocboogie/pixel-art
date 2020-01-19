package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (s *server) routes() {
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)

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

	s.router.Get("/avatar/format", s.handleAvatarFormat())
	s.router.Get("/art/format", s.handleArtFormat())

	s.router.Route("/users/{id}", func(r chi.Router) {
		r.Get("/", s.handleUsersFind())
		r.Get("/posts", s.handleUsersPosts())
	})

	s.router.Route("/posts", func(r chi.Router) {
		r.Get("/{id}", s.handlePostsFind())
		r.Get("/", s.handlePostsAll())
		r.Group(func(r chi.Router) {
			r.Use(s.authenticated)
			r.Post("/", s.handlePostsCreate())
		})
	})

	s.router.Route("/likes/{id}", func(r chi.Router) {
		r.Use(s.authenticated)
		r.Post("/", s.handleLikesLike())
		r.Delete("/", s.handleLikesUnlike())
	})
}
