package main

import "fmt"

// closure is a function that can access variables from the outer scope
func Generate(seed int) func() {
	return func() {
		fmt.Println(seed) // getting variable from the outer scope
		seed += 2         // modifying the variable from the outer scope
	}

}

func main() {
	iterator := Generate(0)
	iterator()
	iterator()
	iterator()
	iterator()
	iterator()

	fmt.Println("---")

	fib := FibonacciGenerator()

	fmt.Println(fib()) // x1 = 1, x2 = 1
	fmt.Println(fib()) // x1 = 1, x2 = 2
	fmt.Println(fib()) // x1 = 2, x2 = 3
	fmt.Println(fib()) // x1 = 3, x2 = 5
	fmt.Println(fib()) // x1 = 5, x2 = 8
	fmt.Println(fib()) // x1 = 8, x2 = 13
}

func FibonacciGenerator() func() int {
	x1, x2 := 0, 1
	// The returned function has access to the variables x1 and x2
	// and can modify them with each call
	return func() int {
		x1, x2 = x2, x1+x2
		return x1
	}
}
