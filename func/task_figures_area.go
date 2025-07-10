package main

import (
	"fmt"
	"math"
)

/*
  Task: Write a function that calculates the area of a figure.
	The function should take a figure as a parameter and return a function that calculates the area of the figure
	for simplicity, the function should take a float64 as a parameter and return a float64
*/

type figures int

const (
	square figures = iota
	circle
	triangle // same-sided
)

func main() {
	ar, ok := area(square)

	if !ok {
		fmt.Println("Invalid figure")

		return
	}

	square5Area := ar(5)
	square10Area := ar(10)

	fmt.Println("Area of square with side 5 is", square5Area)
	fmt.Println("Area of square with side 10 is", square10Area)

	ar, ok = area(circle)

	if !ok {
		fmt.Println("Invalid figure")

		return
	}

	circle5Area := ar(5)
	circle10Area := ar(10)

	fmt.Println("Area of circle with radius 5 is", circle5Area)
	fmt.Println("Area of circle with radius 10 is", circle10Area)

	ar, ok = area(triangle)

	if !ok {
		fmt.Println("Invalid figure")

		return
	}

	triangle5Area := ar(5)
	triangle10Area := ar(10)

	fmt.Println("Area of triangle with side 5 is", triangle5Area)
	fmt.Println("Area of triangle with side 10 is", triangle10Area)
}

func area(figure figures) (func(float64) float64, bool) {
	switch figure {
	case square:
		return func(side float64) float64 {
			return side * side
		}, true
	case circle:
		return func(radius float64) float64 {
			return math.Pi * radius * radius
		}, true
	case triangle:
		return func(side float64) float64 {
			return math.Pow(side, 2) * math.Sqrt(3) / 4
		}, true
	default:
		return nil, false
	}
}
