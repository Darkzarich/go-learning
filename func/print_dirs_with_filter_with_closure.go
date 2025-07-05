package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// TASK:
// the same as func/print_dirs_with_filter.go program but with closures to avoid redundant copying of parameters
// also added passing filter as an argument
// usage: print_dirs_with_filter.go <filter>

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 0 {
		fmt.Println("Usage: print_dirs_with_filter.go <filter>")
		os.Exit(1)
	}

	filter := argsWithoutProg[0]

	printFileTreeWithFilter(".", filter)
}

func printFileTreeWithFilter(path string, filter string) {
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

			if f.Name() == ".git" || !strings.Contains(fullPath, filter) {
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
