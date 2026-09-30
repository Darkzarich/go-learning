package intraday

import (
	"context"

	"consumer/client/storage"
)

type StorageClient interface {
	SaveIntraday(ctx context.Context, intraday storage.Intraday) error
}

type Service struct {
	storageClient StorageClient
}

func NewService(storageClient StorageClient) *Service {
	return &Service{storageClient: storageClient}
}

func (s *Service) ProcessIntraday(ctx context.Context, user storage.Intraday) error {
	return s.storageClient.SaveIntraday(ctx, user)
}
