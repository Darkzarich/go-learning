package main

import "time"

type Ticker struct {
	Name     string
	AvgPrice float64
}

type TickerEmulatorDTO struct {
	Ticker    string    `json:"ticker"`
	Price     float64   `json:"price"`
	Timestamp time.Time `json:"timestamp"`
}
