package puppetdb

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

var (
	// ErrNonTransientResponse is returned if the downstream puppetdb api
	// returns an error response that most likely means a retry of the request
	// will fail. A caller receiving this error should not attempt a retry.
	ErrNonTransientResponse = errors.New("puppetdb: the api response indicates an error that cannot be recovered from")

	// ErrTransientResponse is returned if the downstream puppetdb api returns
	// an error that is transient in nature and a retry of the request is
	// likely to succeed. An example of these could be a gateway timeout error
	// or some kind of temporary reverse proxy issue.
	ErrTransientResponse = errors.New("puppetdb: the api response indicates a recoverable error")
)

// Client for the Orchestrator API
type Client struct {
	resty *resty.Client
}

// NewClient access the orchestrator API via TLS. N.B. The timeout is the resty http client timeout so is all encompassing
// and will incorporate connect, TLS handshake, http/s header receipt and general data transfer. The value used for this in
// ER is 5 seconds which seems reasonable.
func NewClient(hostURL, token string, tlsConfig *tls.Config, timeout time.Duration) *Client {
	r := resty.New()
	if tlsConfig != nil {
		r.SetTLSClientConfig(tlsConfig)
	}
	r.SetHostURL(hostURL)
	r.SetHeader("X-Authentication", token)
	r.SetTimeout(timeout)
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
// var payload *[]Fact
// getRequest(client, "/pdb/query/v4/facts",
//
//	query,
//	&Pagination{Limit: 10, Offset: 20},
//	&OrderBy{Field: "certname", Order: "asc"},
//	&payload)
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
			return fmt.Errorf("%s%s: %w", client.resty.HostURL, path, ue.Err)
		}

		return fmt.Errorf("%s%s: %w", client.resty.HostURL, path, err)
	}

	if r.IsError() {
		var err error

		code := r.StatusCode()
		switch {
		case code >= http.StatusBadRequest || code < http.StatusInternalServerError:
			err = ErrTransientResponse
		case code > http.StatusInternalServerError:
			err = ErrNonTransientResponse
		}

		re := r.Error()
		if re != nil {
			err = fmt.Errorf("client error: %v: %w", re, err)
		}

		return fmt.Errorf("%s%s: %s: \"%s\": %w", client.resty.HostURL, path, r.Status(), r.Body(), err)
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
