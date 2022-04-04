package controllers

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/record", s.CreateRecord).Methods("POST")
	s.Router.HandleFunc("/record", s.GetAllRecords).Methods("GET")
	s.Router.HandleFunc("/record/{key}", s.GetRecord).Methods("GET")
	s.Router.HandleFunc("/record/{key}", s.DeleteRecord).Methods("DELETE")
}
