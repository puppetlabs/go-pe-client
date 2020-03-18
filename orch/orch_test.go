package orch

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func TestInventory(t *testing.T) {
	expected := setupResponder(t, "inventory")
	actual, err := client.Inventory()
	require.Nil(t, err)
	require.Equal(t, expected, normalize(t, actual))
	setupErrorResponder(t, "inventory")
	actual, err = client.Inventory()
	require.Nil(t, actual)
	require.Equal(t, expectedError, err)
}

func TestTasks(t *testing.T) {
	expected := setupResponder(t, "tasks")
	actual, err := client.Tasks()
	require.Nil(t, err)
	require.Equal(t, expected, normalize(t, actual))
	setupErrorResponder(t, "tasks")
	actual, err = client.Tasks()
	require.Nil(t, actual)
	require.Equal(t, expectedError, err)
}

func init() {
	client = NewInsecureClient(hostURL, "xxxx")
	httpmock.Activate()
	httpmock.ActivateNonDefault(client.resty.GetClient())
}

func setupResponder(t *testing.T, url string) *map[string]interface{} {
	httpmock.Reset()
	expectedEncoded, err := ioutil.ReadFile("testdata/apidocs/" + url + ".json")
	require.Nil(t, err)
	expected := &map[string]interface{}{}
	err = json.Unmarshal(expectedEncoded, expected)
	require.Nil(t, err)
	responder, err := httpmock.NewJsonResponder(200, expected)
	require.Nil(t, err)
	httpmock.RegisterResponder("GET", hostURL+"/orchestrator/v1/"+url, responder)
	return expected
}

func setupErrorResponder(t *testing.T, url string) {
	httpmock.Reset()
	data, err := ioutil.ReadFile("testdata/apidocs/error.json")
	require.Nil(t, err)
	expected := &OrchestratorError{}
	err = json.Unmarshal(data, expected)
	require.Nil(t, err)
	responder, err := httpmock.NewJsonResponder(400, expected)
	require.Nil(t, err)
	httpmock.RegisterResponder("GET", hostURL+"/orchestrator/v1/"+url, responder)
}

func normalize(t *testing.T, actual interface{}) *map[string]interface{} {
	// Normalize the actual result into a map via JSON encoding to allow reliable comparison
	actualEncoded, err := json.Marshal(actual)
	require.Nil(t, err)
	actualNormalized := &map[string]interface{}{}
	err = json.Unmarshal(actualEncoded, actualNormalized)
	require.Nil(t, err)
	return actualNormalized
}

var client *Client

var hostURL = "https://test-host:8143"

var expectedError = &OrchestratorError{
	Kind: "puppetlabs.orchestrator/unknown-environment",
	Msg:  "Unknown environment doesnotexist",
}
