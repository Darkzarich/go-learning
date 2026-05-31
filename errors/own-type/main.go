package main

import (
	"errors"
	"fmt"
	"time"
)

// Our own type that implements the error interface.
// TimeError - stores time and text.
type TimeError struct {
	Time time.Time
	Text string
}

// Error method is required by the error interface.
func (te TimeError) Error() string {
	return fmt.Sprintf("%v: %v", te.Time.Format(`2006/01/02 15:04:05`), te.Text)
}

// Constructor for our own type.
func NewTimeError(text string) TimeError {
	return TimeError{
		Time: time.Now(),
		Text: text,
	}
}

// Returning our own type as an error. It works.
func testFunc(i int) error {
	if i == 0 {
		return NewTimeError(`parameter is zero`)
	}
	return nil
}

func main() {
	if err := testFunc(0); err != nil {
		fmt.Println(err)
	}

	// Casting error to our own type that implements the error interface to access the data.
	if err := testFunc(0); err != nil {
		if v, ok := err.(TimeError); ok {
			fmt.Println(v.Time, v.Text)
		} else {
			fmt.Println(err)
		}
	}

	// Another, better way to do it.
	err := testFunc(0)
	var v TimeError
	if errors.As(err, &v) {
		fmt.Println(v.Time, v.Text)
	}
}
