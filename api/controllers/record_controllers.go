package controllers

import (
	"net/http"

	j "github.com/aniruddha2000/goEtcd/api/json"
	"github.com/gorilla/mux"
)

func (s *Server) CreateRecord(w http.ResponseWriter, r *http.Request) {
	// key := r.URL.Query().Get("key")
	// val := r.URL.Query().Get("val")

	r.ParseForm()
	key := r.Form["key"]
	val := r.Form["val"]

	for i := 0; i < len(key); i++ {
		s.Cache.Store(key[i], val[i])
	}

	j.JSON(w, r, http.StatusCreated, "Record created")
}

func (s *Server) GetAllRecords(w http.ResponseWriter, r *http.Request) {
	records := s.Cache.FindAllRecords()
	j.JSON(w, r, http.StatusOK, records)
}

func (s *Server) GetRecord(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]

	val, err := s.Cache.FindRecord(key)
	if err != nil {
		j.JSON(w, r, http.StatusNotFound, err.Error())
		return
	}
	j.JSON(w, r, http.StatusOK, map[string]string{key: val})
}

func (s *Server) DeleteRecord(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]

	err := s.Cache.DeleteRecordByKey(key)
	if err != nil {
		j.JSON(w, r, http.StatusNotFound, err.Error())
		return
	}
	j.JSON(w, r, http.StatusNoContent, map[string]string{"data": "delete"})
}
