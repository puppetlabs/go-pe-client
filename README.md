## Go client library for the Puppet Enterprise APIs

[![Go Report Card](https://goreportcard.com/badge/github.com/puppetlabs/go-pe-client)](https://goreportcard.com/report/github.com/puppetlabs/go-pe-client) ![Travis Build Status](https://travis-ci.com/puppetlabs/go-pe-client.svg?branch=master)

Currently a small subset of the Orchestrator and PuppetDB APIs are supported:
* https://puppet.com/docs/pe/latest/orchestrator_api_usage_endpoint.html
* https://puppet.com/docs/puppetdb/latest/api/index.html
* https://puppet.com/docs/pe/latest/rbac_api_v1.html

## Running the command line
* Assuming you have go installed. Run the following to get an RBAC token from a PE device:

`go run cmd/test/main.go <pe-server> <login> <password>`

EG

```
go run cmd/test/main.go cometary-plot.delivery.puppetlabs.net admin pazzw0rd

Connecting to: cometary-plot.delivery.puppetlabs.net
(*rbac.Token)(0xc000098de0)({
 Token: (string) (len=44) "0OVHPFr4izm980Ll2g0eVikOm53wBizGHdMJ3cbF_8IM"
})

```

* Then you can see example calls to the Orchestrator and PuppetDB:

`go run cmd/test/main.go <pe-server> <token>`

EG

```
go run cmd/test/main.go brown-solidity.delivery.puppetlabs.net APz9o8g392dTK0yEIOLk1Vl-rr2fVWGp0mnhgFH52PMf

Connecting to:  brown-solidity.delivery.puppetlabs.net
(*[]puppetdb.Node)(0xc0000a63c0)((len=7 cap=9) {
 (puppetdb.Node) {
  Deactivated: (interface {}) <nil>,
  LatestReportHash: (string) (len=40) "2a641f45f7497719ebc07e43603aa3bef7d08dd4",
  FactsEnvironment: (string) (len=10) "production",
  CachedCatalogStatus: (string) (len=8) "not_used",
  ReportEnvironment: (string) (len=10) "production",
  LatestReportCorrectiveChange: (bool) false,
  CatalogEnvironment: (string) (len=10) "production",
  FactsTimestamp: (string) (len=24) "2020-03-19T09:20:41.172Z",
  LatestReportNoop: (bool) false,
  Expired: (interface {}) <nil>,
  LatestReportNoopPending: (bool) false,
  ReportTimestamp: (string) (len=24) "2020-03-19T09:20:42.839Z",
  Certname: (string) (len=38) "safest-thicket.delivery.puppetlabs.net",
  CatalogTimestamp: (string) (len=24) "2020-03-19T09:20:42.669Z",
  LatestReportJobID: (string) (len=1) "5",
  LatestReportStatus: (string) (len=9) "unchanged"
 },
 (puppetdb.Node) {
  Deactivated: (interface {}) <nil>,
  LatestReportHash: (string) (len=40) "4824b1af8b06a49a602bf471b574c6a7f32bd85d",
  FactsEnvironment: (string) (len=10) "production",
  CachedCatalogStatus: (string) (len=8) "not_used",
  ReportEnvironment: (string) (len=10) "production",
  LatestReportCorrectiveChange: (bool) false,
  CatalogEnvironment: (string) (len=10) "production",
  FactsTimestamp: (string) (len=24) "2020-03-20T15:28:05.463Z",
  LatestReportNoop: (bool) false,
  Expired: (interface {}) <nil>,
  LatestReportNoopPending: (bool) false,
  ReportTimestamp: (string) (len=24) "2020-03-20T15:28:28.850Z",
  Certname: (string) (len=38) "brown-solidity.delivery.puppetlabs.net",
  CatalogTimestamp: (string) (len=24) "2020-03-20T15:28:09.172Z",
  LatestReportJobID: (string) "",
  LatestReportStatus: (string) (len=9) "unchanged"
 }
})

(*[]puppetdb.Node)(0xc0000a6960)((len=1 cap=4) {
 (puppetdb.Node) {
  Deactivated: (interface {}) <nil>,
  LatestReportHash: (string) (len=40) "4824b1af8b06a49a602bf471b574c6a7f32bd85d",
  FactsEnvironment: (string) (len=10) "production",
  CachedCatalogStatus: (string) (len=8) "not_used",
  ReportEnvironment: (string) (len=10) "production",
  LatestReportCorrectiveChange: (bool) false,
  CatalogEnvironment: (string) (len=10) "production",
  FactsTimestamp: (string) (len=24) "2020-03-20T15:28:05.463Z",
  LatestReportNoop: (bool) false,
  Expired: (interface {}) <nil>,
  LatestReportNoopPending: (bool) false,
  ReportTimestamp: (string) (len=24) "2020-03-20T15:28:28.850Z",
  Certname: (string) (len=38) "brown-solidity.delivery.puppetlabs.net",
  CatalogTimestamp: (string) (len=24) "2020-03-20T15:28:09.172Z",
  LatestReportJobID: (string) "",
  LatestReportStatus: (string) (len=9) "unchanged"
 }
})

(*[]orch.InventoryNode)(0xc0000a6ba0)((len=7 cap=9) {
 (orch.InventoryNode) {
  Name: (string) (len=39) "ancestral-flora.delivery.puppetlabs.net",
  Connected: (bool) true,
  Broker: (string) (len=51) "pcp://brown-solidity.delivery.puppetlabs.net/server",
  Timestamp: (string) (len=24) "2020-03-19T08:58:08.335Z"
 },
 (orch.InventoryNode) {
  Name: (string) (len=39) "taxable-preview.delivery.puppetlabs.net",
  Connected: (bool) true,
  Broker: (string) (len=51) "pcp://brown-solidity.delivery.puppetlabs.net/server",
  Timestamp: (string) (len=24) "2020-03-19T08:58:03.314Z"
 }
})
```

* Also you can see calls to the Classifier:

To hit the classifier node endpoint run :

`go run cmd/test/main.go classifier node <pe-server> <nodename> <token>`

To hit the classifier group endpoint run :

`go run cmd/test/main.go classifier groups <pe-server> <token>`

EG

```
(classifier.Node) {
 Name: (string) (len=37) "anomic-cabana.delivery.puppetlabs.net",
 Environment: (string) (len=10) "production",
 Groups: ([]string) (len=3 cap=4) {
  (string) (len=36) "00000000-0000-4000-8000-000000000000",
  (string) (len=36) "805b4e23-6bb7-445f-914f-1fe3ea48239f",
  (string) (len=36) "c4513bd2-a792-4060-aeff-a3c511329e4b"
 },
 Classes: (struct {}) {
 },
 Parameters: (struct {}) {
 },
 ConfigData: (struct {}) {
 }
}
```

## Examples

There are examples available to run to show how the PE client hangs together.

```bash
examples/
```

This is running the classifer example

```bash
# greghardy @ UisceBeatha in ~/go/src/github.com/puppetlabs/go-pe-client on git:classifier_node x [9:10:07]
$ examples/classifier.sh custom-kimono.delivery.puppetlabs.net
Aquiring token
Token=AOHJvizy1My65MFDq1METTq6X-f5nqlvLMc_-L1yMxR9
Selecting a node from PDB
Selecting a classified node from the result=irate-sphere.delivery.puppetlabs.net
Performing the classifier API call on node=irate-sphere.delivery.puppetlabs.net
Responses=(classifier.Node) {
 Name: (string) (len=36) "irate-sphere.delivery.puppetlabs.net",
 Environment: (string) (len=10) "production",
 Groups: ([]struct { ID string "json:\"id\""; Name string "json:\"name\"" }) (len=2 cap=4) {
  (struct { ID string "json:\"id\""; Name string "json:\"name\"" }) {
   ID: (string) (len=36) "173a9bba-c683-4d10-a2d8-16aa7208fd42",
   Name: (string) (len=16) "All Environments"
  },
  (struct { ID string "json:\"id\""; Name string "json:\"name\"" }) {
   ID: (string) (len=36) "00000000-0000-4000-8000-000000000000",
   Name: (string) (len=9) "All Nodes"
  }
 },
 Classes: (struct {}) {
 },
 Parameters: (struct {}) {
 },
 ConfigData: (struct {}) {
 }
}
```

## Supporting new endpoints

To add support for a new endpoint, you can:
1. Add an example payload from the API docs to the testdata dir (or if you have real data add it to a separate subdir)
2. Use one of the existing implementation files and corresponding tests as a template for your new endpoint (e.g. `inventory.go` and `inventory_test.go`)
3. Create a struct to represent your example payload, perhaps using a tool like https://mholt.github.io/json-to-go
4. Update the resty calls as needed
5. Update tests as needed
