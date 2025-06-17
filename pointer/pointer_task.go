// Task: Write function f that increments the value of the pointer

package main

import (
	"bufio"
	"fmt"
	"os"
)

func f(count *int) {
	*count++ // Increment the value of the pointer
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Interaction counter")

	count := 0
	for {
		fmt.Print("-> ")

		_, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		f(&count)

		fmt.Printf("User input %d lines\n", count)
	}
}
