package orch

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
)

const (
	orchJob       = "/orchestrator/v1/jobs/{job-id}"
	orchJobNodes  = "/orchestrator/v1/jobs/{job-id}/nodes"
	orchJobReport = "/orchestrator/v1/jobs/{job-id}/report"
	orchJobs      = "/orchestrator/v1/jobs"
)

//ErrJobNotFound will be returned when we get a 404 from PE for the specific job.
var ErrJobNotFound = errors.New("job not found")

//processJobResponse will process the response for specific jobs and return the
//special case of not found if the job does not exist on PE.
func processJobResponse(r *resty.Response, message string) error {
	if r.IsError() {
		var err error
		if r.StatusCode() == http.StatusNotFound {
			err = fmt.Errorf("%s error: %w", message, ErrJobNotFound)
		} else {
			if r.Error() != nil {
				return r.Error().(error)
			}
			err = fmt.Errorf("%s error: %s", message, r.Status())
		}
		return err
	}
	return nil
}

// Jobs lists all of the jobs known to the orchestrator (GET /jobs)
func (c *Client) Jobs() (*Jobs, error) {
	payload := &Jobs{}
	r, err := c.resty.R().SetResult(&payload).Get(orchJobs)
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		if r.Error() != nil {
			return nil, r.Error().(error)
		}
		return nil, fmt.Errorf("%s error: %s", orchJobs, r.Status())
	}
	return payload, nil
}

// Job lists all details of a given job (GET /jobs/:job-id)
func (c *Client) Job(jobID string) (*Job, error) {
	payload := &Job{}
	r, err := c.resty.R().
		SetResult(&payload).
		SetPathParams(map[string]string{"job-id": jobID}).
		Get(orchJob)
	if err != nil {
		return nil, err
	}
	if err = processJobResponse(r, strings.ReplaceAll(orchJob, "{job-id}", jobID)); err != nil {
		return nil, err
	}

	return payload, nil
}

// JobReport returns the report for a given job (GET /jobs/:job-id/report)
func (c *Client) JobReport(jobID string) (*JobReport, error) {
	payload := &JobReport{}
	r, err := c.resty.R().
		SetResult(&payload).
		SetPathParams(map[string]string{"job-id": jobID}).
		Get(orchJobReport)
	if err != nil {
		return nil, err
	}
	if err = processJobResponse(r, strings.ReplaceAll(orchJobReport, "{job-id}", jobID)); err != nil {
		return nil, err
	}
	return payload, nil
}

// JobNodes lists all of the nodes associated with a given job (GET /jobs/:job-id/nodes)
func (c *Client) JobNodes(jobID string) (*JobNodes, error) {
	payload := &JobNodes{}
	r, err := c.resty.R().
		SetResult(&payload).
		SetPathParams(map[string]string{"job-id": jobID}).
		Get(orchJobNodes)
	if err != nil {
		return nil, err
	}
	if err = processJobResponse(r, strings.ReplaceAll(orchJobNodes, "{job-id}", jobID)); err != nil {
		return nil, err
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
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	State       string                 `json:"state"`
	Type        string                 `json:"type"`
	Command     string                 `json:"command"`
	Options     map[string]interface{} `json:"options"`
	NodeCount   int                    `json:"node_count"`
	NodeStates  NodeStates             `json:"node_states"`
	Owner       map[string]interface{} `json:"owner"`
	Description string                 `json:"description"`
	Timestamp   string                 `json:"timestamp"`
	Environment Environment            `json:"environment"`
	Status      []Status               `json:"status"`
	Nodes       Nodes                  `json:"nodes"`
	Events      Events                 `json:"events"`
	Report      Report                 `json:"report"`
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
