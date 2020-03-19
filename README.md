## Go client library for the Puppet Enterprise APIs

Currently only a small subset of the Orchestrator API is supported: https://puppet.com/docs/pe/latest/orchestrator_api_usage_endpoint.html

* `/orchestrator/v1/inventory`
* `/orchestrator/v1/tasks`

### Supporting new endpoints

To add support for a new endpoint, you can:
1. Add an example payload from the API docs to the testdata dir (or if you have real data add it to a separate subdir)
2. Copy one of the existing implementation files (e.g. `inventory.go`) for your new endpoint
3. Create a struct to represent your example payload. You can use a tool like https://mholt.github.io/json-to-go/ but be sure to add the 'omitempty' constraint to all fields (or at least all optional fields)
4. Update the resty call to point to the correct URL
5. Add a test to `orch_test.go` using one of the existing tests as a template (e.g. `TestInventory`)
