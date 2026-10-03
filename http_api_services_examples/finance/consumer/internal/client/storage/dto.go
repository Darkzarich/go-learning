package storage

import (
	"time"

	di "consumer/internal/domain/intraday"
)

type saveIntradayReq struct {
	Ticker    string    `json:"ticker"`
	Price     float64   `json:"price"`
	Timestamp time.Time `json:"timestamp"`
}

func newSaveIntradayReq(i di.Intraday) saveIntradayReq {
	return saveIntradayReq{Ticker: i.Ticker, Price: i.Price, Timestamp: i.Timestamp}
}
