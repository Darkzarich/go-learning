package main

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	intervals []time.Duration
	start     time.Time
}

func (s *Stopwatch) Start() {
	s.start = time.Now()
}

func (s *Stopwatch) SaveSplit() {
	s.intervals = append(s.intervals, time.Since(s.start))
}

func (s Stopwatch) GetResults() []time.Duration {
	return s.intervals
}

func main() {
	sw := Stopwatch{}
	sw.Start()

	time.Sleep(1 * time.Second)
	sw.SaveSplit()

	time.Sleep(500 * time.Millisecond)
	sw.SaveSplit()

	time.Sleep(300 * time.Millisecond)
	sw.SaveSplit()

	fmt.Println(sw.GetResults())
}
