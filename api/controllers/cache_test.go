package controllers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/aniruddha2000/eetcede/api/models"
)

func TestCacheServer(t *testing.T) {
	tests := []struct {
		name       string
		server     *Server
		method     func(*Server) *http.Response
		want       string
		statusCode int
	}{
		{
			name:   "Create",
			server: helperServer(t),
			method: func(s *Server) *http.Response {
				req := httptest.NewRequest(http.MethodPost, "/record?key=py&val=con&key=go&val=lang", nil)
				w := httptest.NewRecorder()
				s.Create(w, req)
				return w.Result()
			},
			want:       "Record created",
			statusCode: 201,
		},
		{
			name:   "List",
			server: helperServer(t),
			method: func(s *Server) *http.Response {
				helperServerCreate(t, s)
				req := httptest.NewRequest(http.MethodGet, "/records", nil)
				w := httptest.NewRecorder()
				s.List(w, req)
				return w.Result()
			},
			want: `
				{
					"go": "lang",
					"py": "con"
				}`,
			statusCode: 200,
		},
		{
			name:   "Get existing key-value",
			server: helperServer(t),
			method: func(s *Server) *http.Response {
				helperServerCreate(t, s)
				req := httptest.NewRequest(http.MethodGet, "/record?key=go", nil)
				w := httptest.NewRecorder()
				s.Get(w, req)
				return w.Result()
			},
			want: `
			{
				"go": "lang"
			}`,
			statusCode: 200,
		},
		{
			name:   "Get non existing key value",
			server: helperServer(t),
			method: func(s *Server) *http.Response {
				helperServerCreate(t, s)
				req := httptest.NewRequest(http.MethodGet, "/record?key=goos", nil)
				w := httptest.NewRecorder()
				s.Get(w, req)
				return w.Result()
			},
			want:       "key not found",
			statusCode: 404,
		},
		{
			name:   "Delete exixting key value",
			server: helperServer(t),
			method: func(s *Server) *http.Response {
				helperServerCreate(t, s)
				req := httptest.NewRequest(http.MethodDelete, "/record?key=go", nil)
				w := httptest.NewRecorder()
				s.Delete(w, req)
				return w.Result()
			},
			want:       "",
			statusCode: 204,
		},
		{
			name:   "Delete non exixting key value",
			server: helperServer(t),
			method: func(s *Server) *http.Response {
				helperServerCreate(t, s)
				req := httptest.NewRequest(http.MethodDelete, "/record?key=goos", nil)
				w := httptest.NewRecorder()
				s.Delete(w, req)
				return w.Result()
			},
			want:       "key not found",
			statusCode: 404,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.method(tt.server)
			defer res.Body.Close()

			if res.StatusCode != tt.statusCode {
				t.Errorf("Want %v, got %d", tt.statusCode, res.StatusCode)
			}
			got, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			if reflect.DeepEqual(string(got), tt.want) {
				t.Errorf("expected %v got %v", tt.want, string(got))
			}
		})
	}
}

func helperServer(t *testing.T) *Server {
	t.Helper()
	router := http.NewServeMux()
	return &Server{Router: router, Cache: models.NewCache()}
}

func helperServerCreate(t *testing.T, s *Server) {
	req := httptest.NewRequest(http.MethodPost, "/record?key=py&val=con&key=go&val=lang", nil)
	w := httptest.NewRecorder()
	s.Create(w, req)
}
