package orch

// Scope represents the scope of a job. Only a single field can be specified.
type Scope struct {
	Application string        `json:"application,omitempty"`
	Nodes       []string      `json:"nodes,omitempty"`
	Query       []interface{} `json:"query,omitempty"`
	NodeGroup   string        `json:"node_group,omitempty"`
}

// Owner represents the owner of a job
type Owner struct {
	ID    string `json:"id"`
	Login string `json:"login"`
}

// Pagination information about the current payload
type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

// Interval represents the time Interval for a task to run
type Interval struct {
	Units string `json:"units"`
	Value int    `json:"value"`
}

// ScheduleOptions represents the schedule options for a task
type ScheduleOptions struct {
	Interval Interval `json:"interval"`
}
