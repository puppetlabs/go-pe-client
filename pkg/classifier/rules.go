package classifier

import (
	"encoding/json"
)

const (
	rules = "/classifier-api/v1/rules/translate"
)

// TranslateRules converts a group's rule condition into PuppetDB query syntax.
func (c *Client) TranslateRules(rule string) (string, error) {
	var payload Rule
	err := postRequest(c, rules, rule, &payload)
	if err != nil {
		return "", err
	}

	data, err := json.Marshal(payload.Query)
	return string(data), err
}

// Rule ...
type Rule struct {
	Query interface{}
}
