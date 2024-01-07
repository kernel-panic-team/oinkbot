package tests

import (
	"testing"

	"github.com/kernel-panic-team/oinkbot/internal/models"
)

func TestCertsEquality(t *testing.T) {
	tests := []struct {
		name     string
		want     bool
		oldCerts *models.PorkCerts
		newCerts *models.PorkCerts
	}{
		{
			name: "testing equal certs",
			want: true,
			oldCerts: &models.PorkCerts{
				IntermediateCert: "test_intermediate_cert",
				CertChain:        "test_cert_chain",
				PublicKey:        "test_public_key",
				PrivateKey:       "test_private_key",
			},
			newCerts: &models.PorkCerts{
				IntermediateCert: "test_intermediate_cert",
				CertChain:        "test_cert_chain",
				PublicKey:        "test_public_key",
				PrivateKey:       "test_private_key",
			},
		},
		{
			name: "testing not equal certs",
			want: false,
			oldCerts: &models.PorkCerts{
				IntermediateCert: "test_intermediate_cert",
				CertChain:        "test_cert_chain",
				PublicKey:        "test_public_key",
				PrivateKey:       "test_private_key",
			},
			newCerts: &models.PorkCerts{
				IntermediateCert: "test_intermediate_cert_1",
				CertChain:        "test_cert_chain",
				PublicKey:        "test_public_key_1",
				PrivateKey:       "test_private_key",
			},
		},
		{
			name:     "testing old certs are nil",
			want:     false,
			oldCerts: nil,
			newCerts: &models.PorkCerts{
				IntermediateCert: "test_intermediate_cert",
				CertChain:        "test_cert_chain",
				PublicKey:        "test_public_key",
				PrivateKey:       "test_private_key",
			},
		},
		{
			name: "testing new certs are nil",
			want: false,
			oldCerts: &models.PorkCerts{
				IntermediateCert: "test_intermediate_cert",
				CertChain:        "test_cert_chain",
				PublicKey:        "test_public_key",
				PrivateKey:       "test_private_key",
			},
			newCerts: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.oldCerts.AreEqual(tt.newCerts); got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}
