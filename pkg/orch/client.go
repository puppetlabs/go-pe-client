package orch

import (
	"bytes"
	"crypto/tls"
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

// Client for the Orchestrator API
type Client struct {
	resty  *resty.Client
	strict bool
}

// NewClient access the orchestrator API via TLS
func NewClient(hostURL string, token string, tlsConfig *tls.Config) *Client {
	r := resty.New()
	if tlsConfig != nil {
		r.SetTLSClientConfig(tlsConfig)
	}
	r.SetHostURL(hostURL)
	r.SetHeader("X-Authentication", token)
	r.SetError(OrchestratorError{})
	r.SetRedirectPolicy(resty.NoRedirectPolicy())

	client := Client{resty: r}
	r.JSONUnmarshal = func(data []byte, v interface{}) error {
		d := json.NewDecoder(bytes.NewReader(data))
		if client.strict {
			d.DisallowUnknownFields()
		}
		return d.Decode(v)
	}
	return &client
}

// OrchestratorError represents an error response from the Orchestrator API
type OrchestratorError struct {
	Kind       string `json:"kind"`
	Msg        string `json:"Msg"`
	StatusCode int
}

func (oe *OrchestratorError) Error() string {
	return oe.Msg
}

// GetStatusCode will return the status code.
func (oe *OrchestratorError) GetStatusCode() int {
	return oe.StatusCode
}

// HTTPError represents an error with the HTTP response code
type HTTPError struct {
	Msg        string
	StatusCode int
}

func (he *HTTPError) Error() string {
	return he.Msg
}

// GetStatusCode will return the HTTP status code.
func (he *HTTPError) GetStatusCode() int {
	return he.StatusCode
}
