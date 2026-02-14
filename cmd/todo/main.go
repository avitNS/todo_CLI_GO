package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"todo/internal/app"
	"todo/internal/config"
	"todo/internal/parser"
	"todo/internal/storage"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	cfg, commandArgs, err := config.Load(os.Args[1:])
	if err != nil {
		logger.Error("Failed config", "error", err)
		os.Exit(1)
	}

	st := storage.NewFileStorage(cfg.StoragePath)

	app := app.NewApp(st)

	cmd, err := parser.ParseArgs(commandArgs)
	if err != nil {
		logger.Error("Failed command", "error", err)
		os.Exit(1)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	done := make(chan error, 1)

	go func() {
		done <- app.Execute(cmd)
	}()

	select {
	case <-ctx.Done():
		logger.Info("Received termination signal, exiting")
		os.Exit(1)
	case err := <-done:
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			logger.Error("Failed execute", "error", err)
			os.Exit(1)
		}
	}

}
