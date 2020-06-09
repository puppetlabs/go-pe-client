package classifier

import (
	"fmt"
	"time"
)

const (
	groups = "/classifier-api/v1/groups"
)

// Groups will return all groups.
func (c *Client) Groups(pagination *Pagination) ([]Group, error) {
	payload := []Group{}
	err := getRequest(c, groups, pagination, &payload)
	return payload, err
}

// Group will return the group matching the given id.
func (c *Client) Group(id string) (Group, error) {
	payload := Group{}
	err := getRequest(c, fmt.Sprintf("%s/%s", groups, id), nil, &payload)
	return payload, err
}

// Group represents a group returned by the groups endpoint.
// See https://puppet.com/docs/pe/2018.1/groups_endpoint.html#get_v1_groups__response_format
type Group struct {
	ID                string
	Name              string
	Description       string
	Environment       string
	EnvironmentTrumps bool `json:"environment_trumps"`
	Parent            string
	Rule              interface{}
	Classes           map[string]interface{}
	ConfigData        map[string]interface{} `json:"config_data"`
	Deleted           map[string]interface{}
	Variables         map[string]interface{}
	LastEdited        time.Time `json:"last_edited"`
	SerialNumber      int       `json:"serial_number"`
}
