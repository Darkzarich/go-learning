package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your year: ")

	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input: ", err)
		return
	}

	year, err := strconv.Atoi(strings.TrimSpace(text)) // ASCII to int

	if err != nil {
		fmt.Println("Error parsing input: ", err)
		return
	}

	switch {
	case year >= 1946 && year <= 1964:
		fmt.Println("Hello, boomer!")
	case year >= 1965 && year <= 1980:
		fmt.Println("Hello, gen-x!")
	case year >= 1981 && year <= 1996:
		fmt.Println("Hello, millennial!")
	case year >= 1997 && year <= 2012:
		fmt.Println("Hello, zoomer!")
	case year >= 2013:
		fmt.Println("Hello, gen alpha!")
	}
}
