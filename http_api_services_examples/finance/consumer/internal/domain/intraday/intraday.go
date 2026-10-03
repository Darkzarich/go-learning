package intraday

import "time"

type Intraday struct {
	Ticker    string
	Price     float64
	Timestamp time.Time
}
