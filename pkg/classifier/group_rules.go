package classifier

import (
	"strings"
)

const (
	groupRulesPathTemplate = "/classifier-api/v1/groups/{group-id}/rules"
)

// GroupRules will return the rules for the specified group.
func (c *Client) GroupRules(groupID string) (GroupRules, error) {
	var payload GroupRules
	path := strings.ReplaceAll(groupRulesPathTemplate, "{group-id}", groupID)
	err := getRequest(c, path, nil, &payload)

	return payload, err
}

// GroupRules stores the response from the group rules endpoint.
type GroupRules struct {
	Rule              interface{}
	RuleWithInherited interface{} `json:"rule_with_inherited"`
	Translated        struct {
		NodesQueryFormat     interface{} `json:"nodes_query_format"`
		InventoryQueryFormat interface{} `json:"inventory_query_format"`
	}
}
