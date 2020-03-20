package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/puppetlabs/go-pe-client/orch"
	"github.com/puppetlabs/go-pe-client/puppetdb"
)

func main() {

	pdbHost := "https://lenient-veranda.delivery.puppetlabs.net:8081"
	orchHost := "https://lenient-veranda.delivery.puppetlabs.net:8143"
	token := "xxxx"

	pdbClient := puppetdb.NewInsecureClient(pdbHost, token)
	nodes, err := pdbClient.Nodes("")
	if err != nil {
		panic(err)
	}
	spew.Dump(nodes)
	fmt.Println()

	nodes, err = pdbClient.Nodes(`["=", "certname", "lenient-veranda.delivery.puppetlabs.net"]`)
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
