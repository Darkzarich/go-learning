package main

import (
	"fmt"
	"reflect"
)

func ExtendedPrint(v interface{}) {
	val := reflect.ValueOf(v)
	//  Checking if the value is a pointer
	switch val.Kind() {
	case reflect.Ptr:
		if val.Elem().Kind() != reflect.Struct {
			fmt.Printf("Pointer to %v : %v", val.Elem().Type(), val.Elem())
			return
		}
		// If it's a pointer to a struct, then we will be working with that struct
		val = val.Elem()

	// working with a struct now
	case reflect.Struct:
	default:
		fmt.Printf("%v : %v", val.Type(), val)
		return
	}

	fmt.Printf("Struct of type %v and number of fields %d:\n", val.Type(), val.NumField())
	for fieldIndex := 0; fieldIndex < val.NumField(); fieldIndex++ {
		// Field returns a Value
		field := val.Field(fieldIndex)
		// getting field name out of its type, not the value
		fmt.Printf("\tField %v: %v - val :%v\n", val.Type().Field(fieldIndex).Name, field.Type(), field)
	}
}

type MyStruct struct {
	A int
	B string
	C bool
}

func main() {
	s := MyStruct{
		A: 3,
		B: "some",
		C: false,
	}
	s1 := &MyStruct{
		A: 7,
		B: "text",
		C: true,
	}

	ExtendedPrint(s)
	ExtendedPrint(s1)
	ExtendedPrint(struct {
		E int
		C string
	}{2, "other text"})
	ExtendedPrint("some string")
}
