package main

import (
	"fmt"
	"os"
	"todo/internal/app"
	"todo/internal/config"
	"todo/internal/parser"
	"todo/internal/storage"
)

func main() {

	cfg, commandArgs, err := config.Load(os.Args[1:])

	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	st := storage.NewFileStorage(cfg.StoragePath)
	app := app.NewApp(st)

	cmd, err := parser.ParseArgs(commandArgs)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
	if err := app.Execute(cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
