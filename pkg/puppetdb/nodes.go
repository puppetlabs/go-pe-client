package puppetdb

import (
	"fmt"
	"io"
	"math"
	"strings"
)

const (
	nodes = "/pdb/query/v4/nodes"
	node  = "/pdb/query/v4/nodes/{certname}"
)

// Nodes will return all nodes matching the given query. Deactivated and expired nodes aren’t included in the response.
func (c *Client) Nodes(query string, pagination *Pagination, orderBy *OrderBy) ([]Node, error) {
	payload := []Node{}
	err := getRequest(c, nodes, query, pagination, orderBy, &payload)
	return payload, err
}

// PaginatedNodes works just like Nodes, but returns a NodesCursor that
// provides methods for iterating over N pages of nodes and calculates page
// information for tracking progress. If pagination is nil, then a default
// configuration with a limit of 100 is used instead.
func (c *Client) PaginatedNodes(query string, pagination *Pagination, orderBy *OrderBy) (*NodesCursor, error) {
	if pagination == nil {
		pagination = &Pagination{Limit: 100}
	}

	tempPagination := Pagination{
		Limit:        1,
		IncludeTotal: true,
	}

	// make a call to pdb for 1 node to fetch the total number of nodes for
	// page calculations in the cursor.
	if _, err := c.Nodes(query, &tempPagination, orderBy); err != nil {
		return nil, fmt.Errorf("failed to get node total from pdb: %w", err)
	}

	pagination.Total = tempPagination.Total
	pagination.IncludeTotal = true

	nc := &NodesCursor{
		client:     c,
		pagination: pagination,
		query:      query,
		orderBy:    orderBy,
	}

	return nc, nil
}

// Node will return a single node by certname
func (c *Client) Node(certname string) (*Node, error) {
	payload := &Node{}
	r, err := c.resty.R().
		SetResult(&payload).
		SetPathParams(map[string]string{"certname": certname}).
		Get(node)
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		if r.Error() != nil {
			return nil, r.Error().(error)
		}
		return nil, fmt.Errorf("%s error: %s", strings.ReplaceAll(node, "{certname}", certname), r.Status())
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
	Count                        int         `json:"count"`
}

// NodesCursor is a pagination cursor that provides convenience methods for
// stepping through pages of nodes.
type NodesCursor struct {
	client      *Client
	pagination  *Pagination
	query       string
	orderBy     *OrderBy
	currentPage []Node
}

// Next returns a page of nodes and iterates the pagination cursor by the
// offset. If there are no more results left, the error will be io.EOF.
func (nc *NodesCursor) Next() ([]Node, error) {
	// this block increases the offset and checks of it's greater than or equal
	// to the total only if we have already returned a first page.
	if nc.currentPage != nil {
		nc.pagination.Offset = nc.pagination.Offset + nc.pagination.Limit

		if nc.pagination.Offset >= nc.pagination.Total {
			return []Node{}, io.EOF
		}
	}

	var err error

	nc.currentPage, err = nc.client.Nodes(nc.query, nc.pagination, nc.orderBy)
	if err != nil {
		return nil, fmt.Errorf("client call for Nodes returned an error: %w", err)
	}

	if nc.CurrentPage() == nc.TotalPages() {
		err = io.EOF
	}

	return nc.currentPage, err
}

// TotalPages returns the total number of pages that can returns nodes.
func (nc *NodesCursor) TotalPages() int {
	pagesf := float64(nc.pagination.Total) / float64(nc.pagination.Limit)
	pages := int(math.Ceil(pagesf))

	return pages
}

// CurrentPage returns the current page number the cursor is at.
func (nc *NodesCursor) CurrentPage() int {
	if nc.pagination.Offset == 0 {
		return 1
	}

	return nc.pagination.Offset/nc.pagination.Limit + 1
}
