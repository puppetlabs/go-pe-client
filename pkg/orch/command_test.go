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
		Scope: Scope{
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

func TestCommandScheduleTask(t *testing.T) {

	// Test success
	setupPostResponder(t, "/orchestrator/v1/command/schedule_task", "command-schedule_task-request.json", "command-schedule_task-response.json")
	scheduleTaskRequest := &ScheduleTaskRequest{
		Environment: "test-env-1",
		Task:        "package",
		Params: map[string]string{
			"action":  "install",
			"package": "httpd",
		},
		Scope: Scope{
			Nodes: []string{"node1.example.com"},
		},
		ScheduledTime: "2027-05-05T19:50:08Z",
	}

	actual, err := orchClient.CommandScheduleTask(scheduleTaskRequest)
	require.Nil(t, err)
	require.Equal(t, expectedCommandScheduleTaskResponse, actual)

	// Test error
	setupErrorResponder(t, "/orchestrator/v1/command/schedule_task")
	actual, err = orchClient.CommandScheduleTask(scheduleTaskRequest)
	require.Nil(t, actual)
	require.Equal(t, expectedError, err)
}

var expectedCommandScheduleTaskResponse = &ScheduledJobID{ScheduledJob: struct {
	ID   string "json:\"id\""
	Name string "json:\"name\""
}{ID: "https://orchestrator.example.com:8143/orchestrator/v1/scheduled_jobs/2", Name: "1234"}}

func TestCommandTaskTarget(t *testing.T) {

	// Test success
	setupPostResponder(t, "/orchestrator/v1/command/task_target", "command-task_target-request.json", "command-task_target-response.json")
	taskTargetRequest := &TaskTargetRequest{
		DisplayName: "1",
		NodeGroups:  []string{"3c4df64f-7609-4d31-9c2d-acfa52ed66ec", "4932bfe7-69c4-412f-b15c-ac0a7c2883f1"},
		Nodes:       []string{"wss6c3w9wngpycg", "jjj2h5w8gpycgwn"},
		PQLQuery:    "nodes[certname] { catalog_environment = \"production\" }",
		Tasks:       []string{"package::install", "exec"},
	}

	actual, err := orchClient.CommandTaskTarget(taskTargetRequest)
	require.Nil(t, err)
	require.Equal(t, expectedCommandTaskTargetResponse, actual)

	// Test error
	setupErrorResponder(t, "/orchestrator/v1/command/task_target")
	actual, err = orchClient.CommandTaskTarget(taskTargetRequest)
	require.Nil(t, actual)
	require.Equal(t, expectedError, err)
}

var expectedCommandTaskTargetResponse = &TaskTargetJobID{TaskTargetJob: struct {
	ID   string "json:\"id\""
	Name string "json:\"name\""
}{ID: "https://orchestrator.example.com:8143/orchestrator/v1/scopes/task_targets/1", Name: "1"}}

func TestCommandPlanRun(t *testing.T) {

	// Test success
	setupPostResponder(t, "/orchestrator/v1/command/plan_run", "command-plan_run-request.json", "command-plan_run-response.json")
	planRunRequest := &PlanRunRequest{
		Description: "Start the canary plan on node1 and node2",
		Params: map[string]interface{}{
			"canary":  1,
			"command": "whoami",
			"nodes":   []string{"node1.example.com", "node2.example.com"},
		},
		Name: "canary"}

	actual, err := orchClient.CommandPlanRun(planRunRequest)
	require.Nil(t, err)
	require.Equal(t, expectedCommandPlanRunResponse, actual)

	// Test error
	setupErrorResponder(t, "/orchestrator/v1/command/plan_run")
	actual, err = orchClient.CommandPlanRun(planRunRequest)
	require.Nil(t, actual)
	require.Equal(t, expectedError, err)
}

var expectedCommandPlanRunResponse = &PlanRunJobID{
	Name: "1234",
}
