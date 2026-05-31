package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	_, err := os.ReadFile("non-existing")

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("unable read file")
			return
		}
		fmt.Println("unknown error")
		return
	}
}
