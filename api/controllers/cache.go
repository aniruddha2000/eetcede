package controllers

import (
	"net/http"

	j "github.com/aniruddha2000/goEtcd/api/json"
	"github.com/gorilla/mux"
)

func (s *Server) Create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form["key"]
	val := r.Form["val"]

	for i := 0; i < len(key); i++ {
		s.Cache.Store(key[i], val[i])
	}

	j.JSON(w, r, http.StatusCreated, "Record created")
}

func (s *Server) List(w http.ResponseWriter, r *http.Request) {
	records := s.Cache.List()
	j.JSON(w, r, http.StatusOK, records)
}

func (s *Server) Get(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]

	val, err := s.Cache.Get(key)
	if err != nil {
		j.JSON(w, r, http.StatusNotFound, err.Error())
		return
	}
	j.JSON(w, r, http.StatusOK, map[string]string{key: val})
}

func (s *Server) Delete(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]

	err := s.Cache.Delete(key)
	if err != nil {
		j.JSON(w, r, http.StatusNotFound, err.Error())
		return
	}
	j.JSON(w, r, http.StatusNoContent, map[string]string{"data": "delete"})
}
