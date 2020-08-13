package main

import (
	"crypto/tls"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	prompt "github.com/c-bata/go-prompt"
	"github.com/puppetlabs/go-pe-client/internal/cli"
	"github.com/puppetlabs/go-pe-client/pkg/puppetdb"
	"github.com/sirupsen/logrus"
)

var client *puppetdb.Client
var prompter *prompt.Prompt
var historyFile *os.File
var pdbTimeout = time.Second * 30

var suggestions = []prompt.Suggest{
	//  Methods
	{Text: "nodes", Description: "Get nodes"},
	{Text: "facts", Description: "Get facts"},
	{Text: "factnames", Description: "Get fact names"},
	{Text: "inventory", Description: "Get inventory"},
	{Text: "reports", Description: "Get reports"},

	// Binary Operators
	{Text: "=", Description: "equal to"},
	{Text: ">", Description: "greater than"},
	{Text: "<", Description: "less than"},
	{Text: ">=", Description: "greater than or equal to"},
	{Text: "<=", Description: "less than or equal to"},
	{Text: "~", Description: "regexp match"},
	{Text: "~>", Description: "regexp array match"},
	{Text: "null?", Description: "is null"},

	// Boolean Operators
	{Text: "and", Description: ""},
	{Text: "or", Description: ""},
	{Text: "not", Description: ""},

	// Projection Operators
	{Text: "extract", Description: "To reduce the keypairs returned for each result in the response, you can use extract:"},

	// Command
	{Text: "exit", Description: "Exit pdb"},
}

func executor(in string) {
	in = strings.TrimSpace(in)

	// Parse the input and extract the API call + query
	var api, query, pagination, orderBy = cli.ParseInput(in)
	// If a api has been selected, then execute it with the provided query
	// the command should be recorded in history and the response printed to
	// stdout
	if api != "" {
		err := cli.WriteHistory(historyFile, in)
		if err != nil {
			logrus.Warnf("Unable to write history to %s because : %s", historyFile.Name(), err)
		}
		execute(api, query, pagination, orderBy)
	}
}

func execute(api string, query string, pagination puppetdb.Pagination, orderBy puppetdb.OrderBy) {
	fmt.Printf("Executing Query '%s %s'\n", api, query)
	var err error
	var data interface{}

	switch api {
	case "nodes":
		fmt.Printf("Nodes")
		data, err = client.Nodes(query, &pagination, &orderBy)
	case "facts":
		data, err = client.Facts(query, &pagination, &orderBy)
	case "inventory":
		data, err = client.Inventory(query, &pagination, &orderBy)
	case "reports":
		data, err = client.Reports(query, &pagination, &orderBy)
	case "factnames":
		data, err = client.FactNames(&pagination, &orderBy)
	}

	if err != nil {
		fmt.Println("err: " + err.Error())
		return
	}
	cli.PrintString(data)
}

func completer(in prompt.Document) []prompt.Suggest {
	w := in.GetWordBeforeCursor()
	if w == "" {
		return []prompt.Suggest{}
	}
	return prompt.FilterHasPrefix(suggestions, w, true)
}

// process Args and create a PDB client
func processArgs() (*puppetdb.Client, error) {
	if len(os.Args) < 3 {
		fmt.Println("\tusage: pdb pe.puppetlabs.net aabbccddeeff")
		os.Exit(-1)
	}
	peServer := os.Args[1]
	token := os.Args[2]
	pdbHostURL := "https://" + peServer + ":8081"

	u, err := url.Parse(pdbHostURL)
	if err != nil {
		return nil, err
	}
	pdb := puppetdb.NewClient(u.String(), token, &tls.Config{InsecureSkipVerify: true}, pdbTimeout) // #nosec - this tool is private and for development purpose
	return pdb, nil
}

func main() {

	// Process args and create a context with PDB client
	var err error
	client, err = processArgs()
	if err != nil {
		logrus.Fatal(err)
	}

	// Initialize history file, where we store the command history
	historyFile, err = cli.InitHistoryFile()
	if err != nil {
		logrus.Warnf("Unable to create history file because : %s", err)
	}
	history := cli.ReadHistory(historyFile)
	err = historyFile.Close()
	if err != nil {
		logrus.Warnf("failed to close history file: %s", err)
	}

	prompter = prompt.New(
		executor,
		completer,
		prompt.OptionPrefix("pdb> "),
		prompt.OptionTitle("puppet-db"),
		prompt.OptionHistory(history),
	)
	prompter.Run()
}
