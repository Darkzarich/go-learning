/*
Task:

Write MyCheck function that checks the string for multiple errors.
Define a type that describes a list of errors, and prints all errors
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Errors []error

func (e Errors) Error() string {
	var errs []string
	for _, err := range e {
		errs = append(errs, err.Error())
	}
	return strings.Join(errs, ";")
}

func MyCheck(s string) error {
	var errs Errors

	if strings.ContainsAny(s, "0123456789") {
		errs = append(errs, fmt.Errorf("string contains digits"))
	}

	if strings.Count(s, " ") < 2 {
		errs = append(errs, fmt.Errorf("string does not contain at least two spaces"))
	}

	if len(s) > 20 {
		errs = append(errs, fmt.Errorf("string is too long"))
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Specify any string (\"q\" to quit): ")

		ret, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		ret = strings.TrimRight(ret, "\r\n")
		if ret == `q` {
			break
		}

		if err = MyCheck(ret); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(`The string passed the check`)
		}
	}
}
