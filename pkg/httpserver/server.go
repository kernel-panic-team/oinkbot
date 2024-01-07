package httpserver

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kernel-panic-team/oinkbot/internal/config"
)

type Server struct {
	cfg        *config.Config
	httpServer *http.Server
}

func New(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Start() error {
	mux := chi.NewMux()

	mux.Route("/healthcheck", func(r chi.Router) {
		r.Mount("/", http.HandlerFunc(checkHealth))
	})

	s.httpServer = &http.Server{
		Addr:              s.cfg.AddressAndPort,
		Handler:           mux,
		ReadHeaderTimeout: s.cfg.ReadHeadersTimeout,
	}

	s.httpServer.SetKeepAlivesEnabled(false)

	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	if s.httpServer == nil {
		return nil
	}
	return s.httpServer.Shutdown(ctx)
}
