package main

import "fmt"

const (
	Black = iota
	Gray
	White
)

const (
	Yellow = iota
	Red
	Green = iota
	Blue
)

func main() {
	fmt.Println(Black, Gray, White)       // 0 1 2
	fmt.Println(Yellow, Red, Green, Blue) // 0 1 2 3
	fmt.Println(15 % 7)
}
