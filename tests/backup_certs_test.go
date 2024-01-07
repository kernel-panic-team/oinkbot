package tests

import (
	"os"
	"testing"
	"time"

	"github.com/kernel-panic-team/oinkbot/internal/config"
	"github.com/kernel-panic-team/oinkbot/internal/service"
)

func TestBackupCerts(t *testing.T) {
	cfg := config.New()
	srv := service.New(cfg, nil, nil, nil)

	builderMap := getBuilderMap(t)

	setCustomCertFilenames(cfg, builderMap)

	for _, b := range builderMap {
		if err := os.WriteFile(b.filename, []byte(b.contents), 0644); err != nil {
			t.Fatalf("unable to write %s: %v", b.filename, err)
		}
	}

	suffix := time.Now().Format("2006-01-02_15-04-05")

	err := srv.BackupOldCerts(suffix)
	if err != nil {
		t.Fatalf("unable to rename old certs: %v", err)
	}

	for _, b := range builderMap {
		if _, err := os.Stat(b.filename); err == nil {
			t.Fatalf("%s still exists", b.filename)
		}
		if _, err := os.Stat(b.filename + "." + suffix); err != nil {
			t.Fatalf("%s.%s does not exist", b.filename, suffix)
		}
		if contents, err := os.ReadFile(b.filename + "." + suffix); err != nil {
			t.Fatalf("unable to read %s.%s: %v", b.filename, suffix, err)
		} else if string(contents) != b.contents {
			t.Fatalf("unexpected contents of %s.%s: %s", b.filename, suffix, contents)
		}
	}
}
