package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func GenerateIntoFile(dstFileName string) error {
	// Random number generator implements io.Reader
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	dstFile, err := os.Create(dstFileName)
	if err != nil {
		return err
	}

	// Gotta make sure to use CopyN because normal Copy will never return EOF
	n, err := io.CopyN(dstFile, generator, 10)
	if err != nil {
		return err
	}

	fmt.Printf("Copied %d bytes to %s", n, dstFileName)
	return nil
}

func main() {
	err := GenerateIntoFile("test.txt")
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
