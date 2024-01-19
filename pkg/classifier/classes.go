package classifier

const (
	classes = "/classifier-api/v1/classes"
)

func (c *Client) Classes(pagination *Pagination) ([]Class, error) {
	payload := []Class{}
	err := getRequest(c, classes, nil, &payload)
	return payload, err
}

// Class represents a group returned by the classes endpoint.
// See https://www.puppet.com/docs/pe/2019.8/classes_endpoint#get_v1_classes
type Class struct {
	Name        string
	Environment string
	Parameters  map[string]interface{}
}
