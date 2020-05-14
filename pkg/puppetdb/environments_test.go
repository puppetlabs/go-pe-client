package puppetdb

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/fatih/structs"
	"github.com/stretchr/testify/require"
)

// TestFactNames performs a test on the FactNames endpoint and verifies the expected response is returned.
func TestEnvironment(t *testing.T) {
	setupGetResponder(t, environments, "", "environments-response.json")
	actual, err := pdbClient.Environments()
	require.Nil(t, err)
	require.True(t, len(actual) > 0)
	require.False(t, structs.HasZero(actual[0]), spew.Sdump(actual))
}
