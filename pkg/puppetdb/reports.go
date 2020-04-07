package puppetdb

import (
	"time"
)

const (
	reports = "/pdb/query/v4/reports"
)

// Reports retrieve the reports that Puppet agent nodes submit after their runs. The Puppet master forwards these to PuppetDB. Each report includes:
// Data about the entire run
// Metadata about the report
// Many events, describing what happened during the run
func (c *Client) Reports(query string, pagination *Pagination) ([]Report, error) {
	payload := []Report{}
	err := getRequest(c, reports, query, pagination, &payload)
	return payload, err
}

// Report summaries for all event reports that matched the input parameters.
type Report struct {
	Hash                 string         `json:"hash"`
	PuppetVersion        string         `json:"puppet_version"`
	ReceiveTime          time.Time      `json:"receive_time"`
	ReportFormat         int            `json:"report_format"`
	StartTime            time.Time      `json:"start_time"`
	EndTime              time.Time      `json:"end_time"`
	ProducerTimestamp    time.Time      `json:"producer_timestamp"`
	Producer             string         `json:"producer"`
	TransactionUUID      string         `json:"transaction_uuid"`
	Status               string         `json:"status"`
	Noop                 bool           `json:"noop"`
	NoopPending          bool           `json:"noop_pending"`
	Environment          string         `json:"environment"`
	ConfigurationVersion string         `json:"configuration_version"`
	Certname             string         `json:"certname"`
	CodeID               string         `json:"code_id"`
	CatalogUUID          string         `json:"catalog_uuid"`
	CachedCatalogStatus  string         `json:"cached_catalog_status"`
	ResourceEvents       ResourceEvents `json:"resource_events"`
	Resources            Resources      `json:"resources"`
	Metrics              Metrics        `json:"metrics"`
	Logs                 Logs           `json:"logs"`
}

// ResourceEvents ...
type ResourceEvents struct {
	Href string
	Data []struct {
		Status          string      `json:"status"`
		Timestamp       time.Time   `json:"timestamp"`
		ResourceType    string      `json:"resource_type"`
		ResourceTitle   string      `json:"resource_title"`
		Property        string      `json:"property"`
		Name            string      `json:"name"`
		NewValue        interface{} `json:"new_value"`
		OldValue        interface{} `json:"old_value"`
		Message         string      `json:"message"`
		File            string      `json:"file"`
		Line            int         `json:"line"`
		ContainmentPath []string    `json:"containment_path"`
	}
}

// Resources ...
type Resources struct {
	Href string
	Data []struct {
		Timestamp       time.Time `json:"timestamp"`
		ResourceType    string    `json:"resource_type"`
		ResourceTitle   string    `json:"resource_title"`
		ContainmentPath []string  `json:"containment_path"`
		Skipped         bool      `json:"skipped"`
		Events          []Event   `json:"events"`
	}
}

// Event ...
type Event struct {
	Timestamp time.Time   `json:"timestamp"`
	Property  string      `json:"property"`
	Name      string      `json:"name"`
	NewValue  interface{} `json:"new_value"`
	OldValue  interface{} `json:"old_value"`
	Message   string      `json:"message"`
	Status    string      `json:"status"`
}

// Metrics ...
type Metrics struct {
	Href string
	Data []struct {
		Category string  `json:"category"`
		Name     string  `json:"name"`
		Value    float32 `json:"value"`
	}
}

// Logs returns a single log line per data entry.
// File and line may each be null if the log does not concern a resource.
type Logs struct {
	Href string
	Data []struct {
		File    string    `json:"file"`
		Line    int       `json:"line"`
		Level   string    `json:"level"`
		Message string    `json:"message"`
		Source  string    `json:"source"`
		Tags    []string  `json:"tags"`
		Time    time.Time `json:"time"`
	}
}
