package puppetdb

const (
	factNames = "/pdb/query/v4/fact-names"
	factPaths = "/pdb/query/v4/fact-paths"
	facts     = "/pdb/query/v4/facts"
)

// FactNames will return an alphabetical list of all known fact names, including those which are known only for deactivated nodes.
func (c *Client) FactNames(pagination *Pagination, orderBy *OrderBy) ([]string, error) {
	payload := []string{}
	err := getRequest(c, factNames, "", pagination, orderBy, &payload)
	return payload, err
}

// FactPaths will return a set of all known fact paths for all known nodes, and is intended as a counterpart to the fact-names endpoint.
func (c *Client) FactPaths(query string, pagination *Pagination, orderBy *OrderBy) ([]FactPath, error) {
	payload := []FactPath{}
	err := getRequest(c, factPaths, query, pagination, orderBy, &payload)
	return payload, err
}

// Facts will return all facts matching the given query. Facts for deactivated nodes are not included in the response.
func (c *Client) Facts(query string, pagination *Pagination, orderBy *OrderBy) ([]Fact, error) {
	payload := []Fact{}
	err := getRequest(c, facts, query, pagination, orderBy, &payload)
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
	Count       int         `json:"count"`
}

// FactPath represents a fact-path returned by the facts-paths endpoint.
// Path ([]string): an array of the parts that make up the path
// Type (string): the type of the fact, string, integer etc
type FactPath struct {
	Name string        `json:"name"`
	Path []interface{} `json:"path"`
	Type string        `json:"type"`
}
