package tests

import (
	"testing"
	"time"

	"github.com/kernel-panic-team/oinkbot/internal/config"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name                     string
		isDefaultsTest           bool
		addressAndPort           string
		cronInterval             config.Days
		readHeadersInterval      time.Duration
		requestTimeout           time.Duration
		shutdownTimeout          time.Duration
		porkbunCertsURI          string
		porkbunPingURI           string
		certsFolderName          string
		intermediateCertFilename string
		certChainFilename        string
		privateKeyFilename       string
		publicKeyFilename        string
	}{
		{
			name:                     "test setting default envs",
			isDefaultsTest:           true,
			addressAndPort:           ":8080",
			cronInterval:             config.Days(5 * 24 * time.Hour),
			readHeadersInterval:      5 * time.Second,
			requestTimeout:           5 * time.Second,
			shutdownTimeout:          5 * time.Second,
			porkbunCertsURI:          "https://porkbun.com/api/json/v3/ssl/retrieve",
			porkbunPingURI:           "https://porkbun.com/api/json/v3/ping",
			intermediateCertFilename: "intermediate.crt",
			certChainFilename:        "cert-chain.crt",
			privateKeyFilename:       "private.key",
			publicKeyFilename:        "public.key",
		},
		{
			name:                     "test setting custom envs",
			isDefaultsTest:           false,
			addressAndPort:           ":8085",
			cronInterval:             config.Days(10 * 24 * time.Hour),
			readHeadersInterval:      55 * time.Second,
			requestTimeout:           55 * time.Second,
			shutdownTimeout:          55 * time.Second,
			porkbunCertsURI:          "https://porkbun.com/api/json/v3/ssl/retrieve/custom",
			porkbunPingURI:           "https://porkbun.com/api/json/v3/ping/custom",
			intermediateCertFilename: "intermediate.crt.1",
			certChainFilename:        "cert-chain.crt.1",
			privateKeyFilename:       "private.key.1",
			publicKeyFilename:        "public.key.1",
		},
	}

	for _, tt := range tests {
		if !tt.isDefaultsTest {
			setCustomEnvs(t)
		}
		cfg := config.New()
		t.Run(tt.name, func(t *testing.T) {
			if cfg.AddressAndPort != tt.addressAndPort {
				t.Fatalf("cfg.AddressAndPort: %v; addressAndPort: %v", cfg.AddressAndPort, tt.addressAndPort)
			}
			if cfg.CronInterval != tt.cronInterval {
				t.Fatalf("cfg.CronInterval: %v; cronInterval: %v", cfg.CronInterval, tt.cronInterval)
			}
			if cfg.ReadHeadersTimeout != tt.readHeadersInterval {
				t.Fatalf("cfg.ReadHeadersTimeout: %v; readHeadersInterval: %v", cfg.ReadHeadersTimeout, tt.readHeadersInterval)
			}
			if cfg.RequestTimeout != tt.requestTimeout {
				t.Fatalf("cfg.RequestTimeout: %v; requestTimeout: %v", cfg.RequestTimeout, tt.requestTimeout)
			}
			if cfg.ShutdownTimeout != tt.shutdownTimeout {
				t.Fatalf("cfg.ShutdownTimeout: %v; shutdownTimeout: %v", cfg.ShutdownTimeout, tt.shutdownTimeout)
			}
			if cfg.PorkbunCertsURI != tt.porkbunCertsURI {
				t.Fatalf("cfg.PorkbunCertsURI: %v; porkbunCertsURI: %v", cfg.PorkbunCertsURI, tt.porkbunCertsURI)
			}
			if cfg.PorkbunPingURI != tt.porkbunPingURI {
				t.Fatalf("cfg.PorkbunPingURI: %v; porkbunPingURI: %v", cfg.PorkbunPingURI, tt.porkbunPingURI)
			}
			if cfg.IntermediateCertFilename != tt.intermediateCertFilename {
				t.Fatalf("cfg.IntermediateCertFilename: %v; intermediateCertFilename: %v", cfg.IntermediateCertFilename, tt.intermediateCertFilename)
			}
			if cfg.CertChainFilename != tt.certChainFilename {
				t.Fatalf("cfg.CertChainFilename: %v; certChainFilename: %v", cfg.CertChainFilename, tt.certChainFilename)
			}
			if cfg.PrivateKeyFilename != tt.privateKeyFilename {
				t.Fatalf("cfg.PrivateKeyFilename: %v; privateKeyFilename: %v", cfg.PrivateKeyFilename, tt.privateKeyFilename)
			}
			if cfg.PublicKeyFilename != tt.publicKeyFilename {
				t.Fatalf("cfg.PublicKeyFilename: %v; publicKeyFilename: %v", cfg.PublicKeyFilename, tt.publicKeyFilename)
			}
		})
	}
}

func setCustomEnvs(t *testing.T) {
	t.Setenv("ADDRESS_AND_PORT", ":8085")
	t.Setenv("CRON_INTERVAL_DAYS", "10")
	t.Setenv("READ_HEADERS_TIMEOUT", "55s")
	t.Setenv("REQUEST_TIMEOUT", "55s")
	t.Setenv("SHUTDOWN_TIMEOUT", "55s")
	t.Setenv("PORKBUN_CERTS_URI", "https://porkbun.com/api/json/v3/ssl/retrieve/custom")
	t.Setenv("PORKBUN_PING_URI", "https://porkbun.com/api/json/v3/ping/custom")
	t.Setenv("INTERMEDIATE_CERT_FILENAME", "intermediate.crt.pem1")
	t.Setenv("CERT_CHAIN_FILENAME", "cert-chain.crt.pem1")
	t.Setenv("PRIVATE_KEY_FILENAME", "private.key.pem1")
	t.Setenv("PUBLIC_KEY_FILENAME", "public.key.pem1")
}
