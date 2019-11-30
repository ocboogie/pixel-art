package api

func (s *server) routes() {
	s.router.HandleFunc("/auth/login", s.handleLogin()).Methods("POST")
	s.router.HandleFunc("/auth/signUp", s.handleSignUp()).Methods("POST")

	s.router.HandleFunc("/posts/{id}", s.authenticated(s.handlePostsFind())).Methods("GET")
	s.router.HandleFunc("/posts", s.authenticated(s.handlePostsCreate())).Methods("POST")
	// s.router.HandleFunc("/posts", s.handlePostsAll()).Methods("GET")
}
