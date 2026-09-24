package intraday

import "time"

type createIntradayReq struct {
	Ticker    string    `json:"ticker"`
	Price     float64   `json:"price"`
	Timestamp time.Time `json:"timestamp"`
}

type createIntradayResp struct {
	Message string `json:"message"`
}

type getIntradaysResp struct {
	TickerID int64   `json:"ticker_id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Time     string  `json:"time"`
}
