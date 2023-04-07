package puppetdb

import (
	"fmt"
	"io"
	"math"
	"strconv"
)

// toParams will take the Pagination struct and convert into a form Client SetQueryParam accepts
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

// Pagination is a filter to be used when paginating
type Pagination struct {
	Limit        int
	Offset       int
	IncludeTotal bool
	Total        int
}

func NewDefaultPagination() *Pagination {
	return &Pagination{
		Limit:        100,
		IncludeTotal: true,
	}
}

// pageCursor is a pagination cursor that provides convenience methods for
// stepping through pbd response pages.
type pageCursor struct {
	client      *Client
	path        string
	query       string
	pagination  *Pagination
	orderBy     *OrderBy
	currentPage any
}

// Next returns a page of nodes and iterates the pagination cursor by the
// offset. If there are no more results left, the error will be io.EOF.
func (pc *pageCursor) next(response any) error {
	// this block increases the offset and checks of it's greater than or equal
	// to the total only if we have already returned a first page.
	if pc.currentPage != nil {
		pc.pagination.Offset = pc.pagination.Offset + pc.pagination.Limit

		if pc.pagination.Offset >= pc.pagination.Total {
			return io.EOF
		}
	} else {
		if pc.pagination == nil {
			pc.pagination = NewDefaultPagination()
		}

		pc.pagination.IncludeTotal = true
	}

	var err error

	err = getRequest(pc.client, pc.path, pc.query, pc.pagination, pc.orderBy, response)
	if err != nil {
		return fmt.Errorf("page cursor: client call returned an error: %w", err)
	}

	pc.currentPage = response

	if pc.CurrentPage() == pc.TotalPages() {
		err = io.EOF
	}

	return err
}

// TotalPages returns the total number of pages that can returns nodes.
func (pc *pageCursor) TotalPages() int {
	pagesf := float64(pc.pagination.Total) / float64(pc.pagination.Limit)
	pages := int(math.Ceil(pagesf))

	return pages
}

// CurrentPage returns the current page number the cursor is at.
func (pc *pageCursor) CurrentPage() int {
	if pc.pagination.Offset == 0 {
		return 1
	}

	return pc.pagination.Offset/pc.pagination.Limit + 1
}

func newPageCursor(c *Client, path, query string, p *Pagination, orderBy *OrderBy) (*pageCursor, error) {
	if p == nil {
		p = NewDefaultPagination()
	}

	tempPagination := Pagination{
		Limit:        1,
		IncludeTotal: true,
	}

	// make a call to pdb for 1 object to fetch the total number of results for
	// page calculations in the cursor.
	if err := getRequest(c, path, query, &tempPagination, orderBy, &[]any{}); err != nil {
		return nil, fmt.Errorf("failed to get result total from pdb: %w", err)
	}

	p.Total = tempPagination.Total
	p.IncludeTotal = true

	pc := &pageCursor{
		client:     c,
		path:       path,
		query:      query,
		pagination: p,
		orderBy:    orderBy,
	}

	return pc, nil
}
