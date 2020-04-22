package puppetdb

import "fmt"

func (o OrderBy) toOrderParam() map[string]string {
	orderPram := map[string]string{}

	orderPram["order_by"] = fmt.Sprintf(`[{"field": "%s", "order": "%s"}]`, o.Field, o.Order)

	return orderPram
}

// OrderBy is used to determine a responses ordering
type OrderBy struct {
	Field string
	Order string
}
