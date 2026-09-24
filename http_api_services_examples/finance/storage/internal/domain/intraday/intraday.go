package intraday

import "time"

type Intraday struct {
	ID       int64
	TickerID int64
	Ticker   string
	// float64 only for simplicity, it's better to use
	// github.com/shopspring/decimal
	Price     float64
	Timestamp time.Time
}
