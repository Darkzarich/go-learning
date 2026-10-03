package intraday

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"

	"consumer/internal/config"

	kafka "github.com/segmentio/kafka-go"

	di "consumer/internal/domain/intraday"
)

type IIntradayService interface {
	ProcessIntraday(ctx context.Context, intraday di.Intraday) error
}

type Consumer struct {
	reader *kafka.Reader
	svc    IIntradayService
}

func NewConsumer(cfg config.Config, svc IIntradayService) *Consumer {
	return &Consumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  []string{cfg.KafkaBrokerUrl},
			Topic:    "tickers",
			GroupID:  "consumer",
			MinBytes: 10e3, // 10KB
			MaxBytes: 10e6, // 10MB
		}),
		svc: svc,
	}
}

func (c *Consumer) Run(ctx context.Context) error {
	defer c.reader.Close()

	for {
		msg, err := c.reader.FetchMessage(ctx)
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				slog.Info("Context deadline exceeded")
				return context.DeadlineExceeded
			}

			if errors.Is(err, context.Canceled) {
				slog.Info("Context canceled")
				return context.Canceled
			}

			if errors.Is(err, io.EOF) {
				slog.Info("Kafka reader closed")
				return nil
			}
			slog.Error("Failed to fetch message", "error", err)
			continue
		}

		var intraday IntradayMessage

		if err := json.Unmarshal(msg.Value, &intraday); err != nil {
			slog.Error("Failed to unmarshal message", "error", err)
			continue
		}

		if err := c.svc.ProcessIntraday(ctx, intraday.toDomain()); err != nil {
			slog.Warn("Failed to process message, skipping", "error", err)
		}

		if err := c.reader.CommitMessages(ctx, msg); err != nil {
			slog.Error("Failed to commit message", "error", err)
		}
	}
}
