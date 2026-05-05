package main

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"time"
)

const s = `lorem ipsum dolor sit amet consectetur adipiscing elit sed do eiusmod tempor incididunt 
ut labore et dolore magna aliqua ut enim ad minim veniam quis 
nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat duis aute irure dolor 
in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur excepteur 
sint occaecat cupidatat non proident sunt in culpa qui officia deserunt mollit anim id est laborum`

func main() {
	r := strings.NewReader(s)

	// buffer 16 bytes
	b := make([]byte, 16)

	for {
		// strings.Reader copies 16 bytes into buffer b
		//
		// the structure remembers the last pointer,
		// so the next call it will read the next 16 bytes
		//
		// returns io.EOF error when the end of the string is reached
		n, err := r.Read(b)

		// when working with io.Reader first you should check if n > 0
		// only then err != nil
		//
		// there may be a situation when some data was read into buffer, saved and then an error occurred
		// then you will have both n > 0 and err != nil
		if n > 0 {
			fmt.Printf("[%s]: %v\n", time.Now().Format("01:11:11.000000"), b)
		}

		if err != nil {
			// break cycle when io.EOF
			if errors.Is(err, io.EOF) {
				break
			}

			fmt.Printf("error: %v\n", err)
		}
	}
}
