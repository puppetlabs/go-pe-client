package orch

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInventory(t *testing.T) {

	// Test success
	setupGetResponder(t, orchInventory, "", "inventory-response.json")
	actual, err := orchClient.Inventory()
	require.Nil(t, err)
	require.Equal(t, expectedInventory, actual)

	// Test error
	setupErrorResponder(t, orchInventory)
	actual, err = orchClient.Inventory()
	require.Nil(t, actual)
	require.Equal(t, expectedError, err)

}

func TestInventoryNode(t *testing.T) {

	orchInventoryNodeFoo := strings.ReplaceAll(orchInventoryNode, "{node}", "foo")

	// Test success
	setupGetResponder(t, orchInventoryNodeFoo, "", "inventory-node-response.json")
	actual, err := orchClient.InventoryNode("foo")
	require.Nil(t, err)
	require.Equal(t, expectedInventoryNode, actual)

	// Test error
	setupErrorResponder(t, orchInventoryNodeFoo)
	actual, err = orchClient.InventoryNode("foo")
	require.Nil(t, actual)
	require.Equal(t, expectedError, err)

}

func TestInventoryCheck(t *testing.T) {

	// Test success
	setupPostResponder(t, orchInventory, "inventory-check-request.json", "inventory-check-response.json")
	actual, err := orchClient.InventoryCheck([]string{"foo.example.com", "bar.example.com", "baz.example.com"})
	require.Nil(t, err)
	require.Equal(t, expectedInventoryCheck, actual)

	// Test error
	setupErrorResponder(t, orchInventory)
	actual, err = orchClient.InventoryCheck([]string{"foo.example.com", "bar.example.com", "baz.example.com"})
	require.Nil(t, actual)
	require.Equal(t, expectedError, err)

}

var expectedInventory = []InventoryNode{{Name: "foo.example.com", Connected: true, Broker: "pcp://broker1.example.com/server", Timestamp: "2016-010-22T13:36:41.449Z"}, {Name: "bar.example.com", Connected: true, Broker: "pcp://broker2.example.com/server", Timestamp: "2016-010-22T13:39:16.377Z"}}

var expectedInventoryNode = &InventoryNode{Name: "foo.example.com", Connected: true, Broker: "pcp://broker.example.com/server", Timestamp: "2017-03-29T21:48:09.633Z"}

var expectedInventoryCheck = []InventoryNode{{Name: "foo.example.com", Connected: true, Broker: "pcp://broker.example.com/server", Timestamp: "2017-07-14T15:57:33.640Z"}, {Name: "bar.example.com", Connected: false, Broker: "", Timestamp: ""}, {Name: "baz.example.com", Connected: true, Broker: "pcp://broker.example.com/server", Timestamp: "2017-07-14T15:41:19.242Z"}}
