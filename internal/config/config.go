package config

import (
	"log"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/caarlos0/env/v10"
)

type Days time.Duration

func (d *Days) UnmarshalText(text []byte) error {
	days, err := strconv.Atoi(string(text))
	if err != nil {
		return err
	}
	*d = Days(time.Duration(days) * 24 * time.Hour)
	return nil
}

type Config struct {
	wd                       string
	DomainName               string        `env:"DOMAIN_NAME"`
	AddressAndPort           string        `env:"ADDRESS_AND_PORT" envDefault:":8080"`
	CronInterval             Days          `env:"CRON_INTERVAL_DAYS" envDefault:"5"`
	ReadHeadersTimeout       time.Duration `env:"READ_HEADERS_TIMEOUT" envDefault:"5s"`
	RequestTimeout           time.Duration `env:"REQUEST_TIMEOUT" envDefault:"5s"`
	ShutdownTimeout          time.Duration `env:"SHUTDOWN_TIMEOUT" envDefault:"5s"`
	SecretAPIKey             string        `env:"PORKBUN_SECRET_KEY" unset:"true"`
	APIKey                   string        `env:"PORKBUN_API_KEY" unset:"true"`
	PorkbunCertsURI          string        `env:"PORKBUN_CERTS_URI" envDefault:"https://porkbun.com/api/json/v3/ssl/retrieve"`
	PorkbunPingURI           string        `env:"PORKBUN_PING_URI" envDefault:"https://porkbun.com/api/json/v3/ping"`
	IntermediateCertFilename string        `env:"INTERMEDIATE_CERT_FILENAME" envDefault:"intermediate.crt"`
	CertChainFilename        string        `env:"CERT_CHAIN_FILENAME" envDefault:"cert-chain.crt"`
	PrivateKeyFilename       string        `env:"PRIVATE_KEY_FILENAME" envDefault:"private.key"`
	PublicKeyFilename        string        `env:"PUBLIC_KEY_FILENAME" envDefault:"public.key"`
}

func New() *Config {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("unable to parse config: %v", err)
	}
	wd, _ := os.Getwd()
	cfg.wd = wd

	cfg.IntermediateCertFilename = path.Join(cfg.wd, "ssl", cfg.IntermediateCertFilename)
	cfg.CertChainFilename = path.Join(cfg.wd, "ssl", cfg.CertChainFilename)
	cfg.PrivateKeyFilename = path.Join(cfg.wd, "ssl", cfg.PrivateKeyFilename)
	cfg.PublicKeyFilename = path.Join(cfg.wd, "ssl", cfg.PublicKeyFilename)

	return &cfg
}
