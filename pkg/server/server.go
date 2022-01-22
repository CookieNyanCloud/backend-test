package server

import (
	"context"
	"net/http"
	"time"
)

type server struct {
	httpServer *http.Server
}

//new server
func NewServer(port string, readTimeout, writeTimeout time.Duration, maxHeaderBytes int, handler http.Handler) *server {
	return &server{
		httpServer: &http.Server{
			Addr:           ":" + port,
			Handler:        handler,
			ReadTimeout:    readTimeout,
			WriteTimeout:   writeTimeout,
			MaxHeaderBytes: maxHeaderBytes << 20,
		},
	}
}

//start server
func (s *server) Run() error {
	return s.httpServer.ListenAndServe()
}

//graceful shutdown
func (s *server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
