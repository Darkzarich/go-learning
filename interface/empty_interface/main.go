package main

import (
	"fmt"
	"strconv"
)

func PrintPassedInt(v interface{}) {
	i, ok := v.(int) // if v is not a number then ok will be false
	// No panic here
	if ok {
		fmt.Println("PrintPassedInt: Passed number", i)
	} else {
		fmt.Println("PrintPassedInt: Not passed number, v is", v)
	}
}

func MustPrintPassedInt(v interface{}) {
	// if v is not number then it will panic
	i := v.(int)
	fmt.Println("MustPrintPassedInt: Passed number", i)
}

func PrintWhatTypeIs(v interface{}) {
	switch val := v.(type) {
	case int:
		fmt.Println("PrintWhatTypeIs: Passed int", strconv.Itoa(val))
	case string:
		fmt.Println("PrintWhatTypeIs: Passed string", val)
	case float64:
		fmt.Println("PrintWhatTypeIs: Passed float64", strconv.FormatFloat(val, 'f', -1, 64))
	case bool:
		fmt.Println("PrintWhatTypeIs: Passed bool", strconv.FormatBool(val))
	case fmt.Stringer:
		fmt.Println("PrintWhatTypeIs: Passed value that implements Stringer interface, value is", val.String())
	default:
		fmt.Println("PrintWhatTypeIs: no known type detected, v is", v)
	}
}

type Person struct {
	Name string
}

func (p Person) String() string {
	return p.Name
}

func main() {
	PrintPassedInt(1)
	PrintPassedInt("1")
	PrintPassedInt(1.0)
	PrintPassedInt(true)

	PrintWhatTypeIs(1)
	PrintWhatTypeIs("1")
	PrintWhatTypeIs(1.0)
	PrintWhatTypeIs(true)
	PrintWhatTypeIs(Person{"John"})

	MustPrintPassedInt(1)
	// Everything else will panic
	MustPrintPassedInt("1")
	MustPrintPassedInt(1.0)
	MustPrintPassedInt(true)
}
