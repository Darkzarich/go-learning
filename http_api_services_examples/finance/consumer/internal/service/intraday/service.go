package intraday

import (
	"context"

	di "consumer/internal/domain/intraday"
)

type Client interface {
	SaveIntraday(ctx context.Context, intraday di.Intraday) error
}

type Service struct {
	Client Client
}

func NewService(Client Client) *Service {
	return &Service{Client: Client}
}

func (s *Service) ProcessIntraday(ctx context.Context, intraday di.Intraday) error {
	return s.Client.SaveIntraday(ctx, intraday)
}
