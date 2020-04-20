package orch

const (
	job       = "/orchestrator/v1/jobs/{job-id}"
	jobNodes  = "/orchestrator/v1/jobs/{job-id}/nodes"
	jobReport = "/orchestrator/v1/jobs/{job-id}/report"
	jobs      = "/orchestrator/v1/jobs"
)

// Jobs lists all of the jobs known to the orchestrator (GET /jobs)
func (c *Client) Jobs() (*Jobs, error) {
	payload := &Jobs{}
	r, err := c.resty.R().SetResult(&payload).Get(jobs)
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		return nil, r.Error().(error)
	}
	return payload, nil
}

// Job lists all details of a given job (GET /jobs/:job-id)
func (c *Client) Job(jobID string) (*Job, error) {
	payload := &Job{}
	r, err := c.resty.R().
		SetResult(&payload).
		SetPathParams(map[string]string{"job-id": jobID}).
		Get(job)
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		return nil, r.Error().(error)
	}
	return payload, nil
}

// JobReport returns the report for a given job (GET /jobs/:job-id/report)
func (c *Client) JobReport(jobID string) (*JobReport, error) {
	payload := &JobReport{}
	r, err := c.resty.R().
		SetResult(&payload).
		SetPathParams(map[string]string{"job-id": jobID}).
		Get(jobReport)
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		return nil, r.Error().(error)
	}
	return payload, nil
}

// JobNodes lists all of the nodes associated with a given job (GET /jobs/:job-id/nodes)
func (c *Client) JobNodes(jobID string) (*JobNodes, error) {
	payload := &JobNodes{}
	r, err := c.resty.R().
		SetResult(&payload).
		SetPathParams(map[string]string{"job-id": jobID}).
		Get(jobNodes)
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		return nil, r.Error().(error)
	}
	return payload, nil
}

// Jobs contains data about all jobs
type Jobs struct {
	Items      []Job      `json:"items"`
	Pagination Pagination `json:"pagination"`
}

// Job contains data about a single job
type Job struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Command     string      `json:"command"`
	Options     Options     `json:"options"`
	NodeCount   int         `json:"node_count"`
	Owner       Owner       `json:"owner"`
	Description string      `json:"description"`
	Timestamp   string      `json:"timestamp"`
	Environment Environment `json:"environment"`
	Status      []Status    `json:"status"`
	Nodes       Nodes       `json:"nodes"`
	Report      Report      `json:"report"`
}

// Environment in the current job
type Environment struct {
	Name string `json:"name"`
}

// Status of the current job
type Status struct {
	State     string `json:"state"`
	EnterTime string `json:"enter_time"`
	ExitTime  string `json:"exit_time"`
}

// Nodes in the current job
type Nodes struct {
	ID string `json:"id"`
}

// Events in the current job
type Events struct {
	ID string `json:"id"`
}

// Report for the current job
type Report struct {
	ID string `json:"id"`
}

// NodeStates for the current job
type NodeStates struct {
	Finished int `json:"finished"`
	Errored  int `json:"errored"`
	Failed   int `json:"failed"`
	Running  int `json:"running"`
}

// Options for the current job
type Options struct {
	Concurrency        interface{} `json:"concurrency"`
	Noop               bool        `json:"noop"`
	Trace              bool        `json:"trace"`
	Debug              bool        `json:"debug"`
	Scope              Scope       `json:"scope"`
	EnforceEnvironment bool        `json:"enforce_environment"`
	Environment        string      `json:"environment"`
	Evaltrace          bool        `json:"evaltrace"`
	Target             interface{} `json:"target"`
	Description        string      `json:"description"`
}

// JobReport contains the report for a single job
type JobReport struct {
	Items []struct {
		Node      string     `json:"node"`
		State     string     `json:"state"`
		Timestamp string     `json:"timestamp"`
		Events    []JobEvent `json:"events"`
	} `json:"items"`
}

// JobEvent contains a single event from a job
type JobEvent struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Timestamp string `json:"timestamp"`
	Details   struct {
		Node   string `json:"node"`
		Detail struct {
			Noop bool `json:"noop"`
		} `json:"detail"`
	} `json:"details"`
	Message string `json:"message"`
}

// JobNodes is a list of all nodes associated with a given job
type JobNodes struct {
	Items      []JobNode  `json:"items"`
	NextEvents NextEvents `json:"next-events"`
}

// JobNode is a single node associated with a given job
type JobNode struct {
	Transport       string                 `json:"transport"`
	FinishTimestamp string                 `json:"finish_timestamp"`
	TransactionUUID string                 `json:"transaction_uuid"`
	StartTimestamp  string                 `json:"start_timestamp"`
	Name            string                 `json:"name"`
	Duration        float64                `json:"duration"`
	State           string                 `json:"state"`
	Details         map[string]interface{} `json:"details"`
	Result          map[string]interface{} `json:"result"`
	LatestEventID   int                    `json:"latest-event-id"`
	Timestamp       string                 `json:"timestamp"`
}

// NextEvents section of the response
type NextEvents struct {
	ID    string `json:"id"`
	Event string `json:"event"`
}
