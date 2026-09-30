package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"consumer/client/storage"
	"consumer/config"
	hi "consumer/handler/kafka/intraday"
	si "consumer/service/intraday"
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

	storageClient := storage.NewStorageClient(cfg.StorageUrl)

	intradayService := si.NewService(storageClient)

	intradayConsumer := hi.NewConsumer(cfg, intradayService)

	go intradayConsumer.Run(ctx)

	<-ctx.Done()

	slog.Info("Shutting down...")

	return nil
}
