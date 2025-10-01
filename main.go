package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"time"
)

type history struct {
	code      string
	timeStamp time.Time
}

var codeHistory []history

func main() {
	fmt.Println("-=Marsupilami=- write exit to leave")
	reader := bufio.NewReader(os.Stdin)
	var code strings.Builder

	for {
		fmt.Print(getPrompt())
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		switch line {
		case "exit":
			return
		case "return":
			if code.Len() > 0 {
				appendHistory(code.String())
				runCode(code.String())
				code.Reset()
			}
			continue
		case "history":
			viewHistory()
			continue
		case "current":
			fmt.Println(code.String())
			continue
		default:
			code.WriteString(line + "\n")
		}
	}
}

func appendHistory(code string) {
	hs := history{code: code, timeStamp: time.Now()}
	codeHistory = append(codeHistory, hs)
}

func viewHistory() {
	for i, histr := range codeHistory {
		fmt.Printf("Index:%v\nCode:\n%vTimestamp:%v\n\n", i, histr.code, histr.timeStamp)
	}
}

func getPrompt() string {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("Error getting user: %v\n", err)
	}
	username := currentUser.Username

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error getting hostname: %v\n", err)
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current dir: %v\n", err)
	}

	// Replace home directory with ~
	homeDir := currentUser.HomeDir
	if strings.HasPrefix(dir, homeDir) {
		dir = "~" + strings.TrimPrefix(dir, homeDir)
	}

	green := "\033[32m"
	reset := "\033[0m"

	return fmt.Sprintf("%sMSPL%s %s@%s:%s$ ", green, reset, username, hostname, dir)
}

func runCode(code string) {
	// Wrap it
	full := fmt.Sprintf("package main\nimport \"fmt\"\nfunc main() {\n%s\n}", code)

	tmpFile, _ := os.CreateTemp("", "*.go")
	defer os.Remove(tmpFile.Name())

	tmpFile.WriteString(full)
	tmpFile.Close()

	cmd := exec.Command("go", "run", tmpFile.Name())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
