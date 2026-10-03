package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"

	"consumer/internal/client/storage"
	"consumer/internal/config"
	hi "consumer/internal/handler/kafka/intraday"
	si "consumer/internal/service/intraday"

	"github.com/segmentio/kafka-go"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("config load failed", "error", err)
		os.Exit(1)
	}

	setupSlogger(cfg.Env)

	if err := run(cfg); err != nil {
		slog.Error("fatal", "error", err)
		os.Exit(1)
	}
}

func setupSlogger(env string) {
	var logLevel slog.Level

	switch env {
	case "dev":
		logLevel = slog.LevelDebug
	case "prod":
		logLevel = slog.LevelInfo
	default:
		logLevel = slog.LevelDebug
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))
	slog.SetDefault(logger)
}

func run(cfg config.Config) error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	httpClient := &http.Client{Timeout: 5 * time.Second}

	client := storage.NewClient(httpClient, cfg.StorageURL)

	intradayService := si.NewService(client)

	g, ctx := errgroup.WithContext(ctx)

	intradayConsumer, intradayReader := initIntradayConsumer(cfg, intradayService)
	defer func() {
		if err := intradayReader.Close(); err != nil {
			slog.Error("close intraday reader", "error", err)
		}
	}()

	g.Go(func() error { return intradayConsumer.Run(ctx) })

	slog.Info("Consumer started")

	err := g.Wait()

	slog.Info("Shutting down...")

	return err
}

func initIntradayConsumer(cfg config.Config, intradayService *si.Service) (*hi.Consumer, *kafka.Reader) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{cfg.KafkaBrokerURL},
		Topic:   "tickers",
		GroupID: "consumer",
	})

	return hi.NewConsumer(reader, intradayService), reader
}
