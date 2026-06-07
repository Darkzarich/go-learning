package main

import (
	"fmt"
	"mymath/estimate"
	"mymath/sum"
)

func main() {
	result, err := sum.Sum(1, 2)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Sum:", result)

	textResult := estimate.EstimateValue(10)

	fmt.Println("Estimate:", textResult)
}
