package orch

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	orchTasks    = "/orchestrator/v1/tasks"
	orchTaskName = "/orchestrator/v1/tasks/{module}/{taskname}"
)

var idRegex = regexp.MustCompile(`http.*\/orchestrator\/v1\/tasks\/(.*)/(.*)`)

// Tasks lists all tasks in a given environment (GET /tasks)
func (c *Client) Tasks(environment string) (*Tasks, error) {
	payload := &Tasks{}
	req := c.resty.R().SetResult(payload)
	if environment != "" {
		req.SetQueryParam("environment", environment)
	}
	r, err := req.Get(orchTasks)
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		if r.Error() != nil {
			return nil, r.Error().(error)
		}
		return nil, fmt.Errorf("%s error: %s", orchTasks, r.Status())
	}
	return payload, nil
}

// TaskByID extracts the module and taskname from the supplied ID and calls Task(...)
func (c *Client) TaskByID(environment, taskID string) (*Task, error) {
	results := idRegex.FindStringSubmatch(taskID)
	if len(results) != 3 {
		return nil, fmt.Errorf("unknown task ID format: %s", taskID)
	}
	module := results[1]
	taskname := results[2]
	return c.Task(environment, module, taskname)
}

// Task returns data about a specified task, including metadata and file information. For the default task in a module, taskname is init. (GET /tasks/:module/:taskname)
func (c *Client) Task(environment, module, taskname string) (*Task, error) {
	payload := &Task{}
	req := c.resty.R().
		SetResult(payload).
		SetPathParams(map[string]string{
			"module":   module,
			"taskname": taskname,
		})
	if environment != "" {
		req.SetQueryParam("environment", environment)
	}
	r, err := req.Get(orchTaskName)
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		if r.Error() != nil {
			return nil, r.Error().(error)
		}
		replacer := strings.NewReplacer("{module}", module, "{taskname}", taskname)
		return nil, fmt.Errorf("%s error: %s", replacer.Replace(orchTaskName), r.Status())
	}
	return payload, nil
}

// Tasks is a list all tasks in a single environment
type Tasks struct {
	Environment struct {
		Name   string `json:"name,omitempty"`
		CodeID string `json:"code_id,omitempty"`
	} `json:"environment,omitempty"`
	Items []struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"items,omitempty"`
}

// Task contains data about a specified task, including metadata and file information
type Task struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Environment struct {
		Name   string `json:"name,omitempty"`
		CodeID string `json:"code_id,omitempty"`
	} `json:"environment,omitempty"`
	Metadata TaskMetadata `json:"metadata,omitempty"`
	Files    []TaskFile   `json:"files,omitempty"`
}

// TaskMetadata ...
type TaskMetadata struct {
	Description     string                 `json:"description,omitempty"`
	Private         bool                   `json:"private,omitempty"`
	SupportsNoop    bool                   `json:"supports_noop,omitempty"`
	InputMethod     string                 `json:"input_method,omitempty"`
	Parameters      map[string]TaskParam   `json:"parameters,omitempty"`
	Extensions      map[string]interface{} `json:"extensions,omitempty"`
	Implementations []TaskImplementation   `json:"implementations"`
}

// TaskFile ...
type TaskFile struct {
	Filename string `json:"filename,omitempty"`
	URI      struct {
		Path   string `json:"path,omitempty"`
		Params struct {
			Environment string `json:"environment,omitempty"`
		} `json:"params,omitempty"`
	} `json:"uri,omitempty"`
	Sha256    string `json:"sha256,omitempty"`
	SizeBytes int    `json:"size_bytes,omitempty"`
}

// TaskParam in the task metadata
type TaskParam struct {
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
}

// TaskImplementation in the task metadata
type TaskImplementation struct {
	Name         string   `json:"name"`
	Requirements []string `json:"requirements"`
	InputMethod  string   `json:"input_method"`
}
