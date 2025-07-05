package main

import "fmt"

// A more practical example of using closures

func main() {
	f := CountCalls(SpecialPrint)
	f("Hello")
	f("World")

	// Wrapping another time

	f = CountCalls(f)

	f("!")
	f("!")
}

func SpecialPrint(s string) {
	fmt.Println("SpecialPrint:", s)
}

func CountCalls(f func(string)) func(string) {
	count := 0

	return func(s string) {
		count++

		f(s)

		fmt.Println("Been called", count, "times")
	}
}
