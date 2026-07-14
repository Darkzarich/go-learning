package main

import (
	"fmt"
	"math"
)

// Newton's method for square root

func Sqrt(x float64) float64 {
	z := 1.0

	for i := 0; i < 10; i++ {
		// (z*z - x) - is how far away z² is from where it needs to be (the error)
		// 2z is derivative, the speed of changing
		z -= (z*z - x) / (2 * z)
	}

	return z
}

// Version 2 of SQRT that returns the number of iterations
func SqrtV2(x float64) (float64, int) {
	z := x/2
	iter := 0
	for math.Abs(x-z*z) > 0.001 {
		z -= (z*z - x) / (2 * z)
		iter++
	}
	return z, iter
}

func main() {
	target := 10.0

	fmt.Printf("Sqrt(%f) = %f (took 10 iterations)\n", target, Sqrt(target))

	res, iter := SqrtV2(target)
	fmt.Printf("SqrtV2(%f) = %f (took %d iterations)\n", target, res, iter)
}
