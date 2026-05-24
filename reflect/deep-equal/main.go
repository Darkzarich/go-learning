package main

import (
	"fmt"
	"reflect"
)

type MyType struct {
	intField int
	strField string
}

func (mt MyType) IsEqual(mt2 MyType) bool {
	return mt == mt2
}

type MyTypeWithRefs struct {
	intField     int
	strField     string
	pointerField *float64
	sliceField   []int
}

func (mt MyTypeWithRefs) IsEqual(mt2 MyTypeWithRefs) bool {
	return reflect.DeepEqual(mt, mt2)
}

func main() {
	a := MyType{intField: 1, strField: "str"}
	b := MyType{intField: 1, strField: "str"}

	fmt.Printf("Is a equal to b: %v\n", a.IsEqual(b))

	// References are compared by pointer, not by value, so they are not equal when comparing structs
	floatValue1, floatValue2 := 10.0, 10.0
	a2 := MyTypeWithRefs{intField: 1, strField: "str", pointerField: &floatValue1, sliceField: []int{1}}
	b2 := MyTypeWithRefs{intField: 1, strField: "str", pointerField: &floatValue2, sliceField: []int{1}}

	fmt.Printf("Is a2 equal to b2: %v\n", a2.IsEqual(b2))
}
