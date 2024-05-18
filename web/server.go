package web

import (
	"net/http"
)

type Server struct{}

func (s *Server) Run(port string, handler http.Handler) error {
	return http.ListenAndServe(":"+port, handler)
}
