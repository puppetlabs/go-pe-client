package orch

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	orchPlans    = "/orchestrator/v1/plans"
	orchPlanName = "/orchestrator/v1/plans/{module}/{planname}"
)

var plansRegex = regexp.MustCompile(`http.*\/orchestrator\/v1\/plans\/(.*)/(.*)`)

// Plans lists all known plans in a given environment (GET /plans)
func (c *Client) Plans(environment string) (*Plans, error) {
	payload := &Plans{}
	req := c.resty.R().SetResult(payload)
	if environment != "" {
		req.SetQueryParam("environment", environment)
	}
	r, err := req.Get(orchPlans)
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		if r.Error() != nil {
			return nil, r.Error().(error)
		}
		return nil, fmt.Errorf("%s error: %s", orchPlans, r.Status())
	}
	return payload, nil
}

// PlanByID extracts the module and planname from the supplied ID and calls Plan(...)
func (c *Client) PlanByID(environment, planID string) (*Plan, error) {
	results := plansRegex.FindStringSubmatch(planID)
	if len(results) != 3 {
		return nil, fmt.Errorf("unknown plan ID format: %s", planID)
	}
	module := results[1]
	planname := results[2]
	return c.Plan(environment, module, planname)
}

// Plan returns data about the specified plan, including metadata (GET /plans/:module/:planname)
func (c *Client) Plan(environment, module, planname string) (*Plan, error) {
	payload := &Plan{}
	req := c.resty.R().
		SetResult(payload).
		SetPathParams(map[string]string{
			"module":   module,
			"planname": planname,
		})
	if environment != "" {
		req.SetQueryParam("environment", environment)
	}
	r, err := req.Get(orchPlanName)
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		if r.Error() != nil {
			return nil, r.Error().(error)
		}
		replacer := strings.NewReplacer("{module}", module, "{planname}", planname)
		return nil, fmt.Errorf("%s error: %s", replacer.Replace(orchPlanName), r.Status())
	}
	return payload, nil
}

// Plans lists the known plans in a single environment
type Plans struct {
	Environment struct {
		Name   string `json:"name,omitempty"`
		CodeID string `json:"code_id,omitempty" structs:"-"`
	} `json:"environment,omitempty"`
	Items []struct {
		ID        string `json:"id,omitempty"`
		Name      string `json:"name,omitempty"`
		Permitted bool   `json:"permitted,omitempty"`
	} `json:"items,omitempty"`
}

// Plan contains information about a specific plan
type Plan struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Environment struct {
		Name   string `json:"name,omitempty"`
		CodeID string `json:"code_id,omitempty" structs:"-"`
	} `json:"environment,omitempty"`
	Metadata  TaskMetadata `json:"metadata,omitempty" structs:"-"`
	Permitted bool         `json:"permitted,omitempty"`
}
