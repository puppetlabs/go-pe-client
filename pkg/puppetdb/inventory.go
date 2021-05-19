package puppetdb

const (
	inventory = "/pdb/query/v4/inventory"
)

// Inventory enables an alternative query syntax for digging into structured facts, and can be used instead of the facts,
// fact-contents, and factsets endpoints for most fact-related queries.
func (c *Client) Inventory(query string, pagination *Pagination, orderBy *OrderBy) ([]Inventory, error) {
	payload := []Inventory{}
	err := getRequest(c, inventory, query, pagination, orderBy, &payload)
	return payload, err
}

// InventoryMap an alternative to Inventory which returns an Array of Maps which will allow for fields that do
// not fit into the Inventory struct, i.e. dot notation fields.
func (c *Client) InventoryMap(query string, pagination *Pagination, orderBy *OrderBy) ([]map[string]interface{}, error) {
	var payload []map[string]interface{}
	err := getRequest(c, inventory, query, pagination, orderBy, &payload)
	return payload, err
}

// Inventory is a PuppetDB node with facts and trusted facts
type Inventory struct {
	Certname    string                 `json:"certname"`
	Timestamp   string                 `json:"timestamp"`
	Environment string                 `json:"environment"`
	Facts       map[string]interface{} `json:"facts"`
	Trusted     map[string]interface{} `json:"trusted"`
	Count       int                    `json:"count"`
}
