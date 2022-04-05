package controllers

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/record", s.Create).Methods("POST")
	s.Router.HandleFunc("/record", s.List).Methods("GET")
	s.Router.HandleFunc("/record/{key}", s.Get).Methods("GET")
	s.Router.HandleFunc("/record/{key}", s.Delete).Methods("DELETE")
}
