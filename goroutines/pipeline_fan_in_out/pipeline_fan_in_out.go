package main

import (
	"fmt"
	"sync"
)

func generate(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func fanOut(in chan int, workers int) []chan int {
	channels := make([]chan int, workers)
	for i := 0; i < workers; i++ {
		channels[i] = square(in) // each worker gets its own square stage
	}
	return channels
}

func fanIn(channels ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	for _, c := range channels {
		wg.Add(1)
		go func(ch chan int) {
			defer wg.Done()
			for n := range ch {
				out <- n
			}
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	// generate numbers
	numbers := generate(1, 2, 3, 4, 5)

	// fan out to three workers, each squares the numbers
	workers := fanOut(numbers, 3)

	// fan in results from all workers
	results := fanIn(workers...)

	// collect results, not sorted
	for r := range results {
		fmt.Println(r)
	}
}
