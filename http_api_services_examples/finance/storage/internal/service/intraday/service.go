package intraday

import (
	"context"
	"fmt"
	"storage/internal/domain/intraday"
	"strings"
)

type Service struct {
	repo intraday.Repository
}

func NewService(repo intraday.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) validateCreate(p *intraday.CreatePayload) error {
	switch {
	case strings.TrimSpace(p.Ticker) == "":
		return fmt.Errorf("%w: ticker is required", intraday.ErrInvalidInput)
	case p.Price <= 0:
		return fmt.Errorf("%w: price must be positive", intraday.ErrInvalidInput)
	case p.Timestamp.IsZero():
		return fmt.Errorf("%w: timestamp is required", intraday.ErrInvalidInput)
	}
	return nil
}

func (s *Service) Create(ctx context.Context, payload *intraday.CreatePayload) error {
	if err := s.validateCreate(payload); err != nil {
		return err
	}

	return s.repo.Create(ctx, payload)
}

func (s *Service) List(ctx context.Context, filter intraday.ListFilter) ([]intraday.Intraday, error) {
	return s.repo.List(ctx, filter)
}
