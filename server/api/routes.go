package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (s *server) routes() {
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)

	s.router.Post("/auth/login", s.handleLogin())
	s.router.Post("/auth/signUp", s.handleSignUp())
	s.router.Post("/auth/logout", s.handleLogout())

	s.router.Get("/avatar/format", s.handleAvatarFormat())

	s.router.Route("/posts", func(r chi.Router) {
		r.Use(s.authenticated)

		r.Get("/{id}", s.handlePostsFind())
		r.Post("/", s.handlePostsCreate())
		r.Get("/", s.handlePostsAll())
	})
	// s.router.HandleFunc("/posts/{id}", s.authenticated(s.handlePostsFind())).
	// 	Methods("GET")
	// s.router.HandleFunc("/posts", s.authenticated(s.handlePostsCreate())).
	// 	Methods("POST")
	// s.router.HandleFunc("/posts", s.authenticated(s.handlePostsAll())).
	// 	Methods("GET").
	// 	Queries("limit", "{limit}")
}
