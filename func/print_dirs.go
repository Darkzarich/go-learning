package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	PrintAllFiles(".", 0)
}

func PrintAllFiles(path string, level int) {
	// Getting a list of files
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("unable to get list of files", err)
		return
	}

	for _, f := range files {
		// Full path to the file
		fullPath := filepath.Join(path, f.Name())

		if f.Name() == ".git" {
			continue
		}

		levelSeparator := strings.Repeat("│  ", level)

		// Recursively call the function for a directory
		if f.IsDir() {
			fmt.Printf("%s├─ 📁 %s\n", levelSeparator, f.Name())
			PrintAllFiles(fullPath, level+1)
		} else {
			fmt.Printf("%s├─ 📄 %s\n", levelSeparator, f.Name())
		}
	}
}
