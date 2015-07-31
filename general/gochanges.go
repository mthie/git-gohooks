package general

import (
	"bytes"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

// RunCommand is a simple wrapper to run commands
// result is the output of stdout and errCode is the exit code
func RunCommand(command string, values ...string) (result string, errCode int) {
	cmd := exec.Command(command, values...)
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Start(); err != nil {
		log.Fatalf("cmd.Start: %v", err)
	}

	if err := cmd.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				errCode = status.ExitStatus()
				return result, errCode
			}
		} else {
			log.Fatalf("cmd.Wait: %v", err)
			return result, errCode
		}
	}

	result = out.String()
	return result, errCode
}

// GetChangedGoFiles returns a list of changed .go files
// anything else is filtered out
func GetChangedGoFiles() (result []string) {
	gitDiff, _ := RunCommand("git", "diff", "--name-only", "--cached", "--diff-filter=ACM")
	absolutePath := GetGitRoot()

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

// GetGitRoot returns the path with the .git directory
func GetGitRoot() string {
	gitDir, _ := RunCommand("git", "rev-parse", "--git-dir")
	return filepath.Dir(gitDir)
}
