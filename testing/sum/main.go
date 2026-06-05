package main

import (
	"fmt"
	"mytest/sum"
)

func main() {
	result, err := sum.Sum(1, 2)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Sum:", result)
}
