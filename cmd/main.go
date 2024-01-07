package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kernel-panic-team/oinkbot/internal/config"
	"github.com/kernel-panic-team/oinkbot/internal/service"
	"github.com/kernel-panic-team/oinkbot/pkg/cron"
	"github.com/kernel-panic-team/oinkbot/pkg/httpclient"
	"github.com/kernel-panic-team/oinkbot/pkg/httpserver"
)

func main() {
	cfg := config.New()

	crn := cron.New(cfg)
	httpsrv := httpserver.New(cfg)
	httpcli := httpclient.New(cfg)
	srv := service.New(cfg, crn, httpsrv, httpcli)

	err := srv.Start()
	if err != nil {
		log.Fatalf("service start failed: %v", err)
	}
	defer srv.Stop()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)
	<-ch

	log.Printf("service stopped")
}
