package puppetdb

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNodes(t *testing.T) {
	// Test without query
	setupGetResponder(t, "/pdb/query/v4/nodes", "", "nodes-response.json")
	actual, err := pdbClient.Nodes("", nil, nil)
	require.Nil(t, err)
	require.Equal(t, expectedNodes, actual)

	// Test with query
	query := `["=", "certname", "lenient-veranda.delivery.puppetlabs.net"]`
	setupGetResponder(t, "/pdb/query/v4/nodes", "query="+query, "nodes-response.json")
	actual, err = pdbClient.Nodes(query, nil, nil)
	require.Nil(t, err)
	require.Equal(t, expectedNodes, actual)
}

func TestNode(t *testing.T) {
	nodeFooURL := strings.ReplaceAll(node, "{certname}", "foo")

	// Test success
	setupGetResponder(t, nodeFooURL, "", "node-response.json")
	actual, err := pdbClient.Node("foo")
	require.Nil(t, err)
	require.Equal(t, expectedNode, actual)
}

var expectedNodes = []Node{{Deactivated: interface{}(nil), LatestReportHash: "7ccb6fb17b3fe11cecffe00b43b44f3776bcb89d", FactsEnvironment: "production", CachedCatalogStatus: "not_used", ReportEnvironment: "production", LatestReportCorrectiveChange: false, CatalogEnvironment: "production", FactsTimestamp: "2020-03-20T10:17:30.394Z", LatestReportNoop: false, Expired: interface{}(nil), LatestReportNoopPending: false, ReportTimestamp: "2020-03-20T10:17:54.470Z", Certname: "lenient-veranda.delivery.puppetlabs.net", CatalogTimestamp: "2020-03-20T10:17:33.991Z", LatestReportJobID: "1", LatestReportStatus: "changed"}, {Deactivated: interface{}(nil), LatestReportHash: "", FactsEnvironment: "production", CachedCatalogStatus: "", ReportEnvironment: "", LatestReportCorrectiveChange: false, CatalogEnvironment: "", FactsTimestamp: "2020-03-20T10:10:28.949Z", LatestReportNoop: false, Expired: interface{}(nil), LatestReportNoopPending: false, ReportTimestamp: "", Certname: "inland-ancestor.delivery.puppetlabs.net", CatalogTimestamp: "", LatestReportJobID: "", LatestReportStatus: ""}}

var expectedNode = &Node{Deactivated: interface{}(nil), LatestReportHash: "", FactsEnvironment: "production", CachedCatalogStatus: "", ReportEnvironment: "", LatestReportCorrectiveChange: false, CatalogEnvironment: "", FactsTimestamp: "2020-03-20T10:10:28.949Z", LatestReportNoop: false, Expired: interface{}(nil), LatestReportNoopPending: false, ReportTimestamp: "", Certname: "inland-ancestor.delivery.puppetlabs.net", CatalogTimestamp: "", LatestReportJobID: "", LatestReportStatus: "", Count: 0}
