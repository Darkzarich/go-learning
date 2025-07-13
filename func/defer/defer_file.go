package main

import (
	"log"
	"os"
)

func main() {
	// Open file
	file, err := os.OpenFile("file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	// Close file after return
	defer file.Close()

	_, err = file.WriteString("hello world")

	if err != nil {
		log.Fatal(err)
	}
}
