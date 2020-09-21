package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func convertPath(path string) string {
	formatted := strings.Trim(path, "\\\\wsl$\\Ubuntu")
	formatted = strings.ReplaceAll(formatted, "\\", "/")
	return formatted
}

func main() {
	path := os.Args[1]
	localPath := os.Getenv("LOCALAPPDATA")
	var finalArguments []string
	if strings.HasPrefix(path, "\\\\wsl$") {
		finalArguments = append(finalArguments, "--remote", "wsl+ubuntu", convertPath(path))
	} else {
		finalArguments = append(finalArguments, path)
	}
	cmd := exec.Command(fmt.Sprintf("%v\\Programs\\Microsoft VS Code\\code.exe", localPath), finalArguments...)
	cmd.Run()
}
