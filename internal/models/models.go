package models

type PorkRequest struct {
	SecretAPIKey string `json:"secretapikey"`
	APIKey       string `json:"apikey"`
}

type PorkPingResponse struct {
	Status string `json:"status"`
	YourIp string `json:"yourIp"`
}

type PorkCerts struct {
	IntermediateCert string `json:"intermediatecertificate"`
	CertChain        string `json:"certificatechain"`
	PrivateKey       string `json:"privatekey"`
	PublicKey        string `json:"publickey"`
}

type PorkErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (p *PorkCerts) AreEqual(other *PorkCerts) bool {
	if p == nil || other == nil {
		return false
	}
	return p.IntermediateCert == other.IntermediateCert &&
		p.CertChain == other.CertChain &&
		p.PrivateKey == other.PrivateKey &&
		p.PublicKey == other.PublicKey
}
