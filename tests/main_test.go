package tests

import (
	"path"
	"testing"

	"github.com/kernel-panic-team/oinkbot/internal/config"
)

type builder struct {
	filename string
	contents string
}

func setCustomCertFilenames(cfg *config.Config, builderMap map[string]builder) {
	cfg.IntermediateCertFilename = builderMap["intermediate"].filename
	cfg.CertChainFilename = builderMap["certchain"].filename
	cfg.PrivateKeyFilename = builderMap["private"].filename
	cfg.PublicKeyFilename = builderMap["public"].filename
}

func getBuilderMap(t *testing.T) map[string]builder {
	tmpDir := t.TempDir()
	return map[string]builder{
		"intermediate": {
			filename: path.Join(tmpDir, "intermediate.crt"),
			contents: "intermediate",
		},
		"certchain": {
			filename: path.Join(tmpDir, "certchain.crt"),
			contents: "certchain",
		},
		"private": {
			filename: path.Join(tmpDir, "private.key"),
			contents: "private",
		},
		"public": {
			filename: path.Join(tmpDir, "public.key"),
			contents: "public",
		},
	}
}
