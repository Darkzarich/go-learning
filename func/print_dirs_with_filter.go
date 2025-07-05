package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// TASK:
// the same as func/print_dirs.go program but with filter parameter
// to filter by path and file name
// Usage: print_dirs_with_filter.go <filter> # for example .png

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 0 {
		fmt.Println("Usage: print_dirs_with_filter.go <filter>")
		os.Exit(1)
	}

	filter := argsWithoutProg[0]

	PrintFileTreeWithFilter(".", filter, 0)
}

func PrintFileTreeWithFilter(path string, filter string, level int) {
	// Getting a list of files
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("unable to get list of files", err)
		return
	}

	for _, f := range files {
		// Full path to the file
		fullPath := filepath.Join(path, f.Name())

		if f.Name() == ".git" || !strings.Contains(fullPath, filter) {
			continue
		}

		levelSeparator := strings.Repeat("│  ", level)

		// Recursively call the function for a directory
		if f.IsDir() {
			fmt.Printf("%s├─ 📁 %s\n", levelSeparator, f.Name())

			PrintFileTreeWithFilter(fullPath, filter, level+1)
		} else {
			fmt.Printf("%s├─ 📄 %s\n", levelSeparator, f.Name())
		}
	}
}
