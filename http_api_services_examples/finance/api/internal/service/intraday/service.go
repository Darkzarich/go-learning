package intraday

import (
	"context"
	"fmt"
	"time"

	di "api/internal/domain/intraday"
)

const defaultPeriod = 24 * time.Hour

type StorageClient interface {
	FetchIntradays(ctx context.Context, filter di.ListFilter) ([]di.Intraday, error)
}

type Service struct {
	storageClient StorageClient
}

func NewService(storageClient StorageClient) *Service {
	return &Service{storageClient: storageClient}
}

func (s *Service) ListIntradays(ctx context.Context, filter di.ListFilter) ([]di.Intraday, error) {
	if filter.From.IsZero() && filter.To.IsZero() {
		filter.To = time.Now()
		filter.From = filter.To.Add(-defaultPeriod)
	}

	intradays, err := s.storageClient.FetchIntradays(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("fetch intradays: %w", err)
	}

	return lastPerMinute(intradays), nil
}

func lastPerMinute(trades []di.Intraday) []di.Intraday {
	result := make([]di.Intraday, 0, len(trades))

	for _, t := range trades {
		t.Timestamp = t.Timestamp.Truncate(time.Minute)

		if n := len(result); n > 0 && result[n-1].TickerID == t.TickerID && result[n-1].Timestamp.Equal(t.Timestamp) {
			result[n-1] = t
			continue
		}

		result = append(result, t)
	}

	return result
}
