package main

import (
	"company_and_worker/company"
	"company_and_worker/person"
	"company_and_worker/robot"
	"fmt"
)

func main() {
	c := company.Company{}

	p := person.Person{}
	p.SetName("John")

	r := person.Person{}
	r.SetName("Mary")

	robot := robot.Robot{}

	// Method expects Worker interface as parameter
	// but we can hire Person and Robot because they implement Worker interface
	c.Hire(&p)
	c.Hire(&r)
	c.Hire(&robot)

	fmt.Println(c.Process(0, []string{"task1", "task2"}))
	fmt.Println(c.Process(1, []string{"planning"}))
	fmt.Println(c.Process(2, []string{"cleaning"}))
}
