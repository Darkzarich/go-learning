package main

import "fmt"

type Person struct {
	Name    string
	Surname string
	Age     int
}

type PersonOption func(*Person)

func NewPerson(opts ...PersonOption) *Person {
	// Init person with default values that are not default for their type
	i := &Person{
		Name:    "FirstName",
		Surname: "Surname",
		Age:     -1,
	}

	// Apply options
	for _, opt := range opts {
		opt(i)
	}

	return i
}

func WithName(name string) PersonOption {
	return func(p *Person) {
		p.Name = name
	}
}

func WithSurname(surname string) PersonOption {
	return func(p *Person) {
		p.Surname = surname
	}
}

func WithAge(age int) PersonOption {
	return func(p *Person) {
		p.Age = age
	}
}

func main() {
	john := NewPerson(
		WithName("John"),
		WithSurname("Doe"),
		WithAge(30),
	)

	max := NewPerson(
		WithName("Max"),
		WithSurname("Mustermann"),
		WithAge(42),
	)

	fmt.Println(john)
	fmt.Println(max)
}
