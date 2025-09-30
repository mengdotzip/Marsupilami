package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Marsupilami")
	reader := bufio.NewReader(os.Stdin)
	var code strings.Builder

	for {
		fmt.Print("mspl> ")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "exit" {
			break
		}

		if line == "return" {
			if code.Len() > 0 {
				runCode(code.String())
				code.Reset()
			}
			continue
		}

		code.WriteString(line + "\n")
	}
}

func runCode(code string) {
	// Wrap ir
	full := fmt.Sprintf("package main\nimport \"fmt\"\nfunc main() {\n%s\n}", code)

	// Run using go run with temp file
	tmpFile, _ := os.CreateTemp("", "*.go")
	defer os.Remove(tmpFile.Name())

	tmpFile.WriteString(full)
	tmpFile.Close()

	cmd := exec.Command("go", "run", tmpFile.Name())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
