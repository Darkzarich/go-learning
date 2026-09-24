package intraday

import (
	"context"
	"time"
)

type CreatePayload struct {
	Ticker    string
	Price     float64
	Timestamp time.Time
}

type ListFilter struct {
	TickerID string
	From     time.Time
	To       time.Time
}

type Repository interface {
	Create(ctx context.Context, payload *CreatePayload) error
	List(ctx context.Context, filter ListFilter) ([]Intraday, error)
}
