package main

import (
	"fmt"
	ms "sum_slice_module_task/mathslice"
)

func main() {
	slice := []ms.Element{1, 2, 3, 4, 5}

	fmt.Println("Slice: ", slice)

	fmt.Println("Sum:", ms.SumSlice(slice))

	fmt.Println("Avg: ", ms.AverageSlice(slice))

	ms.MapSlice(slice, func(v ms.Element) ms.Element {
		return v * 2
	})

	fmt.Println("Mapped as x*2: ", slice)
}
