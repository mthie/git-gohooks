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
	hookBase := strings.TrimSuffix(filepath.Base(os.Args[0]), filepath.Ext(os.Args[0]))
	hookPrefix := fmt.Sprintf("%s_", hookBase)
	os.Chdir(gitroot)

	filepath.Walk(gitroot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return filepath.SkipDir
		}
		file := info.Name()

		if strings.HasPrefix(file, hookPrefix) {
			result, errCode := general.RunCommand(filepath.Join(gitroot, "/.git/hooks", file))
			if errCode != 0 {
				fmt.Fprintf(os.Stderr, "Error: %s", result)
				os.Exit(errCode)
				return nil
			}
		}
		return nil
	})
}
