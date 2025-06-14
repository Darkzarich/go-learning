package main

import "fmt"

func main() {
	// создаём массив
	array := [3]int{1, 2, 3}
	// итерируемся
	for arrayIndex, arrayValue := range array {
		fmt.Printf("array[%d]: %d\n", arrayIndex, arrayValue)
	}
}
