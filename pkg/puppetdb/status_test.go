package puppetdb

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestStatus performs a test on the FactNames endpoint and verifies the expected response is returned.
func TestStatus(t *testing.T) {
	// Test FactNames
	setupGetResponder(t, puppetDBStatus, "", "puppetdbstatus-response.json")
	actual, err := pdbClient.PDbStatus()
	require.Nil(t, err)
	require.Equal(t, expectedStatuses, actual)
}

var (
	expectedStatuses = &PDbStatus{"6.8.1-20200122_170412-gc886602"}
)
