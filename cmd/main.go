package main

import (
 	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/puppetlabs/go-pe-client/orch"
	"github.com/puppetlabs/go-pe-client/puppetdb"
)

func main() {
	peServer := os.Args[1]
	token := os.Args[2]
	pdbHost := "https://" + peServer + ":8081"
	orchHost := "https://" + peServer + ":8143"
	fmt.Println("Connecting to: ", peServer)

	pdbClient := puppetdb.NewInsecureClient(pdbHostURL, token)
	nodes, err := pdbClient.Nodes("")
	if err != nil {
		panic(err)
	}
	spew.Dump(nodes)
	fmt.Println()

	nodes, err = pdbClient.Nodes(fmt.Sprintf(`["=", "certname", "%s"]`, peServer))
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
