/*
	While not as common as sync.Mutex, channels can be used for mutual exclusion, e.g., limiting access to a resource.
*/

package main

import (
	"fmt"
	"time"
)

type Resource struct {
	// some shared state
	value int
}

func (r *Resource) DoWork() {
	r.value++
	fmt.Printf("Resource value: %d\n", r.value)
	time.Sleep(time.Second)
}

func main() {
	resource := &Resource{}

	// a channel of capacity 1 acts as a binary semaphore
	sem := make(chan struct{}, 1)

	for i := 0; i < 5; i++ {
		go func(id int) {
			sem <- struct{}{}        // acquire
			defer func() { <-sem }() // release

			fmt.Printf("Worker %d acquired lock\n", id)
			resource.DoWork()
		}(i)
	}

	time.Sleep(6 * time.Second) // just for demo; in real code use WaitGroup
}
