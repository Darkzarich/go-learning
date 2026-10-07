package intraday

type getIntradaysResp struct {
	TickerID int64   `json:"ticker_id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Time     string  `json:"time"`
}
