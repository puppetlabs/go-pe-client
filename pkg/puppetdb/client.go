package puppetdb

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/sirupsen/logrus"

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
//				&OrderBy{Field: "certname", Order: "asc"},
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
		var ue *url.Error
		if errors.As(err, &ue) {
			return fmt.Errorf("%s %s: %w", client.resty.HostURL, path, ue.Err)
		}
		return fmt.Errorf("%s %s: %w", client.resty.HostURL, path, err)
	}
	if r.IsError() {
		re := r.Error()
		if re == nil {
			return fmt.Errorf("%s %s: %s: \"%s\"", client.resty.HostURL, path, r.Status(), r.Body())
		}
		return fmt.Errorf("%s %s: %s: \"%s\": %v", client.resty.HostURL, path, r.Status(), r.Body(), re)
	}

	if pagination != nil && pagination.IncludeTotal {
		pagination.Total = getTotal(r.Header().Get("X-Records"))
	}

	return nil
}

// getTotal extracts the total from the X-Records header
func getTotal(records string) int {
	if records != "" {
		count, err := strconv.Atoi(records)
		if err == nil {
			return count
		}
		logrus.Warnf("Unable to convert X-Records %s to int", records)
	}
	return 0
}
