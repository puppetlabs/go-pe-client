package pe

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func init() {
	peClient = NewClient(peHostURL, "xxxx", nil)
	peClient.strict = true
	httpmock.Activate()
	httpmock.ActivateNonDefault(peClient.resty.GetClient())
}

func setupGetResponder(t *testing.T, url, query, responseFilename string) {
	httpmock.Reset()
	responseBody, err := ioutil.ReadFile("testdata/apidocs/" + responseFilename)
	require.Nil(t, err)
	response := httpmock.NewBytesResponse(200, responseBody)
	response.Header.Set("Content-Type", "application/json")
	if query != "" {
		httpmock.RegisterResponder(http.MethodGet, peHostURL+url, httpmock.ResponderFromResponse(response))
	} else {
		httpmock.RegisterResponderWithQuery(http.MethodGet, peHostURL+url, query, httpmock.ResponderFromResponse(response))
	}
	response.Body.Close()
}

var peClient *Client

var peHostURL = "https://test-host"
