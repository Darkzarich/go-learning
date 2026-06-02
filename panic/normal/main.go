package main

import (
	"fmt"
	"os"
)

func myFunc() {
	if _, err := os.ReadFile(`test.txt`); err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Start")
	myFunc()
	// Never reached
	fmt.Println("Finish")
}
