package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	w := strings.Builder{}

	for i := 0; i < 50; i++ {
		fmt.Fprintf(&w, "%v", math.NaN())
	}

	w.Write([]byte("... BATMAN!"));

	// выводим собранную строку
	fmt.Printf("%s\n", &w)
}
