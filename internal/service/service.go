package service

import (
	"context"
	"log"

	"github.com/kernel-panic-team/oinkbot/internal/config"
)

type httpServer interface {
	Start() error
	Stop(ctx context.Context) error
}

type httpClient interface {
	GetRawResponse(url string) ([]byte, error)
}

type cronClient interface {
	Start(func()) error
	Stop()
}

type Service struct {
	config *config.Config
	cron   cronClient
	server httpServer
	client httpClient
}

func (s *Service) Start() error {
	err := s.cron.Start(s.worker)
	if err != nil {
		return err
	}
	err = s.server.Start()
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), s.config.ShutdownTimeout)
	defer cancel()

	s.cron.Stop()
	err := s.server.Stop(ctx)
	if err != nil {
		log.Printf("error stopping http server: %v", err)
	}
}

func New(config *config.Config, cron cronClient, server httpServer, client httpClient) *Service {
	return &Service{
		config: config,
		cron:   cron,
		server: server,
		client: client,
	}
}
