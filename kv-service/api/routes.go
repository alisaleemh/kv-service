package api

import "net/http"

func (s *Server) routes() {
	s.Router.HandleFunc("/kv/{key}", s.HandleGetKey()).Methods(http.MethodGet)
	s.Router.HandleFunc("/kv/{key}", s.HandleInsertKey()).Methods(http.MethodPut)
	s.Router.HandleFunc("/kv/{key}", s.HandleDeleteKey()).Methods(http.MethodDelete)

}
