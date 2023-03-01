package rbac

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
	rbacClient = NewClient(rbacHostURL, nil)
	rbacClient.strict = true
	httpmock.Activate()
	httpmock.ActivateNonDefault(rbacClient.resty.GetClient())
}

func setupPostResponder(t *testing.T, url, requestFilename, responseFilename string) {
	httpmock.Reset()
	httpmock.RegisterResponder(http.MethodPost, rbacHostURL+url,
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
	httpmock.RegisterResponder(http.MethodGet, rbacHostURL+url, responder)
	httpmock.RegisterResponder(http.MethodPost, rbacHostURL+url, responder)
}

func setupCreateRoleSuccessResponder(t *testing.T, url string, requestFilename string) {
	httpmock.Reset()
	httpmock.RegisterResponder(http.MethodPost, rbacHostURL+url,
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
			response := httpmock.NewBytesResponse(303, []byte{})
			response.Header.Set("Content-Type", "application/json")
			response.Header.Set("Location", "/path/to/role")
			return response, nil
		},
	)
}

func setupCreateRoleErrorResponder(t *testing.T, url string) {
	httpmock.Reset()
	responder, err := httpmock.NewJsonResponder(409, createRoleDuplicateError)
	require.Nil(t, err)
	httpmock.RegisterResponder(http.MethodGet, rbacHostURL+url, responder)
	httpmock.RegisterResponder(http.MethodPost, rbacHostURL+url, responder)
}

var rbacClient *Client

var rbacHostURL = "https://test-host:4433"

var expectedError = &APIError{
	Kind:       "puppetlabs.rbac/unknown-environment",
	Msg:        "Unknown environment doesnotexist",
	StatusCode: 400,
}

var createRoleDuplicateError = &APIError{
	Msg:        "There was a database conflict due to the value(s): Testing",
	StatusCode: 409,
}
