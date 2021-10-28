package puppetdb

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestStatus performs a test on the puppetdb-status endpoint and verifies the expected response is returned.
func TestStatus(t *testing.T) {
	// Test Status success
	setupGetResponder(t, puppetDBStatus, "", "puppetdbstatus-response.json")
	actual, err := pdbClient.PDbStatus()
	require.Nil(t, err)
	require.Equal(t, expectedStatuses, actual)

	// Test error
	setupURLErrorResponder(t, puppetDBStatus)
	actual, err = pdbClient.PDbStatus()
	require.Equal(t, expectedErrorStatuses, actual)
	require.Equal(t, errExpectedURL, err)
}

var (
	expectedStatuses      = &PDbStatus{"6.8.1-20200122_170412-gc886602"}
	expectedErrorStatuses = &PDbStatus{""}
	errExpectedURL        = fmt.Errorf("https://test-host:8081/status/v1/services/puppetdb-status: 404: \"{\"Op\":\"nil\",\"URL\":\"https://test-host:8081\",\"Err\":null}\"")
)
