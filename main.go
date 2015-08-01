package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mthie/git-gohooks/general"
)

func main() {
	os.Chdir(general.GetGitRoot() + "/.git/hooks")
	currentFileSplit := strings.Split(os.Args[0], "/")
	currentFile := currentFileSplit[len(currentFileSplit)-1]
	files := general.GetFilesList()
	for _, file := range files {
		if strings.HasPrefix(file, fmt.Sprintf("%s_", currentFile)) {
			result, errCode := general.RunCommand("./" + file)
			if errCode != 0 {
				fmt.Fprintf(os.Stderr, "Error: %s", result)
				os.Exit(errCode)
				return
			}
		}
	}
}
