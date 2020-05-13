package orch

import (
	"fmt"
	"strings"
)

const (
	orchInventory     = "/orchestrator/v1/inventory"
	orchInventoryNode = "/orchestrator/v1/inventory/{node}"
)

// Inventory lists all nodes that are connected to the PCP broker (GET /inventory)
func (c *Client) Inventory() ([]InventoryNode, error) {
	payload := map[string][]InventoryNode{}
	r, err := c.resty.R().SetResult(&payload).Get(orchInventory)
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		if r.Error() != nil {
			return nil, r.Error().(error)
		}
		return nil, fmt.Errorf("%s error: %s", orchInventory, r.Status())
	}
	inventoryNodes := payload["items"]
	return inventoryNodes, nil
}

// InventoryNode returns information about whether the requested node is connected to the PCP broker (GET /inventory/:node)
func (c *Client) InventoryNode(node string) (*InventoryNode, error) {
	payload := &InventoryNode{}
	req := c.resty.R().
		SetResult(payload).
		SetPathParams(map[string]string{
			"node": node,
		})
	r, err := req.Get(orchInventoryNode)
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		if r.Error() != nil {
			return nil, r.Error().(error)
		}
		inventoryNode := strings.ReplaceAll(orchInventoryNode, "{node}", node)
		return nil, fmt.Errorf("%s error: %s", inventoryNode, r.Status())
	}
	return payload, nil
}

// InventoryCheck checks if the given list of nodes is connected to the PCP broker (POST /inventory)
func (c *Client) InventoryCheck(nodes []string) ([]InventoryNode, error) {
	payload := map[string][]InventoryNode{}
	r, err := c.resty.R().
		SetResult(&payload).
		SetBody(map[string]interface{}{"nodes": nodes}).
		Post(orchInventory)
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		if r.Error() != nil {
			return nil, r.Error().(error)
		}
		return nil, fmt.Errorf("%s error: %s", orchInventory, r.Status())
	}
	inventoryNodes := payload["items"]
	return inventoryNodes, nil
}

// InventoryNode contains data about a single node
type InventoryNode struct {
	Name      string `json:"name,omitempty"`
	Connected bool   `json:"connected,omitempty"`
	Broker    string `json:"broker,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}
