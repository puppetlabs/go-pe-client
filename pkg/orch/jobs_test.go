package orch

import (
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/fatih/structs"
	"github.com/stretchr/testify/require"
)

func TestJobs(t *testing.T) {

	// Test success
	setupGetResponder(t, orchJobs, "", "jobs-response.json")
	actual, err := orchClient.Jobs()
	require.Nil(t, err)
	require.False(t, structs.HasZero(actual), spew.Sdump(actual))

	// Test error
	setupErrorResponder(t, orchJobs)
	actual, err = orchClient.Jobs()
	require.Nil(t, actual)
	require.Equal(t, expectedError, err)

}

func TestJob(t *testing.T) {

	testURL := strings.ReplaceAll(orchJob, "{job-id}", "123")

	// Test success
	setupGetResponder(t, testURL, "", "job-response.json")
	actual, err := orchClient.Job("123")
	require.Nil(t, err)
	require.False(t, structs.HasZero(actual), spew.Sdump(actual))

	// Test error
	setupErrorResponder(t, testURL)
	actual, err = orchClient.Job("123")
	require.Nil(t, actual)
	require.Equal(t, expectedError, err)
	require.False(t, errors.Is(err, expectedJobNotFoundErr))

	//Test HTTP error
	setupResponderWithStatusCodeAndBody(t, testURL, http.StatusNotFound, []byte(`{"StatusCode": 400}`))
	actual, err = orchClient.Job("123")
	testHTTPError(t, actual, err, http.StatusNotFound)
}

func TestJobReport(t *testing.T) {

	testURL := strings.ReplaceAll(orchJobReport, "{job-id}", "123")

	// Test success
	setupGetResponder(t, testURL, "", "job-report-response.json")
	actual, err := orchClient.JobReport("123")
	require.Nil(t, err)
	require.False(t, structs.HasZero(actual), spew.Sdump(actual))

	// Test error
	setupErrorResponder(t, testURL)
	actual, err = orchClient.JobReport("123")
	require.Nil(t, actual)
	require.Equal(t, expectedError, err)

	// Test HTTP error
	setupResponderWithStatusCodeAndBody(t, testURL, http.StatusForbidden, []byte(`{"StatusCode": 400}`))
	actual, err = orchClient.JobReport("123")
	testHTTPError(t, actual, err, http.StatusForbidden)

}

func TestJobNodes(t *testing.T) {

	testURL := strings.ReplaceAll(orchJobNodes, "{job-id}", "123")

	// Test success
	setupGetResponder(t, testURL, "", "jobs-nodes-response.json")
	actual, err := orchClient.JobNodes("123")
	require.Nil(t, err)
	require.False(t, structs.HasZero(actual), spew.Sdump(actual))

	// Test error
	setupErrorResponder(t, testURL)
	actual, err = orchClient.JobNodes("123")
	require.Nil(t, actual)
	require.Equal(t, expectedError, err)

	// Test HTTP error
	setupResponderWithStatusCodeAndBody(t, testURL, http.StatusBadRequest, []byte(`{"StatusCode": 400}`))
	actual, err = orchClient.JobNodes("123")
	testHTTPError(t, actual, err, http.StatusBadRequest)

}
