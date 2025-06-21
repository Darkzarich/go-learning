package main

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
)

func main() {
	var slice []int

	fmt.Println(slice) // []

	slice = append(slice, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println(slice) // [1 2 3 4 5 6 7 8 9 10]

	mySlice1 := make([]int, 0)     // slice [], base array []
	mySlice2 := make([]int, 5)     // слайс [0 0 0 0 0], базовый массив [0 0 0 0 0]
	mySlice3 := make([]int, 5, 10) // слайс [0 0 0 0 0], базовый массив [0 0 0 0 0 0 0 0 0 0]

	mySlice4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(mySlice1) // []
	fmt.Println(mySlice2) // [0 0 0 0 0]
	fmt.Println(mySlice3) // [0 0 0 0 0]
	fmt.Println(mySlice4) // [1 2 3 4 5]

	mySlice4[2] = 100 // changing base array changes all slices based on it

	fmt.Println("Whole slice: ", mySlice4[:])
	fmt.Println("First three elements: ", mySlice4[:3])
	fmt.Println("Last three elements: ", mySlice4[len(mySlice4)-3:])

	var mySlice5 = [5]int{1, 2, 3, 4, 5}

	fmt.Println(mySlice5[0:2]) // [1 2 3 4 5]

	// appending to slice

	a := []int{1, 2, 3, 4}
	b := a[2:3] // b = [3]
	b = append(b, 7)

	fmt.Println(a, len(a), cap(a)) // [1 2 3 7] 4 4
	fmt.Println(b, len(b), cap(b)) // [3 7] 2 2

	b = append(b, 8, 9, 10)
	b[0] = 11

	fmt.Println(a, len(a), cap(a)) // [1 2 3 7] 4 4
	fmt.Println(b, len(b), cap(b)) // [11 7 8 9 10] 5 6

	// combine slices

	a = []int{1, 2, 3, 4}
	b = []int{5, 6, 7, 8}

	c := append(a, b...) // unpacking b, c = [1 2 3 4 5 6 7 8]

	fmt.Println(a, len(a), cap(a)) // [1 2 3 4] 4 4
	fmt.Println(b, len(b), cap(b)) // [5 6 7 8] 4 4
	fmt.Println(c, len(c), cap(c)) // [1 2 3 4 5 6 7 8] 8 8

	// sort

	s := []int{5, 4, 1, 3, 2}
	sort.Ints(s)   // doesn't change length or capacity
	fmt.Println(s) // [1 2 3 4 5] - returned the same slice but sorted

	bSlice := []byte(" \t\n a lone gopher \n\t\r\n")
	fmt.Printf("%s", bytes.TrimSpace(bSlice)) // a lone gopher - changed slice size so returned a new slice
	fmt.Printf("%s", bSlice)                  // \t\n a lone gopher \n\t\r\n

	// copy slice
	originalSlice := []int{1, 2, 3, 4, 5}
	copySlice := make([]int, len(originalSlice))
	copy(copySlice, originalSlice)

	fmt.Println(originalSlice, copySlice) // [1 2 3 4 5] [1 2 3 4 5]

	// range over slice

	for i, v := range []int{1, 2, 3, 4, 5} {
		fmt.Printf("%d: %d\n", i, v)
	}

	// delete the last slice element

	s3 := []int{1, 2, 3}

	if len(s3) != 0 { // protection from index out of range panic
		s3 = s3[:len(s3)-1]
	}
	fmt.Println(s3) // [1 2]

	// delete the first slice element

	s4 := []int{1, 2, 3}
	if len(s4) != 0 { // protection from index out of range panic
		s4 = s4[1:]
	}
	fmt.Println(s4) // [2 3]

	// delete i element from the slice

	s5 := []int{1, 2, 3, 4, 5}
	i := 2

	if len(s5) != 0 && i < len(s5) { // protection from index out of range panic
		s5 = append(s5[:i], s5[i+1:]...)
	}
	fmt.Println(s5) // [1 2 4 5]

	// compare slices

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{1, 2, 3, 4, 5}
	slice3 := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(reflect.DeepEqual(slice1, slice2)) // true
	fmt.Println(reflect.DeepEqual(slice1, slice3)) // false

	// task - fill the slice with numbers from 1 to 100,
	// slice first 10 elements and then the last 10
	// reverse the slice

	length := 100
	sliceToFill := make([]int, 0, length)

	for i := 1; i <= length; i++ {
		sliceToFill = append(sliceToFill, i)
	}

	sliceToFill = append(sliceToFill[:10], sliceToFill[len(sliceToFill)-10:]...)

	length = len(sliceToFill)

	for i := range slice[:length/2] {
		sliceToFill[i], sliceToFill[length-i-1] = sliceToFill[length-i-1], sliceToFill[i]
	}

	fmt.Println(sliceToFill)
}
