package main

import (
	"fmt"
	"os"
	"todo/internal/app"
	"todo/internal/parser"
	"todo/internal/storage"
)

const JsonPath = "tasks.json"

func main() {

	st := storage.NewFileStorage(JsonPath)
	app := app.NewApp(st)

	cmd, err := parser.ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := app.Execute(cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
