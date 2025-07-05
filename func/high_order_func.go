package main

import "fmt"

func main() {
	fmt.Println(HighOrderFunc(func(x int) int { return x + 1 }))
}

func HighOrderFunc(f func(int) int) int {
	return f(1)
}
