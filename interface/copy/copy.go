package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(srcFileName, dstFileName string) error {
	// Returns pointer to file *os.File that has method Read() hence implementing io.Reader
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		return err
	}

	// Returns pointer to file *os.File that has method Write() hence implementing io.Writer
	dstFile, err := os.Create(dstFileName)
	if err != nil {
		return err
	}

	// io.Copy expects both io.Reader and io.Writer
	n, err := io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	fmt.Printf("Copied %d bytes from %s to %s", n, srcFileName, dstFileName)
	return nil
}

func main() {
	err := CopyFile("test.txt", "copy.txt")
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
