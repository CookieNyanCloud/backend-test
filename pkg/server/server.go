package server

import (
	"context"
	"net/http"

	"github.com/cookienyancloud/avito-backend-test/internal/config"
)

type server struct {
	httpServer *http.Server
}

//new server
func NewServer(cfg *config.Config, handler http.Handler) *server {
	return &server{
		httpServer: &http.Server{
			Addr:           ":" + cfg.HTTP.Port,
			Handler:        handler,
			ReadTimeout:    cfg.HTTP.ReadTimeout,
			WriteTimeout:   cfg.HTTP.WriteTimeout,
			MaxHeaderBytes: cfg.HTTP.MaxHeaderMegabytes << 20,
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
