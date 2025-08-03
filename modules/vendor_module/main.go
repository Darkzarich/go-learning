package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

type Person struct {
	Name string
	Age  int
}

type Booth struct {
	People   []Person
	Title    string
	Capacity int
}

func main() {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}

	if cmp.Equal(a, b) {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}

	person1 := Person{Name: "John", Age: 30}
	person2 := Person{Name: "John", Age: 30}

	if cmp.Equal(person1, person2) {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}

	// Recursive too

	booth1 := Booth{
		People: []Person{
			{Name: "John", Age: 30},
			{Name: "Jane", Age: 25},
		},
		Title:    "CEO",
		Capacity: 10,
	}

	booth2 := Booth{
		People: []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
		},
		Title:    "Tech lead",
		Capacity: 10,
	}

	if cmp.Equal(booth1, booth2) {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}
}
