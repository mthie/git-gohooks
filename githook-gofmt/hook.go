package main

import (
	"fmt"
	"log"
	"os"

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
	fmt.Fprint(os.Stderr, "Go files must be formatted with gofmt. Please run:\n\n")
	fmt.Fprint(os.Stderr, "  gofmt -w")
	for _, file := range files {
		fmt.Fprint(os.Stderr, " \\\n")
		fmt.Fprintf(os.Stderr, "    %s", file)
	}
	fmt.Fprint(os.Stderr, "\n")
	os.Exit(1)
	log.Printf("Result: %+v", result)
}
