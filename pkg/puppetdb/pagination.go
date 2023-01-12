package puppetdb

import "strconv"

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
