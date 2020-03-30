package orch

// CommandTask runs a permitted task job across a set of nodes (POST /command/task)
func (c *Client) CommandTask(taskRequest *TaskRequest) (*JobID, error) {
	payload := JobID{}
	r, err := c.resty.R().
		SetResult(&payload).
		SetBody(taskRequest).
		Post("/orchestrator/v1/command/task")
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		return nil, r.Error().(error)
	}
	return &payload, nil
}

// JobID identifies a single Orchestrator job
type JobID struct {
	Job struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"job"`
}

// TaskRequest describes a task to be run
type TaskRequest struct {
	Environment string            `json:"environment,omitempty"`
	Task        string            `json:"task"`
	Params      map[string]string `json:"params"`
	Scope       TaskScope         `json:"scope"`
}

// TaskScope is part of a TaskRequest. Only a single field can be specified.
type TaskScope struct {
	Application string        `json:"application,omitempty"`
	Nodes       []string      `json:"nodes,omitempty"`
	Query       []interface{} `json:"query,omitempty"`
	NodeGroup   string        `json:"node_group,omitempty"`
}
