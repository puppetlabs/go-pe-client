package classifier

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-resty/resty/v2"
)

// Client for the Orchestrator API
type Client struct {
	resty *resty.Client
}

// NewClient access the orchestrator API via TLS
func NewClient(hostURL, token string, tlsConfig *tls.Config) *Client {
	r := resty.New()
	if tlsConfig != nil {
		r.SetTLSClientConfig(tlsConfig)
	}
	r.SetBaseURL(hostURL)
	r.SetHeader("X-Authentication", token)
	r.SetRedirectPolicy(resty.NoRedirectPolicy())

	return &Client{resty: r}
}

// SetTransport lets the caller overwrite the default transport used by the client.
// This is useful when injecting mock transports for testing purposes.
func (c *Client) SetTransport(tripper http.RoundTripper) {
	c.resty.SetTransport(tripper)
}

// getRequest uses the Given client to make a HTTP GET request to the given path, providing
// the query.  The result of the request is marshalled into the response type. e.g.
// var payload *[]Group
// getRequest(client, "/classifier/v1/groups",
//
//	&Pagination{Limit: 10, Offset: 20},
//	&payload)
func getRequest(client *Client, path string, pagination *Pagination, response interface{}) error {
	req := client.resty.R().SetResult(&response)

	if pagination != nil {
		req.SetQueryParams(pagination.toParams())
	}

	r, err := req.Get(path)
	if err != nil {
		var ue *url.Error
		if errors.As(err, &ue) {
			return fmt.Errorf("%s%s: %w", client.resty.HostURL, path, ue.Err)
		}
		return fmt.Errorf("%s%s: %w", client.resty.HostURL, path, err)
	}
	if r.IsError() {
		re := r.Error()
		if re == nil {
			return fmt.Errorf("%s%s: %s: \"%s\"", client.resty.HostURL, path, r.Status(), r.Body())
		}
		return fmt.Errorf("%s%s: %s: \"%s\": %v", client.resty.HostURL, path, r.Status(), r.Body(), re)
	}

	return nil
}

func postRequest(client *Client, path string, body string, response interface{}) error {
	req := client.resty.R().
		SetResult(&response).
		SetHeader("Content-Type", "application/json").
		SetBody(body)

	r, err := req.Post(path)
	if err != nil {
		var ue *url.Error
		if errors.As(err, &ue) {
			return fmt.Errorf("%s%s: %w", client.resty.HostURL, path, ue.Err)
		}
		return fmt.Errorf("%s%s: %w", client.resty.HostURL, path, err)
	}
	if r.IsError() {
		re := r.Error()
		if re == nil {
			return fmt.Errorf("%s%s: %s: \"%s\"", client.resty.HostURL, path, r.Status(), r.Body())
		}
		return fmt.Errorf("%s%s: %s: \"%s\": %v", client.resty.HostURL, path, r.Status(), r.Body(), re)
	}

	return nil
}

// PostRequest posts a request to the specified uri
func PostRequest(client *Client, uri string) ([]byte, error) {
	r, err := client.resty.R().
		Post(uri)
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		if r.Error() != nil {
			return nil, r.Error().(error)
		}
		return nil, fmt.Errorf("%s error: %s", uri, r.Status())
	}

	return r.Body(), nil
}
