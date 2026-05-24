package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func ChangeFieldByName(v interface{}, fname string, newval int) {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return
	}

	field := val.FieldByName(fname)
	if field.IsValid() {
		if field.CanSet() {
			switch field.Kind() {
			case reflect.Int:
				field.SetInt(int64(newval))
			case reflect.String:
				field.SetString(strconv.Itoa(newval))
			}
		}
	}
}

type MyStruct struct {
	Field1 int
	Field2 string
}

func main() {
	s := MyStruct{Field1: 1, Field2: "some"}

	// CanSet will return false because struct is passed by value
	ChangeFieldByName(s, "Field1", 2)

	fmt.Println(s)

	ChangeFieldByName(&s, "Field1", 2)

	fmt.Println(s)
}
