package main

import (
	"io"
	"log"
	"os"
	"strings"
)

type myLimitedReader struct {
	r io.Reader
	n int
}

func (lr *myLimitedReader) Read(p []byte) (n int, err error) {
	if lr.n == 0 {
		return 0, io.EOF
	}

	if lr.n < len(p) {
		p = p[:lr.n]
	}

	readCount, err := lr.r.Read(p)
	lr.n -= readCount

	return readCount, nil
}

func LimitReader(r io.Reader, n int) io.Reader {
	return &myLimitedReader{r, n}
}

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := LimitReader(r, 4)

	_, err := io.Copy(os.Stdout, lr)
	if err != nil {
		log.Fatal(err)
	}
}
