package api

import "net/http"

func (s *Server) routes() {
	s.Router.HandleFunc("/test_delete", s.HandleTestDelete()).Methods(http.MethodPost)
	s.Router.HandleFunc("/test_overwrite", s.HandleTestOverwrite()).Methods(http.MethodPost)

}
