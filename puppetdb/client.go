package puppetdb

import (
	"crypto/tls"
	"net/http"

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

	return &Client{
		resty: r,
	}
}

// SetTransport lets the caller overwrite the default transport used by the client.
// This is useful when injecting mock transports for testing purposes.
func (c *Client) SetTransport(tripper http.RoundTripper) {
	c.resty.SetTransport(tripper)
}
