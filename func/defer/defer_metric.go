package main

import (
	"fmt"
	"time"
)

func main() {
	VeryLongTimeFunction()
}

func metricTime(start time.Time) {
	fmt.Println(time.Now().Sub(start))
}

func VeryLongTimeFunction() {
	defer metricTime(time.Now()) // passing current time, it will be calculated before function return
	time.Sleep(time.Second * 5)
}
