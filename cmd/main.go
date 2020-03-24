package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/puppetlabs/go-pe-client/orch"
	"github.com/puppetlabs/go-pe-client/puppetdb"
)

func main() {

	peHost := "lenient-veranda.delivery.puppetlabs.net"
	pdbHostURL := fmt.Sprintf("https://%s:8081", peHost)
	orchHostURL := fmt.Sprintf("https://%s:8143", peHost)
	token := "0UoFj-u-34x3O7o3AJe3ezh8NYJZ9MWIwhuFoFLEKLiM"

	pdbClient := puppetdb.NewInsecureClient(pdbHostURL, token)
	nodes, err := pdbClient.Nodes("")
	if err != nil {
		panic(err)
	}
	spew.Dump(nodes)
	fmt.Println()

	nodes, err = pdbClient.Nodes(fmt.Sprintf(`["=", "certname", "%s"]`, peHost))
	if err != nil {
		panic(err)
	}
	spew.Dump(nodes)
	fmt.Println()

	orchClient := orch.NewInsecureClient(orchHostURL, token)
	inv, err := orchClient.Inventory()
	if err != nil {
		panic(err)
	}
	spew.Dump(inv)
	fmt.Println()

	job, err := orchClient.CommandTask(&orch.TaskRequest{
		Task: "package",
		Params: map[string]string{
			"action": "status",
			"name":   "openssl",
		},
		Scope: orch.TaskScope{
			Nodes: []string{peHost},
		},
	})
	if err != nil {
		panic(err)
	}
	spew.Dump(job)
	fmt.Println()

}
