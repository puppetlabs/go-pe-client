package puppetdb

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// TestReports performs a test on the Report endpoint and verifies the expected response is returned.
func TestReports(t *testing.T) {
	query := `["=", "certname", "foobar.delivery.puppetlabs.net"]`
	setupGetResponder(t, reports, "query="+query, "reports.json")
	actual, err := pdbClient.Reports(query)
	require.Nil(t, err)
	require.Equal(t, expectedReport, actual)
}

var expectedReport = []Report{{
	Hash:                 "4324324324324324324",
	PuppetVersion:        "10.0",
	ReceiveTime:          time.Time{},
	ReportFormat:         2.0,
	StartTime:            time.Time{},
	EndTime:              time.Time{},
	ProducerTimestamp:    time.Time{},
	Producer:             "foobar-master.puppet.com",
	TransactionUUID:      "342343432432",
	Status:               "Good",
	Noop:                 false,
	NoopPending:          false,
	Environment:          "production",
	ConfigurationVersion: "99",
	Certname:             "foobar",
	CodeID:               "2343",
	CatalogUUID:          "343243243243243243",
	CachedCatalogStatus:  "Good",
	ResourceEvents: ResourceEvents{
		Href: "http://foobar.events.com",
	},
	Resources: Resources{
		Href: "http://foobar.resources.com",
	},
	Metrics: Metrics{
		Href: "http://foobar.metrics.com",
	},
	Logs: Logs{
		Href: "http://foobar.logs.com",
	}}}
