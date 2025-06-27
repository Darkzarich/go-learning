package main

import (
	"fmt"
)

type MyMap map[string]int

func main() {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3

	fmt.Println(m)

	M2 := make(MyMap)
	M2["one1"] = 1
	M2["two2"] = 2
	M2["three3"] = 3

	fmt.Println(M2)

	m3 := MyMap{
		"TWO":   2,
		"THREE": 3,
	}

	fmt.Println(m3)

	// taking values from map

	myMap := make(map[string]int)

	myMap["first"] = 5

	v1, ok1 := myMap["first"] // v1 = 5, ok1 = true

	v2, ok2 := myMap["second"] // v2 = 0, ok2 = false

	println(v1, ok1)
	println(v2, ok2)

	// reference
	referenceMap := myMap

	referenceMap["first"] = 10

	fmt.Println(referenceMap)

	// functions

	referenceMap["second"] = 20

	fmt.Println(len(referenceMap)) // length
	fmt.Println(referenceMap)

	delete(referenceMap, "second") // delete key-value pair

	fmt.Println(referenceMap)

	// range

	for key, value := range referenceMap {
		fmt.Printf("Key \"%v\" has value %v\n", key, value)
	}

	exercise()
}

// the small exercise for map
func exercise() {
	products := map[string]int{
		"bread":     50,
		"milk":      100,
		"butter":    200,
		"sausage":   500,
		"salt":      20,
		"cucumbers": 200,
		"cheese":    600,
		"ham":       700,
		"meat":      900,
		"tomatoes":  250,
		"fish":      300,
		"drink":     1500,
	}

	// task 0: print all products

	fmt.Println("All products:")

	for product, price := range products {
		fmt.Println("- ", product, ":", price)
	}

	// task 1: print products more expensive than 500

	fmt.Println("Products more expensive than 500:")

	for product, price := range products {
		if price > 500 {
			fmt.Println("- ", product, ":", price)
		}
	}

	// task 2: count total price of an order provided by slice

	order := []string{"bread", "meat", "cheese", "cucumbers"}

	totalPrice := 0

	for _, product := range order {
		totalPrice += products[product]
	}

	fmt.Println("Total price of order:", totalPrice)
}
