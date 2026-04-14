package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"todo/internal/config"
	"todo/internal/parser"
	"todo/internal/service"
	"todo/internal/storage"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	cfg, commandArgs, err := config.Load(os.Args[1:])
	if err != nil {
		logger.Error("Failed config", "error", err)
		os.Exit(1)
	}

	repo := storage.NewFileStorage(cfg.StoragePath)
	service := service.NewTaskService(repo)
	cmd, err := parser.ParseArgs(commandArgs, service)
	if err != nil {
		logger.Error("Failed command", "error", err)
		os.Exit(1)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	done := make(chan error, 1)

	go func() {
		done <- cmd.Execute(ctx)
	}()

	select {
	case <-ctx.Done():
		logger.Info("Received termination signal, exiting")
	case err := <-done:
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}

}
