package classifier

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func init() {
	pdbClient = NewClient(hostURL, "xxxx", nil)
	httpmock.Activate()
	httpmock.ActivateNonDefault(pdbClient.resty.GetClient())
}

func TestGroups(t *testing.T) {
	setupGetResponder(t, groups, "", "groups-response.json")
	actual, err := pdbClient.Groups(nil)
	require.Nil(t, err)
	require.True(t, len(actual) == 2)
	require.Equal(t, expected, actual)
}

func TestGroup(t *testing.T) {
	setupGetResponder(t, fmt.Sprintf("%s/%s", groups, g2ID), "", "group.json")
	actual, err := pdbClient.Group(g2ID)
	require.Nil(t, err)
	require.NotNil(t, actual)
	require.Equal(t, g2, actual)
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

var (
	pdbClient *Client
	hostURL   = "https://test-host:4433"
)

var g = Group{
	Parent:            "5239b6da-8194-4f99-ab37-de6cbf665754",
	EnvironmentTrumps: false,
	Name:              "PE Infrastructure Agent",
	Rule:              []interface{}{"and", []interface{}{"~", []interface{}{"fact", "pe_server_version"}, ".+"}},
	Variables:         map[string]interface{}{},
	ID:                "799ab2fb-5ece-4628-ac47-ba61e70d5b54",
	Environment:       "production",
	LastEdited:        time.Time{},
	SerialNumber:      1,
	Classes:           map[string]interface{}{"puppet_enterprise::profile::agent": map[string]interface{}{}},
	ConfigData:        map[string]interface{}{},
}

var (
	g2ID = "913e54b7-b09c-4543-9f74-daff5e51a49f"
	g2   = Group{
		Parent:            "d1fc626d-222a-4616-be89-cc67ef1f8b3a",
		Description:       "Production nodes",
		EnvironmentTrumps: true,
		Name:              "Production environment",
		Rule:              []interface{}{"and", []interface{}{"=", []interface{}{"trusted", "extensions", "pp_environment"}, "production"}},
		Variables:         map[string]interface{}{},
		ID:                g2ID,
		Environment:       "production",
		LastEdited:        time.Time{},
		SerialNumber:      1,
		Classes:           map[string]interface{}{},
		ConfigData:        map[string]interface{}{},
		Deleted:           map[string]interface{}(nil),
	}
)

var expected = []Group{
	g, g2,
}
