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
	s.router.Group(func(r chi.Router) {
		r.Use(s.authenticated)
		r.Get("/me", s.handleMe())
		r.Patch("/me", s.handleUpdateMe())
	})

	s.router.Get("/avatar/format", s.handleAvatarFormat())

	s.router.Route("/posts", func(r chi.Router) {
		r.Get("/{id}", s.handlePostsFind())
		r.Get("/", s.handlePostsAll())
		r.Group(func(r chi.Router) {
			r.Use(s.authenticated)
			r.Post("/", s.handlePostsCreate())
		})
	})
}
