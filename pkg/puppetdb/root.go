package puppetdb

import (
	"errors"
	"fmt"
	"io"
)

const (
	rootQueryEndpoint = "/pdb/query/v4"
)

func (c *Client) PaginatedRootQuery(query string, pagination *Pagination, orderBy *OrderBy) (*RootQueryCursor, error) {
	pc, err := newPageCursor(c, rootQueryEndpoint, query, pagination, orderBy)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize page cursor: %w", err)
	}

	cursor := RootQueryCursor{
		pageCursor: pc,
	}

	return &cursor, nil
}

type RootQueryCursor struct {
	*pageCursor
}

func (rqc *RootQueryCursor) NextInto(target any) error {
	err := rqc.next(target)
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}

	return err
}
