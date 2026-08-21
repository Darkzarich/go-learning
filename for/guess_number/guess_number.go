package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Print("Guess a number : ")
	var guess int
	var attempt int = 0
	var target int = rand.IntN(100)

	for {
		_, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Println("Could not read input")
			continue
		}

		attempt++

		if guess < target {
			fmt.Println("Guessed number is too low!")
		} else if guess > target {
			fmt.Println("Guessed number is too high!")
		} else {
			fmt.Printf("You guessed it in %d attempts!\n", attempt)
			break
		}
	}
}
