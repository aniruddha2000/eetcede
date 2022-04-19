package controllers

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/record", s.Create)
	s.Router.HandleFunc("/records", s.List)
	s.Router.HandleFunc("/get/record", s.Get)
	s.Router.HandleFunc("/del/record", s.Delete)
}
