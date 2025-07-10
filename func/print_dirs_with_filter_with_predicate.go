package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// TASK:
// the same as func/print_dirs_with_filter.go program but printFileTreeWithFilter filter parameter should be a predicate function

func main() {
	printFileTreeWithFilterWithPredicate(".", func(path string) bool {
		return strings.Contains(path, "func")
	})
}

func printFileTreeWithFilterWithPredicate(path string, filter func(string) bool) {
	var walk func(string, int)

	walk = func(path string, level int) {
		// Getting a list of files
		files, err := os.ReadDir(path)

		if err != nil {
			fmt.Println("unable to get list of files", err)
			return
		}

		for _, f := range files {
			// Full path to the file
			fullPath := filepath.Join(path, f.Name())

			if f.Name() == ".git" || !filter(fullPath) {
				continue
			}

			levelSeparator := strings.Repeat("│  ", level)

			// Recursively call the function for a directory
			if f.IsDir() {
				fmt.Printf("%s├─ 📁 %s\n", levelSeparator, f.Name())
				walk(fullPath, level+1)
			} else {
				fmt.Printf("%s├─ 📄 %s\n", levelSeparator, f.Name())
			}
		}
	}

	walk(path, 0)
}
