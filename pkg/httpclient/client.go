package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/kernel-panic-team/oinkbot/internal/config"
	"github.com/kernel-panic-team/oinkbot/internal/models"
)

type Client struct {
	cfg *config.Config
	cli *http.Client
}

func New(cfg *config.Config) *Client {
	return &Client{
		cfg: cfg,
		cli: &http.Client{
			Timeout: cfg.RequestTimeout,
		},
	}
}

func (c *Client) GetRawResponse(URI string) (body []byte, err error) {
	payload, err := json.Marshal(models.PorkRequest{
		SecretAPIKey: c.cfg.SecretAPIKey,
		APIKey:       c.cfg.APIKey,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to marshal payload: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, URI, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := c.cli.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to send request: %v", err)
	}
	defer res.Body.Close()

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error returned: %s: %s", res.Status, decodeErrorResponse(body))
	}

	return body, nil
}

func decodeErrorResponse(body []byte) string {
	var errRes models.PorkErrorResponse
	err := json.Unmarshal(body, &errRes)
	if err != nil {
		return fmt.Sprintf("unable to decode error response: %v", err)
	}
	return errRes.Message
}
