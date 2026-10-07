package intraday

import (
	"context"
	"fmt"

	di "api/internal/domain/intraday"
)

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
	intradays, err := s.storageClient.FetchIntradays(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("fetch intradays: %w", err)
	}

	return intradays, nil
}
