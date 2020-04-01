package puppetdb

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestInventory performs a test on the Inventory endpoint and verifies the expected response is returned.
func TestInventory(t *testing.T) {
	query := `["=", "certname", "foobar.delivery.puppetlabs.net"]`
	setupGetResponder(t, inventory, "query="+query, "inventory.json")
	actual, err := pdbClient.Inventory(query)
	require.Nil(t, err)
	require.Equal(t, expectedInventory, actual)
}

var expectedInventory = []Inventory{{Certname: "foobar.delivery.puppetlabs.net", Timestamp: "2020-03-30T08:55:23.348Z", Environment: "production", Facts: map[string]interface{}{"agent_specified_environment": "production", "aio_agent_build": "6.10.1", "architecture": "x86_64"}, Trusted: map[string]interface{}{"authenticated": "remote", "domain": "delivery.puppetlabs.net"}}}
