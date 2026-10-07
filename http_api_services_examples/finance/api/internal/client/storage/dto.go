package storage

import (
	"time"
)

type getIntradaysRespItem struct {
	TickerID int64     `json:"ticker_id"`
	Name     string    `json:"name"`
	Price    float64   `json:"price"`
	Time     time.Time `json:"time"`
}

type getIntradaysResp = []getIntradaysRespItem
