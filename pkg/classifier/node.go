package classifier

import (
	"encoding/json"
	"fmt"
)

const (
	uri = "/classifier-api/v2/classified/nodes"
)

// Node will return the Node matching the given id.
// certname is the hostname of the node to query.
func (c *Client) Node(certname string) (Node, error) {
	payload, err := PostRequest(c, fmt.Sprintf("%s/%s", uri, certname))

	if err != nil {
		return Node{}, err
	}

	var node Node
	err = json.Unmarshal(payload, &node)
	return node, err
}

// Node represents the response to the classifier nodes v2 endpoint.
type Node struct {
	Name        string `json:"name"`
	Environment string `json:"environment"`
	Groups      []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"groups"`
	Classes struct {
	} `json:"classes"`
	Parameters struct {
	} `json:"parameters"`
	ConfigData struct {
	} `json:"config_data"`
}
