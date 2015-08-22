package main

import (
	"fmt"
	"os"

	"github.com/mthie/git-gohooks/general"
)

func main() {
	files := general.GetChangedGoFiles()
	if len(files) == 0 {
		os.Exit(0)
		return
	}

	os.Chdir(general.GetGitRoot())

	result, status := general.RunCommand("go", "test", "-test.short", "./...")
	if status != 0 {
		fmt.Fprintf(os.Stderr, "Test failed, please commit only stuff that works.\n%s", result)
		os.Exit(1)
		return
	}
}
