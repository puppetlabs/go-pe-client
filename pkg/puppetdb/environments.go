package puppetdb

const (
	environments = "/pdb/query/v4/environments"
)

// Environments returns a list of all known environments
func (c *Client) Environments() ([]Environment, error) {
	payload := []Environment{}
	err := getRequest(c, environments, "", nil, nil, &payload)
	return payload, err
}

// Environment represents a PuppetDB environment
type Environment struct {
	Name string `json:"name"`
}
