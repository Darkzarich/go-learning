package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"storage/internal/config"
	hh "storage/internal/handler/health"
	hi "storage/internal/handler/intraday"
	ri "storage/internal/repository/intraday"
	si "storage/internal/service/intraday"

	"github.com/jackc/pgx/v5/pgxpool"
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
	// Graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	pgConfig, err := pgxpool.ParseConfig(cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("parse DB config: %w", err)
	}

	slog.Info("Connecting to database", "host", pgConfig.ConnConfig.Host, "db", pgConfig.ConnConfig.Database)

	pool, err := pgxpool.NewWithConfig(ctx, pgConfig)
	if err != nil {
		return fmt.Errorf("new pool: %w", err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		return fmt.Errorf("ping db: %w", err)
	}
	slog.Info("connected to database")

	repo := ri.NewPostgresRepo(pool)
	svc := si.NewService(repo)

	mux := http.NewServeMux()

	hi.NewHandler(svc).Routes(mux)
	hh.NewHandler().Routes(mux)

	server := &http.Server{
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	ln, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		return fmt.Errorf("listen on port %q: %w", cfg.Port, err)
	}

	slog.Info("Server is running on " + ln.Addr().String())

	errCh := make(chan error, 1)
	go func() {
		if err := server.Serve(ln); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		slog.Info("Shutting down...")
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	return server.Shutdown(shutdownCtx)
}
