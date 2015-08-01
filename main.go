package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mthie/git-gohooks/general"
)

func main() {
	gitroot, _ := filepath.Abs(filepath.Dir(general.GetGitRoot()))
	os.Chdir(gitroot + "/.git/hooks")
	currentFileSplit := strings.Split(os.Args[0], "/")
	currentFile := currentFileSplit[len(currentFileSplit)-1]
	files := general.GetFilesList()
	os.Chdir(gitroot)

	for _, file := range files {
		if strings.HasPrefix(file, fmt.Sprintf("%s_", currentFile)) {
			result, errCode := general.RunCommand(gitroot + "/.git/hooks/" + file)
			if errCode != 0 {
				fmt.Fprintf(os.Stderr, "Error: %s", result)
				os.Exit(errCode)
				return
			}
		}
	}
}
