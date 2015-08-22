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

	_, status := general.RunCommand("go", "build", ".", "./...")
	if status != 0 {
		fmt.Fprint(os.Stderr, "Build failed, please commit only stuff that builds.\n")
		os.Exit(1)
		return
	}
	general.RunCommand("go", "clean", ".", "./...")
}
