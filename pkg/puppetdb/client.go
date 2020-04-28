package puppetdb

import (
	"crypto/tls"
	"fmt"
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

// getRequest uses the Given client to make a HTTP GET request to the given path, providing
// the query.  The result of the request is marshalled into the response type. e.g.
// var payload *[]Fact
// getRequest(client, "/pdb/query/v4/facts",
//				query,
//				&Pagination{Limit: 10, Offset: 20},
//				&OrderBy{Field: "certname", Order: "asc",},
//				&payload)
func getRequest(client *Client, path string, query string, pagination *Pagination, orderBy *OrderBy, response interface{}) error {
	req := client.resty.R().SetResult(&response)
	if query != "" {
		req.SetQueryParam("query", query)
	}
	if pagination != nil {
		req.SetQueryParams(pagination.toParams())
	}
	if orderBy != nil {
		req.SetQueryParams(orderBy.toParams())
	}
	r, err := req.Get(path)
	if err != nil {
		return err
	}
	if r.IsError() {
		return fmt.Errorf("%s: %s", r.Status(), r.Body())
	}

	return nil
}
