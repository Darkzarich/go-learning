package intraday

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"time"

	kafka "github.com/segmentio/kafka-go"

	di "consumer/internal/domain/intraday"
)

type IntradayService interface {
	ProcessIntraday(ctx context.Context, intraday di.Intraday) error
}

type Consumer struct {
	reader *kafka.Reader
	svc    IntradayService
}

func NewConsumer(reader *kafka.Reader, svc IntradayService) *Consumer {
	return &Consumer{
		reader: reader,
		svc:    svc,
	}
}

func (c *Consumer) Run(ctx context.Context) error {
	for {
		msg, err := c.reader.FetchMessage(ctx)

		if err != nil {
			if ctx.Err() != nil {
				return nil
			}

			slog.Error("Failed to fetch message", "error", err)
			continue
		}

		log := slog.With("topic", msg.Topic, "partition", msg.Partition, "offset", msg.Offset)

		if err := c.handle(ctx, msg); err != nil {
			if ctx.Err() != nil {
				return nil
			}
			log.Error("Skipping message", "error", err)
		}

		if err := c.reader.CommitMessages(ctx, msg); err != nil {
			log.Error("Failed to commit message", "error", err)
		}
	}
}

func (c *Consumer) handle(ctx context.Context, msg kafka.Message) error {
	var m IntradayMessage
	if err := json.Unmarshal(msg.Value, &m); err != nil {
		return fmt.Errorf("unmarshal %q: %w", msg.Value, err)
	}

	backoff := 500 * time.Millisecond
	for {
		err := c.svc.ProcessIntraday(ctx, m.toDomain())
		if err == nil || errors.Is(err, di.ErrRejected) {
			return err
		}

		if ctx.Err() != nil {
			return ctx.Err()
		}

		slog.Warn("Storage unavailable, retrying", "error", err, "backoff", backoff.String())

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(backoff):
		}

		// Экспоненциальное увеличение ожидания перед каждым следующим ретраем
		backoff = min(backoff*2, 30*time.Second)
	}
}
