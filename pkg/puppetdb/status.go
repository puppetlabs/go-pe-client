package puppetdb

const (
	puppetDBStatus = "/status/v1/services/puppetdb-status"
)

// PDbStatus will return the status of the pdb server, specifically the service version.
func (c *Client) PDbStatus() (*PDbStatus, error) {
	payload := &PDbStatus{}
	err := getRequest(c, puppetDBStatus, "", nil, nil, &payload)

	return payload, err
}

// PDbStatus represents the puppet db status returned from the endpoint.
// ServiceVersion (string): the service version of the pe server the endpoint calls out to
type PDbStatus struct {
	ServiceVersion string `json:"service_version,omitempty"`
}
