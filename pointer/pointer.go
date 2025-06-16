package main

import "fmt"

func main() {
	var (
		i int  = 5
		p *int = &i // Set the address of i
	)

	fmt.Printf("i: %d\n", i)
	fmt.Printf("p: %p\n", p)
	fmt.Printf("*p: %d\n", *p)

	type Person struct {
		Name string
		Age  int
	}

	person := Person{"John", 30}            // Struct value in a variable
	fmt.Printf("person: %s\n", person.Name) // Pointer to a struct

	p2 := &Person{"Max", 28}             // Pointer to a struct
	fmt.Printf("p2: %p\n", p2)           // Pointer to a struct
	fmt.Printf("*p2: %d\n", *&p2.Age)    // Explicit dereference
	fmt.Printf("p2.Name: %s\n", p2.Name) // implicit dereference
	fmt.Printf("p2.Name: %d\n", p2.Age)  // implicit dereference

	p3 := new(Person) // Same as &Person{}
	fmt.Printf("p3: %p\n", p3)

	p4 := &p3 // Pointer to a pointer

	fmt.Printf("p4: %p\n", p4)
	fmt.Printf("*p4: %p\n", *p4) // Value of the pointer to a pointer (address)

	p5 := p4
	fmt.Printf("p5 == p4: %t\n", p5 == p4) // Comparison of two pointers

	var panicPointer *int                       // nil
	fmt.Println("panicPointer:", *panicPointer) // panic: runtime error: invalid memory address or nil pointer dereference
}
