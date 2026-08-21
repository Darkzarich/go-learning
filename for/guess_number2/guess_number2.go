package main

import "fmt"

const (
	high = 100
	low  = 1
	// With binary search, attempts never exceed 7
	// if more, the user is a cheater)
	maxAttempts = 7
)

func main() {
	fmt.Printf("🤖 Think of a number from %d to %d: ", low, high)
	var target int

	_, err := fmt.Scan(&target)
	if err != nil {
		fmt.Println("Failed to read the number")
		return
	}

	if target < low || target > high {
		fmt.Println("You entered an invalid number. Error!")
		return
	}

	var curHigh int = high
	var curLow int = low
	var guess int
	var answer string
	var attempts int = 0

	for {
		if attempts > maxAttempts {
			fmt.Printf("🤖 Looks like you cheated me")
			break
		}

		guess = (curHigh + curLow) / 2

		fmt.Printf("🤖 Is your number possibly %d? (y)es/(l)ower/(h)igher\n", guess)

		_, err := fmt.Scan(&answer)
		if err != nil {
			fmt.Println("Failed to read the answer")
			continue
		}

		if answer != "y" && answer != "l" && answer != "h" {
			fmt.Print("🤖 Let's try again!\n")
			continue
		}

		if answer == "y" {
			fmt.Println("🤖 I guessed it!")
			break
		} else if answer == "l" {
			curHigh = guess - 1
		} else if answer == "h" {
			curLow = guess + 1
		}

		attempts++
	}
}
