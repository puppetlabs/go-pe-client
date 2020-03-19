package orch

import "github.com/davecgh/go-spew/spew"

// Inventory lists all nodes that are connected to the PCP broker (GET /inventory)
func (c *Client) Inventory() (*[]InventoryNode, error) {
	payload := map[string][]InventoryNode{}
	r, err := c.resty.R().SetResult(&payload).Get("/orchestrator/v1/inventory")
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		return nil, r.Error().(error)
	}
	inventoryNodes := payload["items"]
	return &inventoryNodes, nil
}

// InventoryNode returns information about whether the requested node is connected to the PCP broker (GET /inventory/:node)
func (c *Client) InventoryNode(node string) (*InventoryNode, error) {
	payload := &InventoryNode{}
	request := c.resty.R().
		SetResult(payload).
		SetPathParams(map[string]string{
			"node": node,
		})
	response, err := request.Get("/orchestrator/v1/inventory/{node}")
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, response.Error().(error)
	}
	return payload, nil
}

// InventoryCheck checks if the given list of nodes is connected to the PCP broker (POST /inventory)
func (c *Client) InventoryCheck(nodes []string) (*[]InventoryNode, error) {
	payload := map[string][]InventoryNode{}
	r, err := c.resty.R().
		SetResult(&payload).
		SetBody(map[string]interface{}{"nodes": nodes}).
		Post("/orchestrator/v1/inventory")
	spew.Dump(r)
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		return nil, r.Error().(error)
	}
	inventoryNodes := payload["items"]
	return &inventoryNodes, nil
}

// InventoryNode contains data about a single node
type InventoryNode struct {
	Name      string `json:"name,omitempty"`
	Connected bool   `json:"connected,omitempty"`
	Broker    string `json:"broker,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}
