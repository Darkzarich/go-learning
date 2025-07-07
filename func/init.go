package main

import "fmt"

var name, surname string

// init() is called before main() and is used to initialize global variables
func init() {
	name = "John"
}
func init() {
	if surname == "" {
		surname = "Doe"
	}
}
func main() {
	fmt.Println("Hello " + name + " " + surname)
}
