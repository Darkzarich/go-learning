package main

import (
	"fmt"
)

func main() {
	fmt.Println(Add(1, 2, 3, 4, 5))

	fmt.Println(Fact(5))
}

func Add(x ...int) int {
	sum := 0
	for _, i := range x {
		sum += i
	}
	return sum
}

func Fact(n int) int {
    if n == 0 {
        return 1
    } else {
        return n * Fact(n-1)
    }
} 