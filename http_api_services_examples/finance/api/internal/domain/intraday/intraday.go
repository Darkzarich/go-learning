package intraday

import "time"

type Intraday struct {
	TickerID  int64
	Name      string
	Price     float64
	Timestamp time.Time
}

type ListFilter struct {
	TickerID int64
	From     time.Time
	To       time.Time
}
