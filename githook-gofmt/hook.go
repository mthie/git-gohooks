package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mthie/git-gohooks/general"
)

func main() {
	files := general.GetChangedGoFiles()
	if files == nil {
		os.Exit(0)
		return
	}

	args := []string{"-l"}
	args = append(args, files...)

	result, _ := general.RunCommand("gofmt", args...)
	if result == "" {
		os.Exit(0)
		return
	}

	resultFiles := strings.Split(result, "\n")

	fmt.Fprint(os.Stderr, "Go files must be formatted with gofmt. Please run:\n\n")
	fmt.Fprint(os.Stderr, "  gofmt -w")
	for _, file := range resultFiles {
		fmt.Fprint(os.Stderr, " \\\n")
		fmt.Fprintf(os.Stderr, "    %s", file)
	}
	fmt.Fprint(os.Stderr, "\n")
	os.Exit(1)
}
