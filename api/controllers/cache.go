package controllers

import (
	"log"
	"net/http"

	j "github.com/aniruddha2000/eetcede/api/json"
)

func (s *Server) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		key := r.Form["key"]
		val := r.Form["val"]

		for i := 0; i < len(key); i++ {
			s.Cache.Store(key[i], val[i])
		}

		j.JSON(w, r, http.StatusCreated, "Record created")
	} else {
		j.JSON(w, r, http.StatusBadRequest, "POST Request accepted")
	}
}

func (s *Server) List(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		records := s.Cache.List()
		j.JSON(w, r, http.StatusOK, records)
	} else {
		j.JSON(w, r, http.StatusBadRequest, "GET Request accepted")
	}
}

func (s *Server) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		keys, ok := r.URL.Query()["key"]
		if !ok || len(keys[0]) < 1 {
			log.Println("Url Param 'key' is missing")
			return
		}
		key := keys[0]

		val, err := s.Cache.Get(key)
		if err != nil {
			j.JSON(w, r, http.StatusNotFound, err.Error())
			return
		}
		j.JSON(w, r, http.StatusOK, map[string]string{key: val})
	} else {
		j.JSON(w, r, http.StatusBadRequest, "POST Request accepted")
	}
}

func (s *Server) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		keys, ok := r.URL.Query()["key"]
		if !ok || len(keys[0]) < 1 {
			log.Println("Url Param 'key' is missing")
			return
		}
		key := keys[0]

		err := s.Cache.Delete(key)
		if err != nil {
			j.JSON(w, r, http.StatusNotFound, err.Error())
			return
		}
		j.JSON(w, r, http.StatusNoContent, map[string]string{"data": "delete"})
	} else {
		j.JSON(w, r, http.StatusBadRequest, "DELETE Request accepted")
	}
}
