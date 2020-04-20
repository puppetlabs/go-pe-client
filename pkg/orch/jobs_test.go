package orch

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJobs(t *testing.T) {

	// Test success
	setupGetResponder(t, jobs, "", "jobs-response.json")
	actual, err := orchClient.Jobs()
	require.Nil(t, err)
	require.Equal(t, expectedJobs, actual)

	// Test error
	setupErrorResponder(t, jobs)
	actual, err = orchClient.Jobs()
	require.Nil(t, actual)
	require.Equal(t, expectedError, err)

}

func TestJob(t *testing.T) {

	testURL := strings.ReplaceAll(job, "{job-id}", "123")

	// Test success
	setupGetResponder(t, testURL, "", "job-response.json")
	actual, err := orchClient.Job("123")
	require.Nil(t, err)
	require.Equal(t, expectedJob, actual)

	// Test error
	setupErrorResponder(t, testURL)
	actual, err = orchClient.Job("123")
	require.Nil(t, actual)
	require.Equal(t, expectedError, err)

}

func TestJobReport(t *testing.T) {

	testURL := strings.ReplaceAll(jobReport, "{job-id}", "123")

	// Test success
	setupGetResponder(t, testURL, "", "job-report-response.json")
	actual, err := orchClient.JobReport("123")
	require.Nil(t, err)
	require.Equal(t, expectedJobReport, actual)

	// Test error
	setupErrorResponder(t, testURL)
	actual, err = orchClient.JobReport("123")
	require.Nil(t, actual)
	require.Equal(t, expectedError, err)

}

func TestJobNodes(t *testing.T) {

	testURL := strings.ReplaceAll(jobNodes, "{job-id}", "123")

	// Test success
	setupGetResponder(t, testURL, "", "jobs-nodes-response.json")
	actual, err := orchClient.JobNodes("123")
	require.Nil(t, err)
	require.Equal(t, expectedJobNodes, actual)

	// Test error
	setupErrorResponder(t, testURL)
	actual, err = orchClient.JobNodes("123")
	require.Nil(t, actual)
	require.Equal(t, expectedError, err)

}

var expectedJobs = &Jobs{Items: []Job{Job{ID: "https://orchestrator.example.com:8143/orchestrator/v1/jobs/1234", Name: "1234", Command: "deploy", Options: Options{Concurrency: interface{}(nil), Noop: false, Trace: false, Debug: false, Scope: Scope{Application: "", Nodes: []string(nil), Query: []interface{}(nil), NodeGroup: ""}, EnforceEnvironment: true, Environment: "production", Evaltrace: false, Target: interface{}(nil), Description: "deploy the web app"}, NodeCount: 5, Owner: Owner{ID: "751a8f7e-b53a-4ccd-9f4f-e93db6aa38ec", Login: "brian"}, Description: "deploy the web app", Timestamp: "2016-05-20T16:45:31Z", Environment: Environment{Name: "production"}, Status: []Status(nil), Nodes: Nodes{ID: "https://localhost:8143/orchestrator/v1/jobs/375/nodes"}, Report: Report{ID: "https://localhost:8143/orchestrator/v1/jobs/375/report"}}}, Pagination: Pagination{Limit: 20, Offset: 0, Total: 42}}

var expectedJob = &Job{ID: "https://orchestrator.example.com:8143/orchestrator/v1/jobs/1234", Name: "1234", Command: "deploy", Options: Options{Concurrency: interface{}(nil), Noop: false, Trace: false, Debug: false, Scope: Scope{Application: "Wordpress_app", Nodes: []string(nil), Query: []interface{}(nil), NodeGroup: ""}, EnforceEnvironment: true, Environment: "production", Evaltrace: false, Target: interface{}(nil), Description: ""}, NodeCount: 5, Owner: Owner{ID: "751a8f7e-b53a-4ccd-9f4f-e93db6aa38ec", Login: "admin"}, Description: "deploy the web app", Timestamp: "2016-05-20T16:45:31Z", Environment: Environment{Name: "production"}, Status: []Status{Status{State: "new", EnterTime: "2016-04-11T18:44:31Z", ExitTime: "2016-04-11T18:44:31Z"}, Status{State: "ready", EnterTime: "2016-04-11T18:44:31Z", ExitTime: "2016-04-11T18:44:31Z"}, Status{State: "running", EnterTime: "2016-04-11T18:44:31Z", ExitTime: "2016-04-11T18:45:31Z"}, Status{State: "finished", EnterTime: "2016-04-11T18:45:31Z", ExitTime: ""}}, Nodes: Nodes{ID: "https://orchestrator.example.com:8143/orchestrator/v1/jobs/1234/nodes"}, Report: Report{ID: "https://orchestrator.example.com:8143/orchestrator/v1/jobs/1234/report"}}

var expectedJobReport = &JobReport{Items: []struct {
	Node      string     "json:\"node\""
	State     string     "json:\"state\""
	Timestamp string     "json:\"timestamp\""
	Events    []JobEvent "json:\"events\""
}{struct {
	Node      string     "json:\"node\""
	State     string     "json:\"state\""
	Timestamp string     "json:\"timestamp\""
	Events    []JobEvent "json:\"events\""
}{Node: "wss6c3w9wngpycg.example.com", State: "running", Timestamp: "2015-07-13T20:37:01Z", Events: []JobEvent{}}, struct {
	Node      string     "json:\"node\""
	State     string     "json:\"state\""
	Timestamp string     "json:\"timestamp\""
	Events    []JobEvent "json:\"events\""
}{Node: "xxyyzz.example.com", State: "running", Timestamp: "2016-07-13T20:37:01Z", Events: []JobEvent{}}}}

var expectedJobNodes = &JobNodes{Items: []JobNode{JobNode{Transport: "pcp", FinishTimestamp: "2020-04-02T15:25:17Z", TransactionUUID: "", StartTimestamp: "2020-04-02T15:25:16Z", Name: "scotch-chapter.delivery.puppetlabs.net", Duration: 1.538, State: "finished", Details: map[string]interface{}{}, Result: map[string]interface{}{"latest": "1.1.1-1ubuntu2.1~18.04.5", "status": "out of date", "version": "1.1.0g-2ubuntu4.1"}, LatestEventID: 11, Timestamp: "2020-04-02T15:25:17Z"}, JobNode{Transport: "pcp", FinishTimestamp: "2020-04-02T15:25:18Z", TransactionUUID: "", StartTimestamp: "2020-04-02T15:25:16Z", Name: "legal-deposit.delivery.puppetlabs.net", Duration: 2.068, State: "finished", Details: map[string]interface{}{"message": "Message of latest event"}, Result: map[string]interface{}{"latest": "1:1.0.2k-19.el7", "status": "out of date", "version": "1:1.0.2k-12.el7"}, LatestEventID: 12, Timestamp: "2020-04-02T15:25:18Z"}}, NextEvents: NextEvents{ID: "https://legal-deposit.delivery.puppetlabs.net:8143/orchestrator/v1/jobs/3/events?start=13", Event: "13"}}
