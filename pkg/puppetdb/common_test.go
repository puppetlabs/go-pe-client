package puppetdb

import (
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func init() {
	pdbClient = NewClient(hostURL, "xxxx", nil, time.Second*1000)
	httpmock.Activate()
	httpmock.ActivateNonDefault(pdbClient.resty.GetClient())
}

func setupGetResponder(t *testing.T, url, query, responseFilename string) {
	httpmock.Reset()
	responseBody, err := os.ReadFile("testdata/" + responseFilename)
	require.Nil(t, err)
	response := httpmock.NewBytesResponse(200, responseBody)
	response.Header.Set("Content-Type", "application/json")
	if query != "" {
		httpmock.RegisterResponder(http.MethodGet, hostURL+url, httpmock.ResponderFromResponse(response))
	} else {
		httpmock.RegisterResponderWithQuery(http.MethodGet, hostURL+url, query, httpmock.ResponderFromResponse(response))
	}
	response.Body.Close()
}

func setupURLErrorResponder(t *testing.T, url string) {
	setupURLResponderWithStatusCode(t, url, http.StatusNotFound)
}

func setupURLResponderWithStatusCode(t *testing.T, url string, statusCode int) {
	setupResponderWithStatusCodeAndBody(t, url, statusCode, expectedURLError)
}

func setupResponderWithStatusCodeAndBody(t *testing.T, url string, statusCode int, response interface{}) {
	httpmock.Reset()
	responder, err := httpmock.NewJsonResponder(statusCode, response)
	require.Nil(t, err)
	httpmock.RegisterResponder(http.MethodGet, hostURL+url, responder)
}

var pdbClient *Client

var hostURL = "https://test-host:8081"

var expectedURLError = url.Error{Op: "nil", URL: hostURL, Err: nil}
