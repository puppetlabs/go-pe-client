package orch

// Tasks contains the response from /orchestrator/v1/tasks
type Tasks struct {
	Environment struct {
		Name   string `json:"name"`
		CodeID string `json:"code_id"`
	} `json:"environment"`
	Items []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"items"`
}

// Tasks sends requests to /orchestrator/v1/tasks
func (c *Client) Tasks() (*Tasks, error) {
	payload := &Tasks{}
	r, err := c.resty.R().SetResult(payload).Get("/orchestrator/v1/tasks")
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		return nil, r.Error().(error)
	}
	return payload, nil
}
