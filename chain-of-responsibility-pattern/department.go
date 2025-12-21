package main

// Base interface for all departments
type Department interface {
	execute(*Patient)
	setNext(Department)
}
