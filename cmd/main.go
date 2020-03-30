package main

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/puppetlabs/go-pe-client/pkg/orch"
	"github.com/puppetlabs/go-pe-client/pkg/puppetdb"
)

func main() {

	if len(os.Args) < 3 {
		panic("usage: go run cmd/main.go <pe-server> <token> e.g. go run cmd/main.go pe.puppetlabs.net aabbccddeeff")
	}

	peServer := os.Args[1]
	token := os.Args[2]
	pdbHostURL := "https://" + peServer + ":8081"
	orchHostURL := "https://" + peServer + ":8143"
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
		Scope: orch.Scope{
			Nodes: []string{peServer},
		},
	})
	if err != nil {
		panic(err)
	}
	spew.Dump(job)
	fmt.Println()

}
