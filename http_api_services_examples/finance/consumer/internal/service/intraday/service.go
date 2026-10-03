package intraday

import (
	"context"

	di "consumer/internal/domain/intraday"
)

type StorageClient interface {
	SaveIntraday(ctx context.Context, intraday di.Intraday) error
}

type Service struct {
	storageClient StorageClient
}

func NewService(storageClient StorageClient) *Service {
	return &Service{storageClient: storageClient}
}

func (s *Service) ProcessIntraday(ctx context.Context, intraday di.Intraday) error {
	return s.storageClient.SaveIntraday(ctx, intraday)
}
