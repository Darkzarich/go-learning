package main

import "fmt"

func myPanic() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(`A panic happened: `, p)
		}
	}()
	panic(`dangerous situation`)
}

func main() {
	fmt.Println("Start")
	myPanic()
	// since recover() is called, the code after panic() will be executed
	fmt.Println("Finish")
}
