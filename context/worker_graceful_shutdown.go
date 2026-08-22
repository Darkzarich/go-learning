package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, jobs <-chan int) {
	defer wg.Done()
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}
			fmt.Println("Processing job", job)
			time.Sleep(500 * time.Millisecond)
		case <-ctx.Done():
			fmt.Println("Worker canceled")
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobs := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, &wg, jobs)
	}

	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- i
			time.Sleep(200 * time.Millisecond)
		}
		close(jobs)
	}()

	// Listen for Ctrl+C
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	<-sigCh

	fmt.Println("Shutdown signal received")
	cancel() // stop workers
	wg.Wait()
	fmt.Println("All workers stopped")
}
