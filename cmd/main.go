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

	pdbClient := puppetdb.NewInsecureClient(pdbHost, token)
	nodes, err := pdbClient.Nodes("")
	if err != nil {
		panic(err)
	}
	spew.Dump(nodes)
	fmt.Println()

	nodes, err = pdbClient.Nodes("[\"=\", \"certname\", \"" + peServer + "\"]")
	if err != nil {
		panic(err)
	}
	spew.Dump(nodes)
	fmt.Println()

	orchClient := orch.NewInsecureClient(orchHost, token)
	inv, err := orchClient.Inventory()
	if err != nil {
		panic(err)
	}
	spew.Dump(inv)
	fmt.Println()

}
