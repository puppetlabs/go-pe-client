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
	Scope       Scope             `json:"scope"`
}

// CommandScheduleTask schedules a task to run at a future date and time (POST /command/schedule_task)
func (c *Client) CommandScheduleTask(scheduleTaskRequest *ScheduleTaskRequest) (*ScheduledJobID, error) {
	payload := ScheduledJobID{}
	r, err := c.resty.R().
		SetResult(&payload).
		SetBody(scheduleTaskRequest).
		Post("/orchestrator/v1/command/schedule_task")
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		return nil, r.Error().(error)
	}
	return &payload, nil
}

// ScheduledJobID identifies a single scheduled job
type ScheduledJobID struct {
	ScheduledJob struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"scheduled_job"`
}

// ScheduleTaskRequest describes a scheduled task
type ScheduleTaskRequest struct {
	Environment   string            `json:"environment,omitempty"`
	Task          string            `json:"task"`
	Params        map[string]string `json:"params"`
	Scope         Scope             `json:"scope"`
	ScheduledTime string            `json:"scheduled_time"`
}
