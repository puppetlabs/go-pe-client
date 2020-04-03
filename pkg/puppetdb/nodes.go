package puppetdb

import (
	"fmt"
	"strconv"
)

func (p Pagination) toParams() map[string]string {
	params := map[string]string{}
	if p.Limit > 0 {
		params["limit"] = strconv.Itoa(p.Limit)
	}
	if p.Offset > 0 {
		params["offset"] = strconv.Itoa(p.Offset)
	}

	if p.IncludeTotal {
		params["include_total"] = strconv.FormatBool(p.IncludeTotal)
	}
	return params
}

// Nodes will return all nodes matching the given query. Deactivated and expired nodes aren’t included in the response.
func (c *Client) Nodes(query string, pagination Pagination) ([]Node, error) {
	payload := []Node{}
	req := c.resty.R().SetResult(&payload)
	if query != "" {
		req.SetQueryParam("query", query)
	}
	if pagination.toParams() != nil {
		req.SetQueryParams(pagination.toParams())
	}

	r, err := req.Get("/pdb/query/v4/nodes")
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		return nil, fmt.Errorf("%s: %s", r.Status(), r.Body())
	}
	return payload, nil
}

// Node is a PuppetDB node
type Node struct {
	Deactivated                  interface{} `json:"deactivated"`
	LatestReportHash             string      `json:"latest_report_hash"`
	FactsEnvironment             string      `json:"facts_environment"`
	CachedCatalogStatus          string      `json:"cached_catalog_status"`
	ReportEnvironment            string      `json:"report_environment"`
	LatestReportCorrectiveChange bool        `json:"latest_report_corrective_change"`
	CatalogEnvironment           string      `json:"catalog_environment"`
	FactsTimestamp               string      `json:"facts_timestamp"`
	LatestReportNoop             bool        `json:"latest_report_noop"`
	Expired                      interface{} `json:"expired"`
	LatestReportNoopPending      bool        `json:"latest_report_noop_pending"`
	ReportTimestamp              string      `json:"report_timestamp"`
	Certname                     string      `json:"certname"`
	CatalogTimestamp             string      `json:"catalog_timestamp"`
	LatestReportJobID            string      `json:"latest_report_job_id"`
	LatestReportStatus           string      `json:"latest_report_status"`
}

// Pagination is a filter to be used when paginating
type Pagination struct {
	Limit        int
	Offset       int
	IncludeTotal bool
}
