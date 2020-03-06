package orch

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

var client *Client
var hostURL = "https://test-host:8143"

func init() {
	client = NewInsecureClient(hostURL, "xxxx")
	httpmock.Activate()
	httpmock.ActivateNonDefault(client.resty.GetClient())
}

func TestTasks(t *testing.T) {
	defer httpmock.Reset()

	// Mock the response
	data, err := ioutil.ReadFile("testdata/tasks.json")
	require.Nil(t, err)
	expected := &Tasks{}
	err = json.Unmarshal(data, expected)
	require.Nil(t, err)
	responder, err := httpmock.NewJsonResponder(200, expected)
	require.Nil(t, err)
	httpmock.RegisterResponder("GET", hostURL+"/orchestrator/v1/tasks", responder)

	// Test the call
	tasks, err := client.Tasks()
	require.Nil(t, err)
	require.NotEmpty(t, tasks.Environment.Name)
	require.NotEmpty(t, tasks.Environment.CodeID)
	require.Len(t, tasks.Items, 12)
	require.NotEmpty(t, tasks.Items[0].ID)
	require.NotEmpty(t, tasks.Items[0].Name)
	require.True(t, tasks.Items[0].Permitted)
}

func TestTasksWithError(t *testing.T) {
	defer httpmock.Reset()

	// Mock the response
	data, err := ioutil.ReadFile("testdata/error.json")
	require.Nil(t, err)
	expected := &OrchestratorError{}
	err = json.Unmarshal(data, expected)
	require.Nil(t, err)
	responder, err := httpmock.NewJsonResponder(400, expected)
	require.Nil(t, err)
	httpmock.RegisterResponder("GET", hostURL+"/orchestrator/v1/tasks", responder)

	// Test the call
	tasks, err := client.Tasks()
	require.Nil(t, tasks)
	require.NotNil(t, err)
	require.NotEmpty(t, err.Error())
	require.NotEmpty(t, err.(*OrchestratorError).Kind)
	require.NotEmpty(t, err.(*OrchestratorError).Msg)
}
