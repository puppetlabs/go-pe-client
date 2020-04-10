package cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"regexp"
	"strconv"
	"strings"

	"github.com/puppetlabs/go-pe-client/pkg/puppetdb"
)

func InitHistoryFile() (*os.File, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("Unable to get users home directory.   The command history wont be saved.")
	}

	filename := fmt.Sprintf("%s/.pdb_history", usr.HomeDir)
	return os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0655)
}

func WriteHistory(historyFile *os.File, cmd string) error {
	if historyFile != nil {
		_, err := historyFile.WriteString(fmt.Sprintf("%s\n", cmd))
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadHistory(historyFile *os.File) []string {
	var lines []string
	if historyFile != nil {
		scanner := bufio.NewScanner(historyFile)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		return lines
	}
	return lines
}

func PrintString(data interface{}) {
	jsonString, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println("err: " + err.Error())
	}
	fmt.Println(string(jsonString))
}

func extractString(str, start, end string) (result string) {
	s := strings.Index(str, start)
	s--
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.LastIndex(str, end)
	e++
	if e == -1 {
		return
	}
	return str[s:e]
}

func createPaginationStruct(options []string) (result puppetdb.Pagination) {
	pagination := puppetdb.Pagination{}
	var err interface{}
	for _, option := range options {

		re := regexp.MustCompile(`Limit=(\d*)`)
		limit := re.FindStringSubmatch(option)

		re2 := regexp.MustCompile(`Offset=(\d*)`)
		offset := re2.FindStringSubmatch(option)

		re3 := regexp.MustCompile(`Include_total=(\d*)`)
		total := re3.FindStringSubmatch(option)

		if len(limit) >= 1 {
			pagination.Limit, err = strconv.Atoi(limit[1])
		}
		if len(offset) >= 1 {

			pagination.Offset, err = strconv.Atoi(offset[1])
		}
		if len(total) >= 1 {
			fmt.Println("Include_total is currently not implemented")
		}
		if err != nil {
			fmt.Println(err)
		}
	}
	return pagination
}

// ParseInput is used to split the users entered query into the relevant parts and returns each part as a string
// For example "nodes ["=", "certname", "jenkins-compose.example.net"] Limit=5 Offset=10"
// Would return "nodes" "["=", "certname", "jenkins-compose.example.net"]" "Limit=5 Offset=10"
// "nodes", "nodes Limit=10", "nodes []", "nodes ["=", "certname", "jenkins-compose.example.net"]" are all accepted by this func
func ParseInput(command string) (string, string, puppetdb.Pagination) {

	checkForQuery, err := regexp.Match(`[\w+]`, []byte(command))
	var query string
	if checkForQuery {
		query = extractString(command, "[", "]")
		if query == "[]" {
			query = ""
		}
	}
	if err != nil {
		fmt.Println("No query parameters detected", err)
	}

	querylessCommand := strings.Replace(command, query, "", 1)
	blocks := strings.Split(querylessCommand, " ")
	api := blocks[0]

	if api == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
	}

	var rex = regexp.MustCompile(`(\w+)=(\w+)`) //nolint
	options := rex.FindAllString(querylessCommand, -1)
	pagination := createPaginationStruct(options)

	return api, query, pagination
}
