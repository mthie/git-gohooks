package general

import (
	"bytes"
	"io/ioutil"
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
	var outErr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &outErr

	if err := cmd.Start(); err != nil {
		log.Fatalf("cmd.Start: %v", err)
	}

	if err := cmd.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, okStatus := exiterr.Sys().(syscall.WaitStatus); okStatus {
				errCode = status.ExitStatus()
				result = outErr.String()
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
	for _, filename := range resultLines {
		if strings.HasSuffix(filename, ".go") {
			result = append(result, filepath.Join(absolutePath, filename))
		}
	}
	return result
}

// GetFilesList returns a list of all files in the current directory
func GetFilesList() (result []string) {
	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
		result = append(result, f.Name())
	}
	return result
}

// GetGitRoot returns the path with the .git directory
func GetGitRoot() string {
	gitDir, _ := RunCommand("git", "rev-parse", "--git-dir")
	return filepath.Dir(gitDir)
}
