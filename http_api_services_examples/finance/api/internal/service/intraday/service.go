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

	return intradays, nil
}
