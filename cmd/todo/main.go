package main

import (
	"todo/internal/app"
	"todo/internal/storage"
)

const JsonPath = "tasks.json"

func main() {

	app := app.NewApp(storage.NewFileStorage(JsonPath))
	app.Execute()

}
