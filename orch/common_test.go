package orch

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func init() {
	client = NewInsecureClient(hostURL, "xxxx")
	httpmock.Activate()
	httpmock.ActivateNonDefault(client.resty.GetClient())
}

func setupGetResponder(t *testing.T, url, query, responseFilename string) {
	httpmock.Reset()
	responseBody, err := ioutil.ReadFile("testdata/apidocs/" + responseFilename)
	require.Nil(t, err)
	response := httpmock.NewBytesResponse(200, responseBody)
	response.Header.Set("Content-Type", "application/json")
	if query != "" {
		httpmock.RegisterResponder(http.MethodGet, hostURL+url, httpmock.ResponderFromResponse(response))
	} else {
		httpmock.RegisterResponderWithQuery(http.MethodGet, hostURL+url, query, httpmock.ResponderFromResponse(response))
	}
}

func setupPostResponder(t *testing.T, url, requestFilename, responseFilename string) {
	httpmock.Reset()
	httpmock.RegisterResponder(http.MethodPost, hostURL+url,
		func(req *http.Request) (*http.Response, error) {

			// Validate the body
			actual := map[string]interface{}{}
			err := json.NewDecoder(req.Body).Decode(&actual)
			require.Nil(t, err, "error decoding actual body for "+url)
			expected := map[string]interface{}{}
			f, err := os.Open("testdata/apidocs/" + requestFilename)
			require.Nil(t, err, "error reading expected body: testdata/apidocs/"+requestFilename)
			err = json.NewDecoder(f).Decode(&expected)
			require.Nil(t, err, "error decoding expected body for "+url)
			require.Equal(t, expected, actual)

			// Build response
			responseBody, err := ioutil.ReadFile("testdata/apidocs/" + responseFilename)
			require.Nil(t, err)
			response := httpmock.NewBytesResponse(200, responseBody)
			response.Header.Set("Content-Type", "application/json")
			return response, nil

		},
	)
}

func setupErrorResponder(t *testing.T, url string) {
	httpmock.Reset()
	responder, err := httpmock.NewJsonResponder(400, expectedError)
	require.Nil(t, err)
	httpmock.RegisterResponder(http.MethodGet, hostURL+url, responder)
	httpmock.RegisterResponder(http.MethodPost, hostURL+url, responder)
}

var client *Client

var hostURL = "https://test-host:8143"

var expectedError = &OrchestratorError{
	Kind: "puppetlabs.orchestrator/unknown-environment",
	Msg:  "Unknown environment doesnotexist",
}
