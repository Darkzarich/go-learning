package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

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

	Client := storage.NewClient(httpClient, cfg.StorageUrl)

	intradayService := si.NewService(Client)

	intradayConsumer := initIntradayConsumer(cfg, intradayService)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		intradayConsumer.Run(ctx)
	}()

	<-ctx.Done()

	slog.Info("Shutting down...")

	wg.Wait()

	return nil
}

func initIntradayConsumer(cfg config.Config, intradayService *si.Service) *hi.Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{cfg.KafkaBrokerUrl},
		Topic:    "tickers",
		GroupID:  "consumer",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	return hi.NewConsumer(reader, intradayService)
}
