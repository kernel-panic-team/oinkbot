package service

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/kernel-panic-team/oinkbot/internal/models"
)

func (s *Service) PingPorkbun() (*models.PorkPingResponse, error) {
	body, err := s.client.GetRawResponse(s.config.PorkbunPingURI)
	if err != nil {
		return nil, err
	}

	var oink models.PorkPingResponse
	err = json.Unmarshal(body, &oink)
	if err != nil {
		return nil, fmt.Errorf("unable to decode response: %v", err)
	}

	return &oink, nil
}

func (s *Service) GetPorkbunCerts() (*models.PorkCerts, error) {
	fullURI, err := url.JoinPath(s.config.PorkbunCertsURI, s.config.DomainName)
	if err != nil {
		return nil, err
	}
	body, err := s.client.GetRawResponse(fullURI)
	if err != nil {
		return nil, err
	}

	var certs models.PorkCerts
	err = json.Unmarshal(body, &certs)
	if err != nil {
		return nil, fmt.Errorf("unable to decode response: %v", err)
	}

	return &certs, nil
}
