package orch

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCommandTask(t *testing.T) {

	// Test success
	setupPostResponder(t, "/orchestrator/v1/command/task", "command-task-request.json", "command-task-response.json")
	taskRequest := &TaskRequest{
		Environment: "test-env-1",
		Task:        "package",
		Params: map[string]string{
			"action": "install",
			"name":   "httpd",
		},
		Scope: TaskScope{
			Application: "Wordpress_app[demo]",
			Nodes:       []string{"node1.example.com"},
			Query:       []interface{}{"from", "nodes", []interface{}{"~", "certname", ".*"}},
			NodeGroup:   "00000000-0000-4000-8000-000000000000",
		},
	}

	actual, err := orchClient.CommandTask(taskRequest)
	require.Nil(t, err)
	require.Equal(t, expectedCommandTaskResponse, actual)

	// Test error
	setupErrorResponder(t, "/orchestrator/v1/command/task")
	actual, err = orchClient.CommandTask(taskRequest)
	require.Nil(t, actual)
	require.Equal(t, expectedError, err)

}

var expectedCommandTaskResponse = &JobID{Job: struct {
	ID   string "json:\"id\""
	Name string "json:\"name\""
}{ID: "https://orchestrator.example.com:8143/orchestrator/v1/jobs/1234", Name: "1234"}}
