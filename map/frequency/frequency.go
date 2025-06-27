package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// random Japanese sentence because I just like Japanese :P
const defaultString = "昨日、友達と映画館に行って、面白い映画を観ました"

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your sentence, the frequency of each symbol will be printed: ")

	sentence, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input: ", err)
		return
	}

	sentence = strings.TrimSpace(sentence)

	if len(sentence) == 0 {
		fmt.Println("Will use the default string: ", defaultString)

		sentence = defaultString
	}

	frequency := make(map[rune]int)

	// go, strings can be iterated over because they are fundamentally treated as read-only slices of bytes
	for _, v := range sentence {
		frequency[v]++
	}

	for k, v := range frequency {
		fmt.Printf("Symbol %c occurs %d times\n", k, v)
	}
}
