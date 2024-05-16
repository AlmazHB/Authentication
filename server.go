package web

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpserver *http.Server
}

func New() *Server {
	return &Server{}
}
func (sv *Server) Run(port string, handler http.Handler) error {
	sv.httpserver = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 28, //1Mb
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return sv.httpserver.ListenAndServe()
}
func (sv *Server) ShutDown(ctx context.Context) error {

	return sv.httpserver.Shutdown(ctx)
}
