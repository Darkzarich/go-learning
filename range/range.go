package main

import "fmt"

func main() {
	// create array
	array := [3]int{1, 2, 3}
	// iterate over array
	for arrayIndex, arrayValue := range array {
		fmt.Printf("array[%d]: %d\n", arrayIndex, arrayValue)
	}
}
