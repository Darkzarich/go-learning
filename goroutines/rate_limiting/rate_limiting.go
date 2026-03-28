/*
	Production systems often need to limit the rate of outgoing requests or processing.
	A channel combined with a time.Ticker makes this straightforward.
*/

package main

import (
	"fmt"
	"time"
)

func rateLimitedWorker(jobs chan int, rate time.Duration) {
	limiter := time.NewTicker(rate)
	defer limiter.Stop()

	for job := range jobs {
		<-limiter.C // wait for next tick
		fmt.Printf("Processing job %d at %s\n", job, time.Now().Format(time.RFC3339))
	}
}

func main() {
	jobs := make(chan int, 10)

	// send 10 jobs
	for i := 1; i <= 10; i++ {
		jobs <- i
	}
	close(jobs)

	// process at most 2 jobs per second
	rateLimitedWorker(jobs, 500*time.Millisecond)
}
