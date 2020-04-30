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

// CommandTaskTarget creates a new task-target (POST /command/task_target)
func (c *Client) CommandTaskTarget(taskTargetRequest *TaskTargetRequest) (*TaskTargetJobID, error) {
	payload := TaskTargetJobID{}
	r, err := c.resty.R().
		SetResult(&payload).
		SetBody(taskTargetRequest).
		Post("/orchestrator/v1/command/task_target")
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		return nil, r.Error().(error)
	}
	return &payload, nil
}

// TaskTargetJobID identifies a task_target job
type TaskTargetJobID struct {
	TaskTargetJob struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"task_target"`
}

// TaskTargetRequest is a collection of tasks, nodes and node groups that define a permission group
type TaskTargetRequest struct {
	DisplayName string   `json:"display_name"`
	Tasks       []string `json:"tasks,omitempty"`
	AllTasks    bool     `json:"all_tasks,omitempty"`
	Nodes       []string `json:"nodes"`
	NodeGroups  []string `json:"node_groups"`
	PQLQuery    string   `json:"pql_query,omitempty"`
}

// CommandPlanRun runs a plan via the plan executor (POST /command/plan_run)
func (c *Client) CommandPlanRun(planRunRequest *PlanRunRequest) (*PlanRunJobID, error) {
	payload := PlanRunJobID{}
	r, err := c.resty.R().
		SetResult(&payload).
		SetBody(planRunRequest).
		Post("/orchestrator/v1/command/plan_run")
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		return nil, r.Error().(error)
	}
	return &payload, nil
}

// PlanRunJobID identifies a plan_run job
type PlanRunJobID struct {
	Name string `json:"name"`
}

// PlanRunRequest describes a plan_run request
type PlanRunRequest struct {
	Name        string                 `json:"plan_name"`
	Params      map[string]interface{} `json:"params"`
	Environment string                 `json:"environment,omitempty"`
	Description string                 `json:"description,omitempty"`
}

// CommandStop stops a orchestrator job that is currently in progress (POST /command/stop)
func (c *Client) CommandStop(stopRequest *StopRequest) (*StopJobID, error) {
	payload := StopJobID{}
	r, err := c.resty.R().
		SetResult(&payload).
		SetBody(stopRequest).
		Post("/orchestrator/v1/command/stop")
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		return nil, r.Error().(error)
	}
	return &payload, nil
}

// StopJobID describes jobs that were stopped successfully
type StopJobID struct {
	Job struct {
		ID    string         `json:"id"`
		Name  string         `json:"name"`
		Nodes map[string]int `json:"nodes"`
	} `json:"job"`
}

// StopRequest describes a stop request
type StopRequest struct {
	Job string `json:"job"`
}
