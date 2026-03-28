package main

import (
    "fmt"
    "sync"
    "time"
)

type Job struct {
    ID int
}

func worker(id int, jobs chan Job, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job.ID)
        time.Sleep(time.Second) // simulate work
    }
}

func main() {
    const numWorkers = 3
    const numJobs = 10

    jobs := make(chan Job, numJobs) // buffered channel

    var wg sync.WaitGroup

    // start workers
    for w := 1; w <= numWorkers; w++ {
        wg.Add(1)
        go worker(w, jobs, &wg)
    }

    // send jobs
    for j := 1; j <= numJobs; j++ {
        jobs <- Job{ID: j}
    }
    close(jobs) // no more jobs - workers will exit after processing remaining

    // wait for all workers to finish
    wg.Wait()
    fmt.Println("All jobs done")
}