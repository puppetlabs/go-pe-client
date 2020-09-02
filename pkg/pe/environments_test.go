package pe

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnvironments(t *testing.T) {

	// Test without environment
	setupGetResponder(t, apiEnvironments, "", "environments.json")
	actual, err := peClient.Environments()
	require.Nil(t, err)
	require.Equal(t, expectedEnvironments, actual)

}

var expectedEnvironments = []string{"production", "test"}
