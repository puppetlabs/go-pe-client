package orch

// Tasks contains the response from /v1/tasks
type Tasks struct {
	Environment struct {
		Name   string      `json:"name"`
		CodeID interface{} `json:"code_id"`
	} `json:"environment,omitempty"`
	Items []struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		Permitted bool   `json:"permitted"`
	} `json:"items,omitempty"`
}

// Tasks returns information about all installed tasks
func (c *Client) Tasks() (*Tasks, error) {
	tasks := &Tasks{}
	r, err := c.resty.R().
		SetResult(tasks).
		Get("/orchestrator/v1/tasks")
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		return nil, r.Error().(error)
	}
	return tasks, nil
}
