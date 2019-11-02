package api

func (s *server) routes() {
	s.e.POST("/signUp", s.handlerSignUp)
	s.e.POST("/login", s.handlerLogin)
}
