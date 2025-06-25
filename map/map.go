package main

import "fmt"

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
		"ONE":   1,
		"TWO":   2,
		"THREE": 3,
	}

	fmt.Println(m3)
}
