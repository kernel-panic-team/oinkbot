package tests

import (
	"testing"
	"time"

	"github.com/kernel-panic-team/oinkbot/internal/config"
	"github.com/kernel-panic-team/oinkbot/internal/service"
	"github.com/kernel-panic-team/oinkbot/pkg/httpclient"
)

func TestPingPorkbun(t *testing.T) {
	cfg := config.New()
	cli := httpclient.New(cfg)
	srv := service.New(cfg, nil, nil, cli)

	tests := []struct {
		name      string
		apiKey    string
		secretKey string
		wantErr   bool
	}{
		{
			name:      "normal ping",
			apiKey:    cfg.APIKey,
			secretKey: cfg.SecretAPIKey,
			wantErr:   false,
		},
		{
			name:      "wrong api key",
			apiKey:    "wrong",
			secretKey: cfg.SecretAPIKey,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg.APIKey = tt.apiKey
			cfg.SecretAPIKey = tt.secretKey

			time.Sleep(1 * time.Second)

			_, err := srv.PingPorkbun()
			if (err != nil) != tt.wantErr {
				t.Errorf("PingPorkbun() returned error: %v, wantErr: %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetPorkbunCerts(t *testing.T) {
	cfg := config.New()
	cli := httpclient.New(cfg)
	srv := service.New(cfg, nil, nil, cli)

	time.Sleep(1 * time.Second)

	certs, err := srv.GetPorkbunCerts()
	if err != nil {
		t.Errorf("GetPorkbunCerts() returner error: %v; want nil", err)
		return
	}

	if certs == nil {
		t.Errorf("certs = %v; want non-nil", certs)
		return
	}
}
