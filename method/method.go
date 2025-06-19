package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u *User) SayHi() {
	fmt.Println("Hi, my name is", u.Name)
}

func main() {
	user := User{
		Name: "Max",
		Age:  28,
	}

	user.SayHi()
}