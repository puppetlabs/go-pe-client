package orch

// Inventory contains the response from /orchestrator/v1/inventory
type Inventory struct {
	Items []struct {
		Name      string `json:"name"`
		Connected bool   `json:"connected"`
		Broker    string `json:"broker"`
		Timestamp string `json:"timestamp"`
	} `json:"items"`
}

// Inventory sends requests to /orchestrator/v1/inventory
func (c *Client) Inventory() (*Inventory, error) {
	payload := &Inventory{}
	r, err := c.resty.R().SetResult(payload).Get("/orchestrator/v1/inventory")
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		return nil, r.Error().(error)
	}
	return payload, nil
}
