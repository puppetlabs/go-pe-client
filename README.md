## Go client library for the Puppet Enterprise APIs

[![Go Report Card](https://goreportcard.com/badge/github.com/puppetlabs/go-pe-client)](https://goreportcard.com/report/github.com/puppetlabs/go-pe-client) ![Travis Build Status](https://travis-ci.com/puppetlabs/go-pe-client.svg?branch=master)

Currently a small subset of the Orchestrator and PuppetDB APIs are supported:
* https://puppet.com/docs/pe/latest/orchestrator_api_usage_endpoint.html
* https://puppet.com/docs/puppetdb/latest/api/index.html

## Running the command line
* Go to your pe server cli and generate an rbac token. Be sure to set the correct username and password.

`curl -k -X POST -H 'Content-Type: application/json' -d '{"login": "admin", "password": "compliance", "lifetime": "4h", "label": "four-hour token"}' https://localhost:4433/rbac-api/v1/auth/token
{"token":"APz9o8g392dTK0yEIOLk1Vl-rr2fVWGp0mnhgFH52PMf"}`

* Assuming you have go installed. Run the following:

` go run cmd/main.go <pe-server> <token>`

EG

```
go run cmd/main.go brown-solidity.delivery.puppetlabs.net APz9o8g392dTK0yEIOLk1Vl-rr2fVWGp0mnhgFH52PMf

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

## Supporting new endpoints

To add support for a new endpoint, you can:
1. Add an example payload from the API docs to the testdata dir (or if you have real data add it to a separate subdir)
2. Use one of the existing implementation files and corresponding tests as a template for your new endpoint (e.g. `inventory.go` and `inventory_test.go`)
3. Create a struct to represent your example payload, perhaps using a tool like https://mholt.github.io/json-to-go
4. Update the resty calls as needed
5. Update tests as needed
