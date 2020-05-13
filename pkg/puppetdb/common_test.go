package puppetdb

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func init() {
	pdbClient = NewClient(hostURL, "xxxx", nil)
	httpmock.Activate()
	httpmock.ActivateNonDefault(pdbClient.resty.GetClient())
}

func setupGetResponder(t *testing.T, url, query, responseFilename string) {
	httpmock.Reset()
	responseBody, err := ioutil.ReadFile("testdata/" + responseFilename)
	require.Nil(t, err)
	response := httpmock.NewBytesResponse(200, responseBody)
	response.Header.Set("Content-Type", "application/json")
	if query != "" {
		httpmock.RegisterResponder(http.MethodGet, hostURL+url, httpmock.ResponderFromResponse(response))
	} else {
		httpmock.RegisterResponderWithQuery(http.MethodGet, hostURL+url, query, httpmock.ResponderFromResponse(response))
	}
}

var pdbClient *Client

var hostURL = "https://test-host:8081"
