package puppetdb

const (
	factnames = "/pdb/query/v4/fact-names"
	facts     = "/pdb/query/v4/facts"
)

// FactNames will return an alphabetical list of all known fact names, including those which are known only for deactivated nodes.
func (c *Client) FactNames(pagination *Pagination) ([]string, error) {
	payload := []string{}
	err := getRequest(c, factnames, "", pagination, &payload)
	return payload, err
}

// Facts will return all facts matching the given query. Facts for deactivated nodes are not included in the response.
func (c *Client) Facts(query string, pagination *Pagination) ([]Fact, error) {
	payload := []Fact{}
	err := getRequest(c, facts, query, pagination, &payload)
	return payload, err
}

// Fact represents a fact returned by the Facts endpoint.
// Name (string): the name of the fact.
// Value (string, numeric, Boolean): the value of the fact.
// Certname (string): the node associated with the fact.
// Environment (string): the environment associated with the fact.
type Fact struct {
	Name        string      `json:"name"`
	Value       interface{} `json:"value"`
	Certname    string      `json:"certname"`
	Environment string      `json:"environment"`
}
