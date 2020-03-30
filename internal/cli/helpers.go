package cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/user"
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
