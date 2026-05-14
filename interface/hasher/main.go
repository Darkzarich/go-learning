package main

import (
	"crypto/rand"
	"fmt"
	"hasher/hashbyte"
)

func main() {

	buf := make([]byte, 16)

	for i := 0; i < 5; i++ {
		n, _ := rand.Read(buf)
		fmt.Printf("Generate bytes: %v size(%d)\n", buf, n)
	}

	hasher := hashbyte.New(0)
	hasher.Write(buf)
	fmt.Printf("Hash: %v \n", hasher.Hash())
}
