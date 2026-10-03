package intraday

import (
	"time"

	di "consumer/internal/domain/intraday"
)

type IntradayMessage struct {
	Ticker    string    `json:"ticker"`
	Price     float64   `json:"price"`
	Timestamp time.Time `json:"timestamp"`
}

func (i IntradayMessage) toDomain() di.Intraday {
	return di.Intraday{
		Ticker:    i.Ticker,
		Price:     i.Price,
		Timestamp: i.Timestamp,
	}
}
