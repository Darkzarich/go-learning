package main

import (
	"fmt"
	"io"
	"math/rand"
	"time"
)

type generator struct {
	// Actually rand.Rand implements io.Reader, but for the sake of example we implement it separately.
	rnd rand.Source
}

// Do note that we do not let use generator struct directly, but instead we create a New function that returns a io.Reader.
// the struct is different type but it implements the same interface
func New(seed int64) io.Reader {
	return &generator{
		rnd: rand.NewSource(seed),
	}
}

// Read is the only method we need to implement.
func (g *generator) Read(bytes []byte) (n int, err error) {
	for i := range bytes {
		randInt := g.rnd.Int63()  // positive number from 0 to 2^63
		randByte := byte(randInt) // cast to byte
		bytes[i] = randByte
	}
	return len(bytes), nil
}

func main() {
	generator := New(time.Now().UnixNano()) // seed is the current time in nanoseconds

	buf := make([]byte, 16)

	for i := 0; i < 5; i++ {
		n, _ := generator.Read(buf) // the only method available on interface
		fmt.Printf("Generate bytes: %v size(%d)\n", buf, n)
	}

	// Using existing in std library rand.Rand
	generator2 := rand.New(rand.NewSource(time.Now().UnixNano())) // seed is the current time in nanoseconds

	for i := 0; i < 5; i++ {
		n, _ := generator2.Read(buf) // the only method available on interface
		fmt.Printf("Generate bytes [2]: %v size(%d)\n", buf, n)
	}
}
