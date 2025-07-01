package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Person struct {
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	dateOfBirth time.Time `json:"-"` // This field is ignored by encoding/json
}

func NewPerson(name, lastName string, dobYear, dobMonth, dobDay int) (Person, error) {
	if len(name) == 0 {
		return Person{}, errors.New("name cannot be empty")
	}

	return Person{
		FirstName:   name,
		LastName:    lastName,
		dateOfBirth: time.Date(dobYear, time.Month(dobMonth), dobDay, 0, 0, 0, 0, time.UTC),
	}, nil
}

func main() {
	// Creating a struct
	person := Person{}
	person.FirstName = "Max"
	person.LastName = "Jones"

	person2 := Person{
		FirstName: "Max",
		LastName:  "Jones",
	}

	person3 := Person{"Max", "Jones", time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)}

	// Using constructor function
	person4, err := NewPerson("Max", "Jones", 1990, 1, 1)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(person)
	fmt.Println(person2)
	fmt.Println(person3)
	fmt.Println(person4)

	jsonPerson, err := json.Marshal(person3)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonPerson))

	// Anonymous struct
	var person5 = struct {
		FirstName string
		LastName  string
	}{
		FirstName: "Max",
		LastName:  "Jones",
	}

	fmt.Println(person5)
}
