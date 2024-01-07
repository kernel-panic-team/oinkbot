package tests

import (
	"os"
	"testing"

	"github.com/kernel-panic-team/oinkbot/internal/config"
	"github.com/kernel-panic-team/oinkbot/internal/service"
)

func TestReadCerts(t *testing.T) {
	cfg := config.New()
	srv := service.New(cfg, nil, nil, nil)

	builderMap := getBuilderMap(t)

	setCustomCertFilenames(cfg, builderMap)

	for _, b := range builderMap {
		if err := os.WriteFile(b.filename, []byte(b.contents), 0644); err != nil {
			t.Fatalf("unable to write %s: %v", b.filename, err)
		}
	}

	oldCerts, err := srv.ReadOldCerts()
	if err != nil {
		t.Fatalf("unable to read old certs: %v", err)
	}

	if oldCerts.IntermediateCert != builderMap["intermediate"].contents {
		t.Fatalf("expected %s, got %s", builderMap["intermediate"].contents, oldCerts.IntermediateCert)
	}
	if oldCerts.CertChain != builderMap["certchain"].contents {
		t.Fatalf("expected %s, got %s", builderMap["certchain"].contents, oldCerts.CertChain)
	}
	if oldCerts.PrivateKey != builderMap["private"].contents {
		t.Fatalf("expected %s, got %s", builderMap["private"].contents, oldCerts.PrivateKey)
	}
	if oldCerts.PublicKey != builderMap["public"].contents {
		t.Fatalf("expected %s, got %s", builderMap["public"].contents, oldCerts.PublicKey)
	}
}
