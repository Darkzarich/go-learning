package main

import (
	"fmt"
	"time"
)

type User struct {
	Login       string
	lastVisited time.Time
}

func UpdateVisitedTime(p *User) {
	p.lastVisited = time.Now()
}

func main() {
	user := User{
		Login: "test_user",
	}

	UpdateVisitedTime(&user) // Pass pointer to a user

	fmt.Printf("user.lastVisited: %s\n", user.lastVisited)
}
