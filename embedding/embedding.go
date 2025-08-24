package main

import (
	"fmt"
)

type Person struct {
	Name string
	Year int
}

func NewPerson(name string, year int) Person {
	return Person{
		Name: name,
		Year: year,
	}
}

// String returns a string representation of Person
func (p Person) String() string {
	return fmt.Sprintf("Name: %s, Year of birth: %d", p.Name, p.Year)
}

// Will call the String() method of Person
func (p Person) Print() {
	fmt.Println(p)
}

// Embedding Person into Student
type Student struct {
	Person
	Group string
}

func NewStudent(name string, year int, group string) Student {
	return Student{
		Person: NewPerson(name, year), // Creating a new Person explicitly
		Group:  group,
	}
}

// String returns a string representation of Student
// Also using the String() method of Person
func (s Student) String() string {
	return fmt.Sprintf("%s, Group: %s", s.Person, s.Group)
}

func main() {
	s := NewStudent("John Doe", 1980, "701")
	// Creating Student but calling the Print() method of embedded Person
	s.Print()
	// Will call the String() method of Person
	fmt.Println(s)
	fmt.Println(s.Name, s.Year, s.Group)
}
