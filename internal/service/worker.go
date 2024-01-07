package service

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"time"

	"github.com/kernel-panic-team/oinkbot/internal/models"
)

func (s *Service) worker() {
	oldCerts, err := s.ReadOldCerts()
	if err != nil {
		log.Println(err)
	}

	newCerts, err := s.GetPorkbunCerts()
	if err != nil {
		log.Println(err)
	}

	if newCerts.AreEqual(oldCerts) {
		log.Println("certs are up to date")
		return
	}

	err = s.BackupOldCerts(time.Now().Format("2006-01-02_15-04-05"))
	if err != nil {
		log.Println(err)
	}

	err = s.SaveNewCerts(newCerts)
	if err != nil {
		log.Println(err)
	}
}

func (s *Service) ReadOldCerts() (*models.PorkCerts, error) {
	var err error
	var oldCerts models.PorkCerts

	oldCerts.IntermediateCert, err = readCert(s.config.IntermediateCertFilename)
	if err != nil {
		return nil, err
	}
	oldCerts.CertChain, err = readCert(s.config.CertChainFilename)
	if err != nil {
		return nil, err
	}
	oldCerts.PrivateKey, err = readCert(s.config.PrivateKeyFilename)
	if err != nil {
		return nil, err
	}
	oldCerts.PublicKey, err = readCert(s.config.PublicKeyFilename)
	if err != nil {
		return nil, err
	}

	return &oldCerts, err
}

func readCert(filename string) (string, error) {
	raw, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("readCert error: %w", err)
	}
	return string(raw), nil
}

func (s *Service) SaveNewCerts(newCerts *models.PorkCerts) error {
	err := os.WriteFile(s.config.IntermediateCertFilename, []byte(newCerts.IntermediateCert), fs.ModePerm)
	if err != nil {
		return err
	}
	err = os.WriteFile(s.config.CertChainFilename, []byte(newCerts.CertChain), fs.ModePerm)
	if err != nil {
		return err
	}
	err = os.WriteFile(s.config.PrivateKeyFilename, []byte(newCerts.PrivateKey), fs.ModePerm)
	if err != nil {
		return err
	}
	err = os.WriteFile(s.config.PublicKeyFilename, []byte(newCerts.PublicKey), fs.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) BackupOldCerts(suffix string) error {
	err := os.Rename(s.config.IntermediateCertFilename, s.config.IntermediateCertFilename+"."+suffix)
	if err != nil {
		return err
	}
	err = os.Rename(s.config.CertChainFilename, s.config.CertChainFilename+"."+suffix)
	if err != nil {
		return err
	}
	err = os.Rename(s.config.PrivateKeyFilename, s.config.PrivateKeyFilename+"."+suffix)
	if err != nil {
		return err
	}
	err = os.Rename(s.config.PublicKeyFilename, s.config.PublicKeyFilename+"."+suffix)
	if err != nil {
		return err
	}
	return nil
}
