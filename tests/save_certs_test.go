package tests

import (
	"os"
	"testing"

	"github.com/kernel-panic-team/oinkbot/internal/config"
	"github.com/kernel-panic-team/oinkbot/internal/models"
	"github.com/kernel-panic-team/oinkbot/internal/service"
)

func TestWriteCerts(t *testing.T) {
	cfg := config.New()
	srv := service.New(cfg, nil, nil, nil)

	builderMap := getBuilderMap(t)

	setCustomCertFilenames(cfg, builderMap)

	newCerts := models.PorkCerts{
		IntermediateCert: builderMap["intermediate"].contents,
		CertChain:        builderMap["certchain"].contents,
		PrivateKey:       builderMap["private"].contents,
		PublicKey:        builderMap["public"].contents,
	}

	err := srv.SaveNewCerts(&newCerts)
	if err != nil {
		t.Fatalf("SaveNewCerts returned error: %v, want nil", err)
	}

	for _, b := range builderMap {
		if _, err := os.Stat(b.filename); err != nil {
			t.Fatalf("SaveNewCerts did not create %s", b.filename)
		}
		contents, err := os.ReadFile(b.filename)
		if err != nil {
			t.Fatalf("Can't read saved file %s: error: %v", b.filename, err)
		}
		if b.contents != string(contents) {
			t.Fatalf("SaveNewCerts wrote wrong contents to %s", b.filename)
		}
	}
}
