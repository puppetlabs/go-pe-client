package pe

import (
	"bytes"
	"crypto/tls"
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

// Client for the PE API
type Client struct {
	resty  *resty.Client
	strict bool
}

// NewClient access the PE API via TLS
func NewClient(hostURL string, token string, tlsConfig *tls.Config) *Client {
	r := resty.New()
	if tlsConfig != nil {
		r.SetTLSClientConfig(tlsConfig)
	}
	r.SetBaseURL(hostURL)
	r.SetHeader("X-Authentication", token)
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
