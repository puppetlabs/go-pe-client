package pe

import (
	"fmt"
)

const (
	apiEnvironments = "/api/environments"
)

// Environments gets the list of environments from the PE API (GET /api/environments)
func (c *Client) Environments() ([]string, error) {
	payload := []string{}
	r, err := c.resty.R().
		SetResult(&payload).
		Get(apiEnvironments)
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		if r.Error() != nil {
			return nil, r.Error().(error)
		}
		return nil, fmt.Errorf("%s error: %s", apiEnvironments, r.Status())
	}
	return payload, nil
}
