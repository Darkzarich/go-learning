/*
	context is widely used in production Go to propagate cancellation across goroutines.
	Under the hood, it often uses a channel to signal done.
*/

package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int, jobs chan int, results chan int) {
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return // jobs channel closed
			}
			// simulate work
			time.Sleep(time.Duration(job) * 200 * time.Millisecond)
			results <- job * 2
		case <-ctx.Done():
			fmt.Printf("Worker %d cancelled\n", id)
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// start workers
	for w := 1; w <= 3; w++ {
		go worker(ctx, w, jobs, results)
	}

	// send jobs
	for i := 1; i <= 10; i++ {
		jobs <- i
	}
	close(jobs)

	// collect results (or timeout)
	for i := 0; i < 10; i++ {
		select {
		case res := <-results:
			fmt.Println("Result:", res)
		case <-ctx.Done():
			fmt.Println("Timeout, stopping")
			return
		}
	}
}
