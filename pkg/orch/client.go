package orch

import (
	"crypto/tls"

	"github.com/go-resty/resty/v2"
)

// Client for the Orchestrator API
type Client struct {
	resty *resty.Client
}

// NewInsecureClient access the orchestrator API in an insecure manner
func NewInsecureClient(hostURL, token string) *Client {
	r := resty.New()
	r.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	r.SetHostURL(hostURL)
	r.SetHeader("X-Authentication", token)
	r.SetError(OrchestratorError{})
	return &Client{
		resty: r,
	}
}

// OrchestratorError represents an error response from the Orchestrator API
type OrchestratorError struct {
	Kind string `json:"kind"`
	Msg  string `json:"msg"`
}

func (oe *OrchestratorError) Error() string {
	return oe.Msg
}
