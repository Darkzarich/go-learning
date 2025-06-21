package main

import "fmt"

func main() {
	var lastWeekTemp [7]int

	lastWeekTemp[0] = -1

	fmt.Println(lastWeekTemp)    // [-1, 0, 0, 0, 0, 0, 0]
	fmt.Println(lastWeekTemp[0]) // -1

	thisWeekTemp := [7]int{-3, 5, 7} // [-3 5 7 0 0 0 0]

	fmt.Println(thisWeekTemp) // [-3 5 7 0 0 0 0]

	rgbColor := [...]uint8{255, 255, 128}     // [255 255 128] len = 3
	rgbaColor := [...]uint8{255, 255, 128, 1} // [255 255 128 1] len = 4

	fmt.Println(rgbColor)  // [255 255 128]
	fmt.Println(rgbaColor) // [255 255 128 1]

	thisWeekTemp2 := [7]int{6: 11, 2: 3} // [0 0 3 0 0 0 11]

	fmt.Println(thisWeekTemp)       // [0 0 3 0 0 0 11]
	fmt.Println(len(thisWeekTemp2)) // 7

	var thisMonthTemp [4][7]int // multidimensional array

	fmt.Println(thisMonthTemp) // [[0 0 0 0 0 0 0] [0 0 0 0 0 0 0] [0 0 0 0 0 0 0] [0 0 0 0 0 0 0]]

	var weekTemp = [7]int{5, 4, 6, 8, 11, 9, 5}
	fmt.Printf("%f\n", averageTempWithFor(weekTemp))    // 6.857...
	fmt.Printf("%f\n", averageTempWithRange(&weekTemp)) // 6.857...
}

func averageTempWithFor(temps [7]int) float64 {
	sumTemp := 0

	for i := 0; i < len(temps); i++ {
		sumTemp += temps[i]
	}

	return float64(sumTemp) / float64(len(temps))
}

func averageTempWithRange(tempsPointer *[7]int) float64 {
	sumTemp := 0

	// range cycles through the array
	// better use pointers to arrays so that it won't copy the array
	for i := range tempsPointer {
		sumTemp += tempsPointer[i]
	}

	return float64(sumTemp) / float64(len(tempsPointer))
}
