package general

import (
	"bytes"
"os/exec"
	"path/filepath"
	"strings"
)

// RunCommand is a simple wrapper to run commands
func RunCommand(command string, values ...string) string {
	cmd := exec.Command(command, values...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	return out.String()
}

// GetChangedGoFiles returns a list of changed .go files
// anything else is filtered out
func GetChangedGoFiles() (result []string) {
	gitDir := RunCommand("git", "rev-parse", "--git-dir")
	gitDiff := RunCommand("git", "diff", "--name-only", "--cached", "--diff-filter=ACM")
	absolutePath := filepath.Dir(gitDir)

	resultLines := strings.Split(gitDiff, "\n")
	if resultLines == nil {
		return
	}
	for _, filename := range resultLines {
		if filename != "" && strings.HasSuffix(filename, ".go") {
			result = append(result, absolutePath+"/"+filename)
		}
	}
	return result
}
