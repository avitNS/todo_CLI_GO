package main

import (
	"fmt"
	"log/slog"
	"os"
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

	if err := app.Execute(cmd); err != nil {
		fmt.Printf("Error: %v\n", err)
		logger.Error("Failed execute", "error", err)
		os.Exit(1)
	}

}
