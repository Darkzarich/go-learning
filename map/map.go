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

	// Exercise 1: Print all products
	exercise()

	// Exercise 2: Print all pairs of numbers that sum to k
	exercise2(13)

	// Exercise 3: Return a slice of two integers that sum to k
	result := exercise3(4)

	fmt.Println(result)

	// Exercise 4: Remove duplicates from a slice of strings
	words := []string{
		"cat",
		"dog",
		"bird",
		"bird",
		"dog",
		"parrot",
		"parrot",
		"cat",
		"hamster",
	}

	fmt.Println(RemoveDuplicates(words))
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

// Exercise 2: Print all pairs of numbers that sum to k
func exercise2(k int) {
	visited := make(map[int]bool)
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	for i1, v1 := range numbers {
		for i2, v2 := range numbers {
			if v1+v2 == k && !visited[i2] {
				fmt.Printf("%d (index: %d) + %d (index: %d) = %d\n", v1, i1, v2, i2, v1+v2)
			}
		}

		visited[i1] = true
	}
}

// Exercise 3: Return a slice of two integers that sum to k
// if there is no such pair of integers, return empty slice
func exercise3(k int) []int {
	m := make(map[int]int)
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	for i, v := range numbers {
		if _, ok := m[k-v]; ok {
			return []int{v, k - v}
		}
		m[v] = i
	}

	return nil
}

// Exercise 4: Remove duplicates from a slice of strings
func RemoveDuplicates(words []string) []string {
	result := make([]string, 0)
	seen := make(map[string]bool)

	for _, word := range words {
		if !seen[word] {
			seen[word] = true
			result = append(result, word)
		}
	}

	return result
}
