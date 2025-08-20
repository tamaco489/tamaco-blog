package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/handler"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/library/logger"
)

func main() {
	// Context for initialization
	ctx := context.Background()

	// Logger initialization
	logger.Init(logger.Config{
		Level:  logger.LevelInfo,
		Format: "json",
	})

	// Handler initialization with context
	srv, err := handler.NewHandler(ctx)
	if err != nil {
		logger.Error("failed to create handler", "error", err)
		os.Exit(1)
	}

	go func() {
		logger.Info("starting server", "addr", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("failed to start server", "error", err)
			os.Exit(1)
		}
	}()

	logger.Info("server started successfully", "addr", srv.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	logger.Info("shutdown signal received, gracefully shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("server shutdown error", "error", err)
		os.Exit(1)
	}

	<-ctx.Done()

	switch ctx.Err() {
	case context.DeadlineExceeded:
		logger.Warn("shutdown timeout exceeded, forcing exit")
	default:
		logger.Info("server shutdown successfully")
	}
}
